// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: acs/roles.proto

package acs

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId      *int32   `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3,oneof" json:"role_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetRoleId() int32 {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	Role           *Role  `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *CreateRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type CreateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRoleResponse) Reset() {
	*x = CreateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleResponse) ProtoMessage() {}

func (x *CreateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{2}
}

type ReadRoleOwnersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	RoleName       string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *ReadRoleOwnersRequest) Reset() {
	*x = ReadRoleOwnersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRoleOwnersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRoleOwnersRequest) ProtoMessage() {}

func (x *ReadRoleOwnersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRoleOwnersRequest.ProtoReflect.Descriptor instead.
func (*ReadRoleOwnersRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRoleOwnersRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *ReadRoleOwnersRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type ReadRoleOwnersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds [][]byte `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *ReadRoleOwnersResponse) Reset() {
	*x = ReadRoleOwnersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRoleOwnersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRoleOwnersResponse) ProtoMessage() {}

func (x *ReadRoleOwnersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRoleOwnersResponse.ProtoReflect.Descriptor instead.
func (*ReadRoleOwnersResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRoleOwnersResponse) GetUserIds() [][]byte {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type CheckRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string   `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	UserId         []byte   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleNames      []string `protobuf:"bytes,3,rep,name=role_names,json=roleNames,proto3" json:"role_names,omitempty"`
}

func (x *CheckRolesRequest) Reset() {
	*x = CheckRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRolesRequest) ProtoMessage() {}

func (x *CheckRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRolesRequest.ProtoReflect.Descriptor instead.
func (*CheckRolesRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{5}
}

func (x *CheckRolesRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *CheckRolesRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CheckRolesRequest) GetRoleNames() []string {
	if x != nil {
		return x.RoleNames
	}
	return nil
}

type CheckRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckRolesResponse) Reset() {
	*x = CheckRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRolesResponse) ProtoMessage() {}

func (x *CheckRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRolesResponse.ProtoReflect.Descriptor instead.
func (*CheckRolesResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{6}
}

func (x *CheckRolesResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	RoleName       string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRoleRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *DeleteRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type DeleteRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoleResponse) Reset() {
	*x = DeleteRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleResponse) ProtoMessage() {}

func (x *DeleteRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{8}
}

type AddRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	UserId         []byte `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleName       string `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *AddRoleRequest) Reset() {
	*x = AddRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleRequest) ProtoMessage() {}

func (x *AddRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleRequest.ProtoReflect.Descriptor instead.
func (*AddRoleRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{9}
}

func (x *AddRoleRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *AddRoleRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *AddRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type AddRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRoleResponse) Reset() {
	*x = AddRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleResponse) ProtoMessage() {}

func (x *AddRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleResponse.ProtoReflect.Descriptor instead.
func (*AddRoleResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{10}
}

type RemoveRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterToken string `protobuf:"bytes,1,opt,name=requester_token,json=requesterToken,proto3" json:"requester_token,omitempty"`
	UserId         []byte `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleName       string `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *RemoveRoleRequest) Reset() {
	*x = RemoveRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleRequest) ProtoMessage() {}

func (x *RemoveRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoleRequest) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveRoleRequest) GetRequesterToken() string {
	if x != nil {
		return x.RequesterToken
	}
	return ""
}

func (x *RemoveRoleRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *RemoveRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type RemoveRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRoleResponse) Reset() {
	*x = RemoveRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acs_roles_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleResponse) ProtoMessage() {}

func (x *RemoveRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acs_roles_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleResponse.ProtoReflect.Descriptor instead.
func (*RemoveRoleResponse) Descriptor() ([]byte, []int) {
	return file_acs_roles_proto_rawDescGZIP(), []int{12}
}

var File_acs_roles_proto protoreflect.FileDescriptor

var file_acs_roles_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x63, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x63, 0x73, 0x22, 0x66, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x5b,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x63, 0x73,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5d, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x33, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x74, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x22, 0x59, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6f, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xf0, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61,
	0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x73,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e,
	0x61, 0x63, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x73,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x61, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acs_roles_proto_rawDescOnce sync.Once
	file_acs_roles_proto_rawDescData = file_acs_roles_proto_rawDesc
)

func file_acs_roles_proto_rawDescGZIP() []byte {
	file_acs_roles_proto_rawDescOnce.Do(func() {
		file_acs_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_acs_roles_proto_rawDescData)
	})
	return file_acs_roles_proto_rawDescData
}

