// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: movie_service.proto

package moviepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type MediaType int32

const (
	MediaType_UNKNOWN MediaType = 0
	MediaType_IMAGE   MediaType = 1
	MediaType_VIDEO   MediaType = 2
	MediaType_AUDIO   MediaType = 3
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "UNKNOWN",
		1: "IMAGE",
		2: "VIDEO",
		3: "AUDIO",
	}
	MediaType_value = map[string]int32{
		"UNKNOWN": 0,
		"IMAGE":   1,
		"VIDEO":   2,
		"AUDIO":   3,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_movie_service_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_movie_service_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{0}
}

type CreateMovieRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TmdbId        string                 `protobuf:"bytes,3,opt,name=tmdb_id,json=tmdbId,proto3" json:"tmdb_id,omitempty"`
	MediaId       string                 `protobuf:"bytes,4,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMovieRequest) Reset() {
	*x = CreateMovieRequest{}
	mi := &file_movie_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieRequest) ProtoMessage() {}

func (x *CreateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMovieRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMovieRequest) GetTmdbId() string {
	if x != nil {
		return x.TmdbId
	}
	return ""
}

func (x *CreateMovieRequest) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

type CreateMovieResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MovieId       string                 `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMovieResponse) Reset() {
	*x = CreateMovieResponse{}
	mi := &file_movie_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieResponse) ProtoMessage() {}

func (x *CreateMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieResponse.ProtoReflect.Descriptor instead.
func (*CreateMovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMovieResponse) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type GetMovieByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MovieId       string                 `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMovieByIDRequest) Reset() {
	*x = GetMovieByIDRequest{}
	mi := &file_movie_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByIDRequest) ProtoMessage() {}

func (x *GetMovieByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByIDRequest.ProtoReflect.Descriptor instead.
func (*GetMovieByIDRequest) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovieByIDRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type Movie struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	MediaInfo     *Media                 `protobuf:"bytes,6,opt,name=media_info,json=mediaInfo,proto3" json:"media_info,omitempty"`
	TmdbInfo      *TMDBInfo              `protobuf:"bytes,7,opt,name=tmdb_info,json=tmdbInfo,proto3" json:"tmdb_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Movie) Reset() {
	*x = Movie{}
	mi := &file_movie_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{3}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Movie) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetMediaInfo() *Media {
	if x != nil {
		return x.MediaInfo
	}
	return nil
}

func (x *Movie) GetTmdbInfo() *TMDBInfo {
	if x != nil {
		return x.TmdbInfo
	}
	return nil
}

type ListMoviesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int64                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMoviesRequest) Reset() {
	*x = ListMoviesRequest{}
	mi := &file_movie_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesRequest) ProtoMessage() {}

func (x *ListMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesRequest.ProtoReflect.Descriptor instead.
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListMoviesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListMoviesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type MovieList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*Movie               `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Limit         int64                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MovieList) Reset() {
	*x = MovieList{}
	mi := &file_movie_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MovieList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieList) ProtoMessage() {}

func (x *MovieList) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieList.ProtoReflect.Descriptor instead.
func (*MovieList) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{5}
}

func (x *MovieList) GetList() []*Movie {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MovieList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MovieList) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *MovieList) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UpdateMovieByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MovieId       string                 `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TmdbId        string                 `protobuf:"bytes,4,opt,name=tmdb_id,json=tmdbId,proto3" json:"tmdb_id,omitempty"`
	MediaId       string                 `protobuf:"bytes,5,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMovieByIDRequest) Reset() {
	*x = UpdateMovieByIDRequest{}
	mi := &file_movie_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMovieByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieByIDRequest) ProtoMessage() {}

func (x *UpdateMovieByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovieByIDRequest) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMovieByIDRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *UpdateMovieByIDRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMovieByIDRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMovieByIDRequest) GetTmdbId() string {
	if x != nil {
		return x.TmdbId
	}
	return ""
}

func (x *UpdateMovieByIDRequest) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

type DeleteMovieByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MovieId       string                 `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMovieByIDRequest) Reset() {
	*x = DeleteMovieByIDRequest{}
	mi := &file_movie_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMovieByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieByIDRequest) ProtoMessage() {}

