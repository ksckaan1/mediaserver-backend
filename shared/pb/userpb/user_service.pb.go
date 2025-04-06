// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user_service.proto

package userpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int64                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	mi := &file_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUsersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*User                `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Limit         int64                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListUsersResponse) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListUsersResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListUsersResponse) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUsersResponse) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetUserByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIDRequest) Reset() {
	*x = GetUserByIDRequest{}
	mi := &file_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRequest) ProtoMessage() {}

func (x *GetUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserByUsernameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByUsernameRequest) Reset() {
	*x = GetUserByUsernameRequest{}
	mi := &file_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameRequest) ProtoMessage() {}

func (x *GetUserByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UpdateUserPasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserPasswordRequest) Reset() {
	*x = UpdateUserPasswordRequest{}
	mi := &file_user_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPasswordRequest) ProtoMessage() {}

func (x *UpdateUserPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserPasswordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_user_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

const file_user_service_proto_rawDesc = "" +
	"\n" +
	"\x12user_service.proto\x12\x06userpb\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xc4\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x1a\n" +
	"\busername\x18\x04 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x05 \x01(\tR\bpassword\"K\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"-\n" +
	"\x12CreateUserResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"@\n" +
	"\x10ListUsersRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x03R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x03R\x06offset\"y\n" +
	"\x11ListUsersResponse\x12 \n" +
	"\x04list\x18\x01 \x03(\v2\f.userpb.UserR\x04list\x12\x14\n" +
	"\x05count\x18\x02 \x01(\x03R\x05count\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x03R\x05limit\x12\x16\n" +
	"\x06offset\x18\x04 \x01(\x03R\x06offset\"$\n" +
	"\x12GetUserByIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x18GetUserByUsernameRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\"G\n" +
	"\x19UpdateUserPasswordRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"#\n" +
	"\x11DeleteUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xa4\x03\n" +
	"\vUserService\x12C\n" +
	"\n" +
	"CreateUser\x12\x19.userpb.CreateUserRequest\x1a\x1a.userpb.CreateUserResponse\x12@\n" +
	"\tListUsers\x12\x18.userpb.ListUsersRequest\x1a\x19.userpb.ListUsersResponse\x127\n" +
	"\vGetUserByID\x12\x1a.userpb.GetUserByIDRequest\x1a\f.userpb.User\x12C\n" +
	"\x11GetUserByUsername\x12 .userpb.GetUserByUsernameRequest\x1a\f.userpb.User\x12O\n" +
	"\x12UpdateUserPassword\x12!.userpb.UpdateUserPasswordRequest\x1a\x16.google.protobuf.Empty\x12?\n" +
	"\n" +
	"DeleteUser\x12\x19.userpb.DeleteUserRequest\x1a\x16.google.protobuf.EmptyB\x19Z\x17shared/pb/userpb;userpbb\x06proto3"

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData []byte
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)))
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_service_proto_goTypes = []any{
	(*User)(nil),                      // 0: userpb.User
	(*CreateUserRequest)(nil),         // 1: userpb.CreateUserRequest
	(*CreateUserResponse)(nil),        // 2: userpb.CreateUserResponse
	(*ListUsersRequest)(nil),          // 3: userpb.ListUsersRequest
	(*ListUsersResponse)(nil),         // 4: userpb.ListUsersResponse
	(*GetUserByIDRequest)(nil),        // 5: userpb.GetUserByIDRequest
	(*GetUserByUsernameRequest)(nil),  // 6: userpb.GetUserByUsernameRequest
	(*UpdateUserPasswordRequest)(nil), // 7: userpb.UpdateUserPasswordRequest
	(*DeleteUserRequest)(nil),         // 8: userpb.DeleteUserRequest
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 10: google.protobuf.Empty
}
var file_user_service_proto_depIdxs = []int32{
	9,  // 0: userpb.User.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: userpb.User.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: userpb.ListUsersResponse.list:type_name -> userpb.User
	1,  // 3: userpb.UserService.CreateUser:input_type -> userpb.CreateUserRequest
	3,  // 4: userpb.UserService.ListUsers:input_type -> userpb.ListUsersRequest
	5,  // 5: userpb.UserService.GetUserByID:input_type -> userpb.GetUserByIDRequest
	6,  // 6: userpb.UserService.GetUserByUsername:input_type -> userpb.GetUserByUsernameRequest
	7,  // 7: userpb.UserService.UpdateUserPassword:input_type -> userpb.UpdateUserPasswordRequest
	8,  // 8: userpb.UserService.DeleteUser:input_type -> userpb.DeleteUserRequest
	2,  // 9: userpb.UserService.CreateUser:output_type -> userpb.CreateUserResponse
	4,  // 10: userpb.UserService.ListUsers:output_type -> userpb.ListUsersResponse
	0,  // 11: userpb.UserService.GetUserByID:output_type -> userpb.User
	0,  // 12: userpb.UserService.GetUserByUsername:output_type -> userpb.User
	10, // 13: userpb.UserService.UpdateUserPassword:output_type -> google.protobuf.Empty
	10, // 14: userpb.UserService.DeleteUser:output_type -> google.protobuf.Empty
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
