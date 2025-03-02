// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/workspace_setting.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkspaceSettingKey int32

const (
	WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED WorkspaceSettingKey = 0
	// The license key.
	WorkspaceSettingKey_WORKSPACE_SETTING_LICENSE_KEY WorkspaceSettingKey = 1
	// The secret session key used to encrypt session data.
	WorkspaceSettingKey_WORKSPACE_SETTING_SECRET_SESSION WorkspaceSettingKey = 2
	// Whether to enable other users to sign up.
	WorkspaceSettingKey_WORKSAPCE_SETTING_ENABLE_SIGNUP WorkspaceSettingKey = 3
	// The custom style.
	WorkspaceSettingKey_WORKSPACE_SETTING_CUSTOM_STYLE WorkspaceSettingKey = 4
	// The custom script.
	WorkspaceSettingKey_WORKSPACE_SETTING_CUSTOM_SCRIPT WorkspaceSettingKey = 5
	// The auto backup setting.
	WorkspaceSettingKey_WORKSPACE_SETTING_AUTO_BACKUP WorkspaceSettingKey = 6
	// The instance URL.
	WorkspaceSettingKey_WORKSPACE_SETTING_INSTANCE_URL WorkspaceSettingKey = 7
)

// Enum value maps for WorkspaceSettingKey.
var (
	WorkspaceSettingKey_name = map[int32]string{
		0: "WORKSPACE_SETTING_KEY_UNSPECIFIED",
		1: "WORKSPACE_SETTING_LICENSE_KEY",
		2: "WORKSPACE_SETTING_SECRET_SESSION",
		3: "WORKSAPCE_SETTING_ENABLE_SIGNUP",
		4: "WORKSPACE_SETTING_CUSTOM_STYLE",
		5: "WORKSPACE_SETTING_CUSTOM_SCRIPT",
		6: "WORKSPACE_SETTING_AUTO_BACKUP",
		7: "WORKSPACE_SETTING_INSTANCE_URL",
	}
	WorkspaceSettingKey_value = map[string]int32{
		"WORKSPACE_SETTING_KEY_UNSPECIFIED": 0,
		"WORKSPACE_SETTING_LICENSE_KEY":     1,
		"WORKSPACE_SETTING_SECRET_SESSION":  2,
		"WORKSAPCE_SETTING_ENABLE_SIGNUP":   3,
		"WORKSPACE_SETTING_CUSTOM_STYLE":    4,
		"WORKSPACE_SETTING_CUSTOM_SCRIPT":   5,
		"WORKSPACE_SETTING_AUTO_BACKUP":     6,
		"WORKSPACE_SETTING_INSTANCE_URL":    7,
	}
)

func (x WorkspaceSettingKey) Enum() *WorkspaceSettingKey {
	p := new(WorkspaceSettingKey)
	*p = x
	return p
}

func (x WorkspaceSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[0].Descriptor()
}

func (WorkspaceSettingKey) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[0]
}

