// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pkg/services/grpc/auth/auth.proto

package auth

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IsAuthorizedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Role         string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	CompanyID    string `protobuf:"bytes,3,opt,name=companyID,proto3" json:"companyID,omitempty"`
	RepositoryID string `protobuf:"bytes,4,opt,name=repositoryID,proto3" json:"repositoryID,omitempty"`
}

func (x *IsAuthorizedData) Reset() {
	*x = IsAuthorizedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAuthorizedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedData) ProtoMessage() {}

func (x *IsAuthorizedData) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedData.ProtoReflect.Descriptor instead.
func (*IsAuthorizedData) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *IsAuthorizedData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *IsAuthorizedData) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *IsAuthorizedData) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

func (x *IsAuthorizedData) GetRepositoryID() string {
	if x != nil {
		return x.RepositoryID
	}
	return ""
}

type IsAuthorizedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuthorized bool `protobuf:"varint,1,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
}

func (x *IsAuthorizedResponse) Reset() {
	*x = IsAuthorizedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAuthorizedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedResponse) ProtoMessage() {}

func (x *IsAuthorizedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedResponse.ProtoReflect.Descriptor instead.
func (*IsAuthorizedResponse) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *IsAuthorizedResponse) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

type GetAccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAccountData) Reset() {
	*x = GetAccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountData) ProtoMessage() {}

func (x *GetAccountData) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountData.ProtoReflect.Descriptor instead.
func (*GetAccountData) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetAccountDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID   string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *GetAccountDataResponse) Reset() {
	*x = GetAccountDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDataResponse) ProtoMessage() {}

func (x *GetAccountDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDataResponse.ProtoReflect.Descriptor instead.
func (*GetAccountDataResponse) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountDataResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *GetAccountDataResponse) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type GetAuthConfigData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthConfigData) Reset() {
	*x = GetAuthConfigData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthConfigData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConfigData) ProtoMessage() {}

func (x *GetAuthConfigData) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConfigData.ProtoReflect.Descriptor instead.
func (*GetAuthConfigData) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{4}
}

type GetAuthConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationAdminEnable bool   `protobuf:"varint,1,opt,name=ApplicationAdminEnable,proto3" json:"ApplicationAdminEnable,omitempty"`
	AuthType               string `protobuf:"bytes,2,opt,name=AuthType,proto3" json:"AuthType,omitempty"`
	DisabledBroker         bool   `protobuf:"varint,3,opt,name=DisabledBroker,proto3" json:"DisabledBroker,omitempty"`
}

func (x *GetAuthConfigResponse) Reset() {
	*x = GetAuthConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConfigResponse) ProtoMessage() {}

func (x *GetAuthConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConfigResponse.ProtoReflect.Descriptor instead.
func (*GetAuthConfigResponse) Descriptor() ([]byte, []int) {
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthConfigResponse) GetApplicationAdminEnable() bool {
	if x != nil {
		return x.ApplicationAdminEnable
	}
	return false
}

func (x *GetAuthConfigResponse) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *GetAuthConfigResponse) GetDisabledBroker() bool {
	if x != nil {
		return x.DisabledBroker
	}
	return false
}

var File_development_kit_pkg_services_grpc_auth_auth_proto protoreflect.FileDescriptor

var file_development_kit_pkg_services_grpc_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x31, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6b, 0x69,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x7e, 0x0a, 0x10, 0x49, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x14, 0x49, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x32, 0xe2, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescOnce sync.Once
	file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescData = file_development_kit_pkg_services_grpc_auth_auth_proto_rawDesc
)

func file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescGZIP() []byte {
	file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescOnce.Do(func() {
		file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescData)
	})
	return file_development_kit_pkg_services_grpc_auth_auth_proto_rawDescData
}

var file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_development_kit_pkg_services_grpc_auth_auth_proto_goTypes = []interface{}{
	(*IsAuthorizedData)(nil),       // 0: grpc.IsAuthorizedData
	(*IsAuthorizedResponse)(nil),   // 1: grpc.IsAuthorizedResponse
	(*GetAccountData)(nil),         // 2: grpc.GetAccountData
	(*GetAccountDataResponse)(nil), // 3: grpc.GetAccountDataResponse
	(*GetAuthConfigData)(nil),      // 4: grpc.GetAuthConfigData
	(*GetAuthConfigResponse)(nil),  // 5: grpc.GetAuthConfigResponse
}
var file_development_kit_pkg_services_grpc_auth_auth_proto_depIdxs = []int32{
	0, // 0: grpc.AuthService.IsAuthorized:input_type -> grpc.IsAuthorizedData
	2, // 1: grpc.AuthService.GetAccountID:input_type -> grpc.GetAccountData
	4, // 2: grpc.AuthService.GetAuthConfig:input_type -> grpc.GetAuthConfigData
	1, // 3: grpc.AuthService.IsAuthorized:output_type -> grpc.IsAuthorizedResponse
	3, // 4: grpc.AuthService.GetAccountID:output_type -> grpc.GetAccountDataResponse
	5, // 5: grpc.AuthService.GetAuthConfig:output_type -> grpc.GetAuthConfigResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_development_kit_pkg_services_grpc_auth_auth_proto_init() }
func file_development_kit_pkg_services_grpc_auth_auth_proto_init() {
	if File_development_kit_pkg_services_grpc_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAuthorizedData); i {
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
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAuthorizedResponse); i {
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
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountData); i {
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
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDataResponse); i {
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
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthConfigData); i {
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
		file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthConfigResponse); i {
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
			RawDescriptor: file_development_kit_pkg_services_grpc_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_development_kit_pkg_services_grpc_auth_auth_proto_goTypes,
		DependencyIndexes: file_development_kit_pkg_services_grpc_auth_auth_proto_depIdxs,
		MessageInfos:      file_development_kit_pkg_services_grpc_auth_auth_proto_msgTypes,
	}.Build()
	File_development_kit_pkg_services_grpc_auth_auth_proto = out.File
	file_development_kit_pkg_services_grpc_auth_auth_proto_rawDesc = nil
	file_development_kit_pkg_services_grpc_auth_auth_proto_goTypes = nil
	file_development_kit_pkg_services_grpc_auth_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedData, opts ...grpc.CallOption) (*IsAuthorizedResponse, error)
	GetAccountID(ctx context.Context, in *GetAccountData, opts ...grpc.CallOption) (*GetAccountDataResponse, error)
	GetAuthConfig(ctx context.Context, in *GetAuthConfigData, opts ...grpc.CallOption) (*GetAuthConfigResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) IsAuthorized(ctx context.Context, in *IsAuthorizedData, opts ...grpc.CallOption) (*IsAuthorizedResponse, error) {
	out := new(IsAuthorizedResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccountID(ctx context.Context, in *GetAccountData, opts ...grpc.CallOption) (*GetAccountDataResponse, error) {
	out := new(GetAccountDataResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAuthConfig(ctx context.Context, in *GetAuthConfigData, opts ...grpc.CallOption) (*GetAuthConfigResponse, error) {
	out := new(GetAuthConfigResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetAuthConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	IsAuthorized(context.Context, *IsAuthorizedData) (*IsAuthorizedResponse, error)
	GetAccountID(context.Context, *GetAccountData) (*GetAccountDataResponse, error)
	GetAuthConfig(context.Context, *GetAuthConfigData) (*GetAuthConfigResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) IsAuthorized(context.Context, *IsAuthorizedData) (*IsAuthorizedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}
func (*UnimplementedAuthServiceServer) GetAccountID(context.Context, *GetAccountData) (*GetAccountDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountID not implemented")
}
func (*UnimplementedAuthServiceServer) GetAuthConfig(context.Context, *GetAuthConfigData) (*GetAuthConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthConfig not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsAuthorized(ctx, req.(*IsAuthorizedData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAccountID(ctx, req.(*GetAccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAuthConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthConfigData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAuthConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetAuthConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAuthConfig(ctx, req.(*GetAuthConfigData))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _AuthService_IsAuthorized_Handler,
		},
		{
			MethodName: "GetAccountID",
			Handler:    _AuthService_GetAccountID_Handler,
		},
		{
			MethodName: "GetAuthConfig",
			Handler:    _AuthService_GetAuthConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "development-kit/pkg/services/grpc/auth/auth.proto",
}
