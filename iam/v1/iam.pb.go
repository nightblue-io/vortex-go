// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: iam/v1/iam.proto

package iam

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Request message for the Login rpc.
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If set, use to call user profile.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Optional. If non-empty, do a device auth flow.
	DeviceAuth string `protobuf:"bytes,4,opt,name=device_auth,json=deviceAuth,proto3" json:"device_auth,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetDeviceAuth() string {
	if x != nil {
		return x.DeviceAuth
	}
	return ""
}

// Response message for the Login rpc.
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceAuth  *LoginResponse_DeviceAuth `protobuf:"bytes,1,opt,name=device_auth,json=deviceAuth,proto3" json:"device_auth,omitempty"`
	AccessToken string                    `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Status      *status.Status            `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetDeviceAuth() *LoginResponse_DeviceAuth {
	if x != nil {
		return x.DeviceAuth
	}
	return nil
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// Request message for the WhoAmI rpc.
type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{2}
}

func (x *WhoAmIRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// Request message for the WhoAmI rpc.
type WhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *WhoAmIResponse) Reset() {
	*x = WhoAmIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResponse) ProtoMessage() {}

func (x *WhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{3}
}

func (x *WhoAmIResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WhoAmIResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WhoAmIResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginResponse_DeviceAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCode                string `protobuf:"bytes,1,opt,name=user_code,json=userCode,proto3" json:"user_code,omitempty"`
	VerificationUriComplete string `protobuf:"bytes,2,opt,name=verification_uri_complete,json=verificationUriComplete,proto3" json:"verification_uri_complete,omitempty"`
	ExpiresIn               int64  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *LoginResponse_DeviceAuth) Reset() {
	*x = LoginResponse_DeviceAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_v1_iam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse_DeviceAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse_DeviceAuth) ProtoMessage() {}

func (x *LoginResponse_DeviceAuth) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_iam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse_DeviceAuth.ProtoReflect.Descriptor instead.
func (*LoginResponse_DeviceAuth) Descriptor() ([]byte, []int) {
	return file_iam_v1_iam_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LoginResponse_DeviceAuth) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *LoginResponse_DeviceAuth) GetVerificationUriComplete() string {
	if x != nil {
		return x.VerificationUriComplete
	}
	return ""
}

func (x *LoginResponse_DeviceAuth) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_iam_v1_iam_proto protoreflect.FileDescriptor

var file_iam_v1_iam_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x22, 0xb4, 0x02,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x84, 0x01,
	0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x22, 0x32, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x0e, 0x57, 0x68, 0x6f, 0x41,
	0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x32, 0x89, 0x03, 0x0a, 0x03, 0x49, 0x61, 0x6d, 0x12, 0x6f, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x30, 0x01, 0x12, 0x6e, 0x0a,
	0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x68, 0x6f,
	0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x6f, 0x72,
	0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x1a, 0xa0, 0x01,
	0x92, 0x41, 0x9c, 0x01, 0x12, 0x33, 0x28, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x29, 0x20, 0x49, 0x41,
	0x4d, 0x20, 0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a,
	0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x62, 0x6c, 0x75, 0x65, 0x2e, 0x69, 0x6f, 0x1a, 0x65, 0x0a, 0x24, 0x53, 0x65, 0x65,
	0x20, 0x68, 0x65, 0x72, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x12, 0x3d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x62, 0x6c, 0x75, 0x65, 0x2d,
	0x69, 0x6f, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x6f, 0x72, 0x74, 0x65, 0x78, 0x2d, 0x69, 0x6f, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_v1_iam_proto_rawDescOnce sync.Once
	file_iam_v1_iam_proto_rawDescData = file_iam_v1_iam_proto_rawDesc
)

func file_iam_v1_iam_proto_rawDescGZIP() []byte {
	file_iam_v1_iam_proto_rawDescOnce.Do(func() {
		file_iam_v1_iam_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_v1_iam_proto_rawDescData)
	})
	return file_iam_v1_iam_proto_rawDescData
}

var file_iam_v1_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_iam_v1_iam_proto_goTypes = []any{
	(*LoginRequest)(nil),             // 0: vortexproto.iam.v1.LoginRequest
	(*LoginResponse)(nil),            // 1: vortexproto.iam.v1.LoginResponse
	(*WhoAmIRequest)(nil),            // 2: vortexproto.iam.v1.WhoAmIRequest
	(*WhoAmIResponse)(nil),           // 3: vortexproto.iam.v1.WhoAmIResponse
	(*LoginResponse_DeviceAuth)(nil), // 4: vortexproto.iam.v1.LoginResponse.DeviceAuth
	(*status.Status)(nil),            // 5: google.rpc.Status
}
var file_iam_v1_iam_proto_depIdxs = []int32{
	4, // 0: vortexproto.iam.v1.LoginResponse.device_auth:type_name -> vortexproto.iam.v1.LoginResponse.DeviceAuth
	5, // 1: vortexproto.iam.v1.LoginResponse.status:type_name -> google.rpc.Status
	0, // 2: vortexproto.iam.v1.Iam.Login:input_type -> vortexproto.iam.v1.LoginRequest
	2, // 3: vortexproto.iam.v1.Iam.WhoAmI:input_type -> vortexproto.iam.v1.WhoAmIRequest
	1, // 4: vortexproto.iam.v1.Iam.Login:output_type -> vortexproto.iam.v1.LoginResponse
	3, // 5: vortexproto.iam.v1.Iam.WhoAmI:output_type -> vortexproto.iam.v1.WhoAmIResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_iam_v1_iam_proto_init() }
func file_iam_v1_iam_proto_init() {
	if File_iam_v1_iam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_v1_iam_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
		file_iam_v1_iam_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WhoAmIRequest); i {
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
		file_iam_v1_iam_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WhoAmIResponse); i {
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
		file_iam_v1_iam_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse_DeviceAuth); i {
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
			RawDescriptor: file_iam_v1_iam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_v1_iam_proto_goTypes,
		DependencyIndexes: file_iam_v1_iam_proto_depIdxs,
		MessageInfos:      file_iam_v1_iam_proto_msgTypes,
	}.Build()
	File_iam_v1_iam_proto = out.File
	file_iam_v1_iam_proto_rawDesc = nil
	file_iam_v1_iam_proto_goTypes = nil
	file_iam_v1_iam_proto_depIdxs = nil
}
