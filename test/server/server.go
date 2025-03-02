package testserver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/yourselfhosted/slash/api/auth"
	"github.com/yourselfhosted/slash/server"
	"github.com/yourselfhosted/slash/server/profile"
	"github.com/yourselfhosted/slash/store"
	"github.com/yourselfhosted/slash/store/db"
	"github.com/yourselfhosted/slash/test"
)

type TestingServer struct {
	server  *server.Server
	client  *http.Client
	profile *profile.Profile
	cookie  string
}

func NewTestingServer(ctx context.Context, t *testing.T) (*TestingServer, error) {
	profile := test.GetTestingProfile(t)
	dbDriver, err := db.NewDBDriver(profile)
	if err != nil {
		fmt.Printf("failed to create db driver, error: %+v\n", err)
		return nil, err
	}
	if err := dbDriver.Migrate(ctx); err != nil {
		fmt.Printf("failed to migrate db, error: %+v\n", err)
		return nil, err
	}

	store := store.New(dbDriver, profile)
	server, err := server.NewServer(ctx, profile, store)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create server")
	}

	s := &TestingServer{
		server:  server,
		client:  &http.Client{},
		profile: profile,
		cookie:  "",
	}
	errChan := make(chan error, 1)

	go func() {
		if err := s.server.Start(ctx); err != nil {
			if err != http.ErrServerClosed {
				errChan <- errors.Wrap(err, "failed to run main server")
			}
		}
	}()

	if err := s.waitForServerStart(errChan); err != nil {
		return nil, errors.Wrap(err, "failed to start server")
	}

	return s, nil
}

func (s *TestingServer) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

func (s *TestingServer) waitForServerStart(errChan <-chan error) error {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if s == nil {
				continue
			}
			e := s.server.GetEcho()
			if e == nil {
				continue
			}
			addr := e.ListenerAddr()
			if addr != nil && strings.Contains(addr.String(), ":") {
				return nil // was started
			}
		case err := <-errChan:
			if err == http.ErrServerClosed {
				return nil
			}
			return err
		}
	}
}

func (s *TestingServer) request(method, uri string, body io.Reader, params, header map[string]string) (io.ReadCloser, error) {
	fullURL := fmt.Sprintf("http://localhost:%d%s", s.profile.Port, uri)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to create a new %s request(%q)", method, fullURL)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to send a %s request(%q)", method, fullURL)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read http response body")
		}
		return nil, errors.Errorf("http response error code %v body %q", resp.StatusCode, string(body))
	}

	if method == "POST" {
		if strings.Contains(uri, "/api/v1/auth/signin") || strings.Contains(uri, "/api/v1/auth/signup") {
			cookie := ""
			h := resp.Header.Get("Set-Cookie")
			parts := strings.Split(h, "; ")
			for _, p := range parts {
				if strings.HasPrefix(p, fmt.Sprintf("%s=", auth.AccessTokenCookieName)) {
					cookie = p
					break
				}
			}
			if cookie == "" {
				return nil, errors.Errorf("unable to find access token in the login response headers")
			}
			s.cookie = cookie
		} else if strings.Contains(uri, "/api/v1/auth/logout") {
			s.cookie = ""
		}
	}
	return resp.Body, nil
}

// get sends a GET client request.
func (s *TestingServer) get(url string, params map[string]string) (io.ReadCloser, error) {
	return s.request("GET", url, nil, params, map[string]string{
		"Cookie": s.cookie,
	})
}

// post sends a POST client request.
func (s *TestingServer) post(url string, body io.Reader, params map[string]string) (io.ReadCloser, error) {
	return s.request("POST", url, body, params, map[string]string{
		"Cookie": s.cookie,
	})
}

// patch sends a PATCH client request.
func (s *TestingServer) patch(url string, body io.Reader, params map[string]string) (io.ReadCloser, error) {
	return s.request("PATCH", url, body, params, map[string]string{
		"Cookie": s.cookie,
	})
}

// delete sends a DELETE client request.
func (s *TestingServer) delete(url string, params map[string]string) (io.ReadCloser, error) {
	return s.request("DELETE", url, nil, params, map[string]string{
		"Cookie": s.cookie,
	})
}
