// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/subscription_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlanType int32

const (
	PlanType_PLAN_TYPE_UNSPECIFIED PlanType = 0
	PlanType_FREE                  PlanType = 1
	PlanType_PRO                   PlanType = 2
)

// Enum value maps for PlanType.
var (
	PlanType_name = map[int32]string{
		0: "PLAN_TYPE_UNSPECIFIED",
		1: "FREE",
		2: "PRO",
	}
	PlanType_value = map[string]int32{
		"PLAN_TYPE_UNSPECIFIED": 0,
		"FREE":                  1,
		"PRO":                   2,
	}
)

func (x PlanType) Enum() *PlanType {
	p := new(PlanType)
	*p = x
	return p
}

func (x PlanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_subscription_service_proto_enumTypes[0].Descriptor()
}

func (PlanType) Type() protoreflect.EnumType {
	return &file_api_v2_subscription_service_proto_enumTypes[0]
}

func (x PlanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanType.Descriptor instead.
func (PlanType) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan        PlanType               `protobuf:"varint,1,opt,name=plan,proto3,enum=slash.api.v2.PlanType" json:"plan,omitempty"`
	StartedTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	ExpiresTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_subscription_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_subscription_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetPlan() PlanType {
	if x != nil {
		return x.Plan
	}
	return PlanType_PLAN_TYPE_UNSPECIFIED
}

func (x *Subscription) GetStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedTime
	}
	return nil
}

func (x *Subscription) GetExpiresTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresTime
	}
	return nil
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_subscription_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_subscription_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{1}
}

type GetSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *GetSubscriptionResponse) Reset() {
	*x = GetSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_subscription_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionResponse) ProtoMessage() {}

func (x *GetSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_subscription_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubscriptionResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type UpdateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseKey string `protobuf:"bytes,1,opt,name=license_key,json=licenseKey,proto3" json:"license_key,omitempty"`
}

func (x *UpdateSubscriptionRequest) Reset() {
	*x = UpdateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_subscription_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionRequest) ProtoMessage() {}

func (x *UpdateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_subscription_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSubscriptionRequest) GetLicenseKey() string {
	if x != nil {
		return x.LicenseKey
	}
	return ""
}

type UpdateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *UpdateSubscriptionResponse) Reset() {
	*x = UpdateSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_subscription_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionResponse) ProtoMessage() {}

func (x *UpdateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_subscription_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_subscription_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSubscriptionResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

var File_api_v2_subscription_service_proto protoreflect.FileDescriptor

var file_api_v2_subscription_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x41, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x4b, 0x65, 0x79, 0x22, 0x5c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x38, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x52, 0x45, 0x45,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x52, 0x4f, 0x10, 0x02, 0x32, 0x96, 0x02, 0x0a, 0x13,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x84, 0x01,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x32, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0xb6, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64,
	0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03,
	0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x32, 0xca, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x32, 0xe2, 0x02, 0x18, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_subscription_service_proto_rawDescOnce sync.Once
	file_api_v2_subscription_service_proto_rawDescData = file_api_v2_subscription_service_proto_rawDesc
)

func file_api_v2_subscription_service_proto_rawDescGZIP() []byte {
	file_api_v2_subscription_service_proto_rawDescOnce.Do(func() {
		file_api_v2_subscription_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_subscription_service_proto_rawDescData)
	})
	return file_api_v2_subscription_service_proto_rawDescData
}

var file_api_v2_subscription_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v2_subscription_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v2_subscription_service_proto_goTypes = []interface{}{
	(PlanType)(0),                      // 0: slash.api.v2.PlanType
	(*Subscription)(nil),               // 1: slash.api.v2.Subscription
	(*GetSubscriptionRequest)(nil),     // 2: slash.api.v2.GetSubscriptionRequest
	(*GetSubscriptionResponse)(nil),    // 3: slash.api.v2.GetSubscriptionResponse
	(*UpdateSubscriptionRequest)(nil),  // 4: slash.api.v2.UpdateSubscriptionRequest
	(*UpdateSubscriptionResponse)(nil), // 5: slash.api.v2.UpdateSubscriptionResponse
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
}
var file_api_v2_subscription_service_proto_depIdxs = []int32{
	0, // 0: slash.api.v2.Subscription.plan:type_name -> slash.api.v2.PlanType
	6, // 1: slash.api.v2.Subscription.started_time:type_name -> google.protobuf.Timestamp
	6, // 2: slash.api.v2.Subscription.expires_time:type_name -> google.protobuf.Timestamp
	1, // 3: slash.api.v2.GetSubscriptionResponse.subscription:type_name -> slash.api.v2.Subscription
	1, // 4: slash.api.v2.UpdateSubscriptionResponse.subscription:type_name -> slash.api.v2.Subscription
	2, // 5: slash.api.v2.SubscriptionService.GetSubscription:input_type -> slash.api.v2.GetSubscriptionRequest
	4, // 6: slash.api.v2.SubscriptionService.UpdateSubscription:input_type -> slash.api.v2.UpdateSubscriptionRequest
	3, // 7: slash.api.v2.SubscriptionService.GetSubscription:output_type -> slash.api.v2.GetSubscriptionResponse
	5, // 8: slash.api.v2.SubscriptionService.UpdateSubscription:output_type -> slash.api.v2.UpdateSubscriptionResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v2_subscription_service_proto_init() }
func file_api_v2_subscription_service_proto_init() {
	if File_api_v2_subscription_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v2_subscription_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_api_v2_subscription_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_api_v2_subscription_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionResponse); i {
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
		file_api_v2_subscription_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubscriptionRequest); i {
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
		file_api_v2_subscription_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubscriptionResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_subscription_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_subscription_service_proto_goTypes,
		DependencyIndexes: file_api_v2_subscription_service_proto_depIdxs,
		EnumInfos:         file_api_v2_subscription_service_proto_enumTypes,
		MessageInfos:      file_api_v2_subscription_service_proto_msgTypes,
	}.Build()
	File_api_v2_subscription_service_proto = out.File
	file_api_v2_subscription_service_proto_rawDesc = nil
	file_api_v2_subscription_service_proto_goTypes = nil
	file_api_v2_subscription_service_proto_depIdxs = nil
}