func (x *DeleteMovieByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovieByIDRequest) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMovieByIDRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type TMDBInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data          *structpb.Struct       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TMDBInfo) Reset() {
	*x = TMDBInfo{}
	mi := &file_movie_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TMDBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TMDBInfo) ProtoMessage() {}

func (x *TMDBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[8]
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
	return file_movie_service_proto_rawDescGZIP(), []int{8}
}

func (x *TMDBInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TMDBInfo) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type Media struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Path          string                 `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Type          MediaType              `protobuf:"varint,6,opt,name=type,proto3,enum=moviepb.MediaType" json:"type,omitempty"`
	MimeType      string                 `protobuf:"bytes,7,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size          int64                  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Media) Reset() {
	*x = Media{}
	mi := &file_movie_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_movie_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_movie_service_proto_rawDescGZIP(), []int{9}
}

func (x *Media) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Media) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Media) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Media) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Media) GetType() MediaType {
	if x != nil {
		return x.Type
	}
	return MediaType_UNKNOWN
}

func (x *Media) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Media) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_movie_service_proto protoreflect.FileDescriptor

const file_movie_service_proto_rawDesc = "" +
	"\n" +
	"\x13movie_service.proto\x12\amoviepb\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"\x80\x01\n" +
	"\x12CreateMovieRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x17\n" +
	"\atmdb_id\x18\x03 \x01(\tR\x06tmdbId\x12\x19\n" +
	"\bmedia_id\x18\x04 \x01(\tR\amediaId\"0\n" +
	"\x13CreateMovieResponse\x12\x19\n" +
	"\bmovie_id\x18\x01 \x01(\tR\amovieId\"0\n" +
	"\x13GetMovieByIDRequest\x12\x19\n" +
	"\bmovie_id\x18\x01 \x01(\tR\amovieId\"\xa4\x02\n" +
	"\x05Movie\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12-\n" +
	"\n" +
	"media_info\x18\x06 \x01(\v2\x0e.moviepb.MediaR\tmediaInfo\x12.\n" +
	"\ttmdb_info\x18\a \x01(\v2\x11.moviepb.TMDBInfoR\btmdbInfo\"A\n" +
	"\x11ListMoviesRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x03R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x03R\x06offset\"s\n" +
	"\tMovieList\x12\"\n" +
	"\x04list\x18\x01 \x03(\v2\x0e.moviepb.MovieR\x04list\x12\x14\n" +
	"\x05count\x18\x02 \x01(\x03R\x05count\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x03R\x05limit\x12\x16\n" +
	"\x06offset\x18\x04 \x01(\x03R\x06offset\"\x9f\x01\n" +
	"\x16UpdateMovieByIDRequest\x12\x19\n" +
	"\bmovie_id\x18\x01 \x01(\tR\amovieId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x17\n" +
	"\atmdb_id\x18\x04 \x01(\tR\x06tmdbId\x12\x19\n" +
	"\bmedia_id\x18\x05 \x01(\tR\amediaId\"3\n" +
	"\x16DeleteMovieByIDRequest\x12\x19\n" +
	"\bmovie_id\x18\x01 \x01(\tR\amovieId\"G\n" +
	"\bTMDBInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12+\n" +
	"\x04data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04data\"\x90\x02\n" +
	"\x05Media\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12\x12\n" +
	"\x04path\x18\x05 \x01(\tR\x04path\x12&\n" +
	"\x04type\x18\x06 \x01(\x0e2\x12.moviepb.MediaTypeR\x04type\x12\x1b\n" +
	"\tmime_type\x18\a \x01(\tR\bmimeType\x12\x12\n" +
	"\x04size\x18\b \x01(\x03R\x04size*9\n" +
	"\tMediaType\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\t\n" +
	"\x05IMAGE\x10\x01\x12\t\n" +
	"\x05VIDEO\x10\x02\x12\t\n" +
	"\x05AUDIO\x10\x032\xec\x02\n" +
	"\fMovieService\x12H\n" +
	"\vCreateMovie\x12\x1b.moviepb.CreateMovieRequest\x1a\x1c.moviepb.CreateMovieResponse\x12<\n" +
	"\fGetMovieByID\x12\x1c.moviepb.GetMovieByIDRequest\x1a\x0e.moviepb.Movie\x12<\n" +
	"\n" +
	"ListMovies\x12\x1a.moviepb.ListMoviesRequest\x1a\x12.moviepb.MovieList\x12J\n" +
	"\x0fUpdateMovieByID\x12\x1f.moviepb.UpdateMovieByIDRequest\x1a\x16.google.protobuf.Empty\x12J\n" +
	"\x0fDeleteMovieByID\x12\x1f.moviepb.DeleteMovieByIDRequest\x1a\x16.google.protobuf.EmptyB\x1bZ\x19shared/pb/moviepb;moviepbb\x06proto3"