var file_acs_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_acs_roles_proto_goTypes = []interface{}{
	(*Role)(nil),                   // 0: acs.Role
	(*CreateRoleRequest)(nil),      // 1: acs.CreateRoleRequest
	(*CreateRoleResponse)(nil),     // 2: acs.CreateRoleResponse
	(*ReadRoleOwnersRequest)(nil),  // 3: acs.ReadRoleOwnersRequest
	(*ReadRoleOwnersResponse)(nil), // 4: acs.ReadRoleOwnersResponse
	(*CheckRolesRequest)(nil),      // 5: acs.CheckRolesRequest
	(*CheckRolesResponse)(nil),     // 6: acs.CheckRolesResponse
	(*DeleteRoleRequest)(nil),      // 7: acs.DeleteRoleRequest
	(*DeleteRoleResponse)(nil),     // 8: acs.DeleteRoleResponse
	(*AddRoleRequest)(nil),         // 9: acs.AddRoleRequest
	(*AddRoleResponse)(nil),        // 10: acs.AddRoleResponse
	(*RemoveRoleRequest)(nil),      // 11: acs.RemoveRoleRequest
	(*RemoveRoleResponse)(nil),     // 12: acs.RemoveRoleResponse
}
var file_acs_roles_proto_depIdxs = []int32{
	0,  // 0: acs.CreateRoleRequest.role:type_name -> acs.Role
	1,  // 1: acs.Roles.Create:input_type -> acs.CreateRoleRequest
	3,  // 2: acs.Roles.ReadOwners:input_type -> acs.ReadRoleOwnersRequest
	5,  // 3: acs.Roles.CheckRoles:input_type -> acs.CheckRolesRequest
	7,  // 4: acs.Roles.Delete:input_type -> acs.DeleteRoleRequest
	9,  // 5: acs.Roles.Add:input_type -> acs.AddRoleRequest
	11, // 6: acs.Roles.Remove:input_type -> acs.RemoveRoleRequest
	2,  // 7: acs.Roles.Create:output_type -> acs.CreateRoleResponse
	4,  // 8: acs.Roles.ReadOwners:output_type -> acs.ReadRoleOwnersResponse
	6,  // 9: acs.Roles.CheckRoles:output_type -> acs.CheckRolesResponse
	8,  // 10: acs.Roles.Delete:output_type -> acs.DeleteRoleResponse
	10, // 11: acs.Roles.Add:output_type -> acs.AddRoleResponse
	12, // 12: acs.Roles.Remove:output_type -> acs.RemoveRoleResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_acs_roles_proto_init() }
func file_acs_roles_proto_init() {
	if File_acs_roles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acs_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_acs_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_acs_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleResponse); i {
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
		file_acs_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRoleOwnersRequest); i {
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
		file_acs_roles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRoleOwnersResponse); i {
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
		file_acs_roles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRolesRequest); i {
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
		file_acs_roles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRolesResponse); i {
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
		file_acs_roles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
		file_acs_roles_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleResponse); i {
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
		file_acs_roles_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleRequest); i {
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
		file_acs_roles_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleResponse); i {
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
		file_acs_roles_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoleRequest); i {
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
		file_acs_roles_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoleResponse); i {
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
	file_acs_roles_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acs_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acs_roles_proto_goTypes,
		DependencyIndexes: file_acs_roles_proto_depIdxs,
		MessageInfos:      file_acs_roles_proto_msgTypes,
	}.Build()
	File_acs_roles_proto = out.File
	file_acs_roles_proto_rawDesc = nil
	file_acs_roles_proto_goTypes = nil
	file_acs_roles_proto_depIdxs = nil
}