func (x WorkspaceSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceSettingKey.Descriptor instead.
func (WorkspaceSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

type WorkspaceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key WorkspaceSettingKey `protobuf:"varint,1,opt,name=key,proto3,enum=slash.store.WorkspaceSettingKey" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*WorkspaceSetting_LicenseKey
	//	*WorkspaceSetting_SecretSession
	//	*WorkspaceSetting_EnableSignup
	//	*WorkspaceSetting_CustomStyle
	//	*WorkspaceSetting_CustomScript
	//	*WorkspaceSetting_AutoBackup
	//	*WorkspaceSetting_InstanceUrl
	Value isWorkspaceSetting_Value `protobuf_oneof:"value"`
}

func (x *WorkspaceSetting) Reset() {
	*x = WorkspaceSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting) ProtoMessage() {}

func (x *WorkspaceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceSetting) GetKey() WorkspaceSettingKey {
	if x != nil {
		return x.Key
	}
	return WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED
}

func (m *WorkspaceSetting) GetValue() isWorkspaceSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *WorkspaceSetting) GetLicenseKey() string {
	if x, ok := x.GetValue().(*WorkspaceSetting_LicenseKey); ok {
		return x.LicenseKey
	}
	return ""
}

func (x *WorkspaceSetting) GetSecretSession() string {
	if x, ok := x.GetValue().(*WorkspaceSetting_SecretSession); ok {
		return x.SecretSession
	}
	return ""
}

func (x *WorkspaceSetting) GetEnableSignup() bool {
	if x, ok := x.GetValue().(*WorkspaceSetting_EnableSignup); ok {
		return x.EnableSignup
	}
	return false
}

func (x *WorkspaceSetting) GetCustomStyle() string {
	if x, ok := x.GetValue().(*WorkspaceSetting_CustomStyle); ok {
		return x.CustomStyle
	}
	return ""
}

func (x *WorkspaceSetting) GetCustomScript() string {
	if x, ok := x.GetValue().(*WorkspaceSetting_CustomScript); ok {
		return x.CustomScript
	}
	return ""
}

func (x *WorkspaceSetting) GetAutoBackup() *AutoBackupWorkspaceSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_AutoBackup); ok {
		return x.AutoBackup
	}
	return nil
}

func (x *WorkspaceSetting) GetInstanceUrl() string {
	if x, ok := x.GetValue().(*WorkspaceSetting_InstanceUrl); ok {
		return x.InstanceUrl
	}
	return ""
}

type isWorkspaceSetting_Value interface {
	isWorkspaceSetting_Value()
}

type WorkspaceSetting_LicenseKey struct {
	LicenseKey string `protobuf:"bytes,2,opt,name=license_key,json=licenseKey,proto3,oneof"`
}

type WorkspaceSetting_SecretSession struct {
	SecretSession string `protobuf:"bytes,3,opt,name=secret_session,json=secretSession,proto3,oneof"`
}

type WorkspaceSetting_EnableSignup struct {
	EnableSignup bool `protobuf:"varint,4,opt,name=enable_signup,json=enableSignup,proto3,oneof"`
}

type WorkspaceSetting_CustomStyle struct {
	CustomStyle string `protobuf:"bytes,5,opt,name=custom_style,json=customStyle,proto3,oneof"`
}

type WorkspaceSetting_CustomScript struct {
	CustomScript string `protobuf:"bytes,6,opt,name=custom_script,json=customScript,proto3,oneof"`
}

type WorkspaceSetting_AutoBackup struct {
	AutoBackup *AutoBackupWorkspaceSetting `protobuf:"bytes,7,opt,name=auto_backup,json=autoBackup,proto3,oneof"`
}

type WorkspaceSetting_InstanceUrl struct {
	InstanceUrl string `protobuf:"bytes,8,opt,name=instance_url,json=instanceUrl,proto3,oneof"`
}

func (*WorkspaceSetting_LicenseKey) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_SecretSession) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_EnableSignup) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_CustomStyle) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_CustomScript) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_AutoBackup) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_InstanceUrl) isWorkspaceSetting_Value() {}

type AutoBackupWorkspaceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether auto backup is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The cron expression for auto backup.
	// For example, "0 0 0 * * *" means backup at 00:00:00 every day.
	// See https://en.wikipedia.org/wiki/Cron for more details.
	CronExpression string `protobuf:"bytes,2,opt,name=cron_expression,json=cronExpression,proto3" json:"cron_expression,omitempty"`
	// The maximum number of backups to keep.
	MaxKeep int32 `protobuf:"varint,3,opt,name=max_keep,json=maxKeep,proto3" json:"max_keep,omitempty"`
}

func (x *AutoBackupWorkspaceSetting) Reset() {
	*x = AutoBackupWorkspaceSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_workspace_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoBackupWorkspaceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoBackupWorkspaceSetting) ProtoMessage() {}

func (x *AutoBackupWorkspaceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoBackupWorkspaceSetting.ProtoReflect.Descriptor instead.
func (*AutoBackupWorkspaceSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{1}
}

func (x *AutoBackupWorkspaceSetting) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AutoBackupWorkspaceSetting) GetCronExpression() string {
	if x != nil {
		return x.CronExpression
	}
	return ""
}

func (x *AutoBackupWorkspaceSetting) GetMaxKeep() int32 {
	if x != nil {
		return x.MaxKeep
	}
	return 0
}

var File_store_workspace_setting_proto protoreflect.FileDescriptor

var file_store_workspace_setting_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xff, 0x02, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x12, 0x23, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a,
	0x0a, 0x1a, 0x41, 0x75, 0x74, 0x6f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x4b, 0x65, 0x65, 0x70, 0x2a, 0xba, 0x02, 0x0a, 0x13, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x4f, 0x52,
	0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20,
	0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x41, 0x50, 0x43, 0x45, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53,
	0x49, 0x47, 0x4e, 0x55, 0x50, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x57, 0x4f, 0x52, 0x4b, 0x53,
	0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x05,
	0x12, 0x21, 0x0a, 0x1d, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55,
	0x50, 0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43,
	0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x07, 0x42, 0xa6, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x15, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x2f,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_workspace_setting_proto_rawDescOnce sync.Once
	file_store_workspace_setting_proto_rawDescData = file_store_workspace_setting_proto_rawDesc
)

func file_store_workspace_setting_proto_rawDescGZIP() []byte {
	file_store_workspace_setting_proto_rawDescOnce.Do(func() {
		file_store_workspace_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_workspace_setting_proto_rawDescData)
	})
	return file_store_workspace_setting_proto_rawDescData
}

var file_store_workspace_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_workspace_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_workspace_setting_proto_goTypes = []interface{}{
	(WorkspaceSettingKey)(0),           // 0: slash.store.WorkspaceSettingKey
	(*WorkspaceSetting)(nil),           // 1: slash.store.WorkspaceSetting
	(*AutoBackupWorkspaceSetting)(nil), // 2: slash.store.AutoBackupWorkspaceSetting
}
var file_store_workspace_setting_proto_depIdxs = []int32{
	0, // 0: slash.store.WorkspaceSetting.key:type_name -> slash.store.WorkspaceSettingKey
	2, // 1: slash.store.WorkspaceSetting.auto_backup:type_name -> slash.store.AutoBackupWorkspaceSetting
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_workspace_setting_proto_init() }
func file_store_workspace_setting_proto_init() {
	if File_store_workspace_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_workspace_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_workspace_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoBackupWorkspaceSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_store_workspace_setting_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WorkspaceSetting_LicenseKey)(nil),
		(*WorkspaceSetting_SecretSession)(nil),
		(*WorkspaceSetting_EnableSignup)(nil),
		(*WorkspaceSetting_CustomStyle)(nil),
		(*WorkspaceSetting_CustomScript)(nil),
		(*WorkspaceSetting_AutoBackup)(nil),
		(*WorkspaceSetting_InstanceUrl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_workspace_setting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_workspace_setting_proto_goTypes,
		DependencyIndexes: file_store_workspace_setting_proto_depIdxs,
		EnumInfos:         file_store_workspace_setting_proto_enumTypes,
		MessageInfos:      file_store_workspace_setting_proto_msgTypes,
	}.Build()
	File_store_workspace_setting_proto = out.File
	file_store_workspace_setting_proto_rawDesc = nil
	file_store_workspace_setting_proto_goTypes = nil
	file_store_workspace_setting_proto_depIdxs = nil
}