var (
	file_movie_service_proto_rawDescOnce sync.Once
	file_movie_service_proto_rawDescData []byte
)

func file_movie_service_proto_rawDescGZIP() []byte {
	file_movie_service_proto_rawDescOnce.Do(func() {
		file_movie_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_movie_service_proto_rawDesc), len(file_movie_service_proto_rawDesc)))
	})
	return file_movie_service_proto_rawDescData
}

var file_movie_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_movie_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_movie_service_proto_goTypes = []any{
	(MediaType)(0),                 // 0: moviepb.MediaType
	(*CreateMovieRequest)(nil),     // 1: moviepb.CreateMovieRequest
	(*CreateMovieResponse)(nil),    // 2: moviepb.CreateMovieResponse
	(*GetMovieByIDRequest)(nil),    // 3: moviepb.GetMovieByIDRequest
	(*Movie)(nil),                  // 4: moviepb.Movie
	(*ListMoviesRequest)(nil),      // 5: moviepb.ListMoviesRequest
	(*MovieList)(nil),              // 6: moviepb.MovieList
	(*UpdateMovieByIDRequest)(nil), // 7: moviepb.UpdateMovieByIDRequest
	(*DeleteMovieByIDRequest)(nil), // 8: moviepb.DeleteMovieByIDRequest
	(*TMDBInfo)(nil),               // 9: moviepb.TMDBInfo
	(*Media)(nil),                  // 10: moviepb.Media
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 12: google.protobuf.Struct
	(*emptypb.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_movie_service_proto_depIdxs = []int32{
	11, // 0: moviepb.Movie.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: moviepb.Movie.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: moviepb.Movie.media_info:type_name -> moviepb.Media
	9,  // 3: moviepb.Movie.tmdb_info:type_name -> moviepb.TMDBInfo
	4,  // 4: moviepb.MovieList.list:type_name -> moviepb.Movie
	12, // 5: moviepb.TMDBInfo.data:type_name -> google.protobuf.Struct
	11, // 6: moviepb.Media.created_at:type_name -> google.protobuf.Timestamp
	11, // 7: moviepb.Media.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 8: moviepb.Media.type:type_name -> moviepb.MediaType
	1,  // 9: moviepb.MovieService.CreateMovie:input_type -> moviepb.CreateMovieRequest
	3,  // 10: moviepb.MovieService.GetMovieByID:input_type -> moviepb.GetMovieByIDRequest
	5,  // 11: moviepb.MovieService.ListMovies:input_type -> moviepb.ListMoviesRequest
	7,  // 12: moviepb.MovieService.UpdateMovieByID:input_type -> moviepb.UpdateMovieByIDRequest
	8,  // 13: moviepb.MovieService.DeleteMovieByID:input_type -> moviepb.DeleteMovieByIDRequest
	2,  // 14: moviepb.MovieService.CreateMovie:output_type -> moviepb.CreateMovieResponse
	4,  // 15: moviepb.MovieService.GetMovieByID:output_type -> moviepb.Movie
	6,  // 16: moviepb.MovieService.ListMovies:output_type -> moviepb.MovieList
	13, // 17: moviepb.MovieService.UpdateMovieByID:output_type -> google.protobuf.Empty
	13, // 18: moviepb.MovieService.DeleteMovieByID:output_type -> google.protobuf.Empty
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_movie_service_proto_init() }
func file_movie_service_proto_init() {
	if File_movie_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_movie_service_proto_rawDesc), len(file_movie_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_service_proto_goTypes,
		DependencyIndexes: file_movie_service_proto_depIdxs,
		EnumInfos:         file_movie_service_proto_enumTypes,
		MessageInfos:      file_movie_service_proto_msgTypes,
	}.Build()
	File_movie_service_proto = out.File
	file_movie_service_proto_goTypes = nil
	file_movie_service_proto_depIdxs = nil
}
