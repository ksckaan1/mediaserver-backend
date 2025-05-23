// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: tmdb_service.proto

package tmdbpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type GetTMDBInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTMDBInfoRequest) Reset() {
	*x = GetTMDBInfoRequest{}
	mi := &file_tmdb_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTMDBInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTMDBInfoRequest) ProtoMessage() {}

func (x *GetTMDBInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tmdb_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTMDBInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTMDBInfoRequest) Descriptor() ([]byte, []int) {
	return file_tmdb_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTMDBInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TMDBInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Data          *structpb.Struct       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TMDBInfo) Reset() {
	*x = TMDBInfo{}
	mi := &file_tmdb_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TMDBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TMDBInfo) ProtoMessage() {}

func (x *TMDBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tmdb_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TMDBInfo.ProtoReflect.Descriptor instead.
func (*TMDBInfo) Descriptor() ([]byte, []int) {
	return file_tmdb_service_proto_rawDescGZIP(), []int{1}
}

func (x *TMDBInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TMDBInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *TMDBInfo) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tmdb_service_proto protoreflect.FileDescriptor

const file_tmdb_service_proto_rawDesc = "" +
	"\n" +
	"\x12tmdb_service.proto\x12\x06tmdbpb\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\"$\n" +
	"\x12GetTMDBInfoRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x82\x01\n" +
	"\bTMDBInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"updated_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12+\n" +
	"\x04data\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04data2J\n" +
	"\vTMDBService\x12;\n" +
	"\vGetTMDBInfo\x12\x1a.tmdbpb.GetTMDBInfoRequest\x1a\x10.tmdbpb.TMDBInfoB\x19Z\x17shared/pb/tmdbpb;tmdbpbb\x06proto3"

var (
	file_tmdb_service_proto_rawDescOnce sync.Once
	file_tmdb_service_proto_rawDescData []byte
)

func file_tmdb_service_proto_rawDescGZIP() []byte {
	file_tmdb_service_proto_rawDescOnce.Do(func() {
		file_tmdb_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tmdb_service_proto_rawDesc), len(file_tmdb_service_proto_rawDesc)))
	})
	return file_tmdb_service_proto_rawDescData
}

var file_tmdb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tmdb_service_proto_goTypes = []any{
	(*GetTMDBInfoRequest)(nil),    // 0: tmdbpb.GetTMDBInfoRequest
	(*TMDBInfo)(nil),              // 1: tmdbpb.TMDBInfo
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 3: google.protobuf.Struct
}
var file_tmdb_service_proto_depIdxs = []int32{
	2, // 0: tmdbpb.TMDBInfo.updated_at:type_name -> google.protobuf.Timestamp
	3, // 1: tmdbpb.TMDBInfo.data:type_name -> google.protobuf.Struct
	0, // 2: tmdbpb.TMDBService.GetTMDBInfo:input_type -> tmdbpb.GetTMDBInfoRequest
	1, // 3: tmdbpb.TMDBService.GetTMDBInfo:output_type -> tmdbpb.TMDBInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tmdb_service_proto_init() }
func file_tmdb_service_proto_init() {
	if File_tmdb_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tmdb_service_proto_rawDesc), len(file_tmdb_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tmdb_service_proto_goTypes,
		DependencyIndexes: file_tmdb_service_proto_depIdxs,
		MessageInfos:      file_tmdb_service_proto_msgTypes,
	}.Build()
	File_tmdb_service_proto = out.File
	file_tmdb_service_proto_goTypes = nil
	file_tmdb_service_proto_depIdxs = nil
}
