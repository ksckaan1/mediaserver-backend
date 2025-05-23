// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: episode_service.proto

package episodepb

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

type CreateEpisodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SeasonId      string                 `protobuf:"bytes,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MediaId       string                 `protobuf:"bytes,4,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEpisodeRequest) Reset() {
	*x = CreateEpisodeRequest{}
	mi := &file_episode_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEpisodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEpisodeRequest) ProtoMessage() {}

func (x *CreateEpisodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEpisodeRequest.ProtoReflect.Descriptor instead.
func (*CreateEpisodeRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEpisodeRequest) GetSeasonId() string {
	if x != nil {
		return x.SeasonId
	}
	return ""
}

func (x *CreateEpisodeRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEpisodeRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEpisodeRequest) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

type CreateEpisodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EpisodeId     string                 `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEpisodeResponse) Reset() {
	*x = CreateEpisodeResponse{}
	mi := &file_episode_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEpisodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEpisodeResponse) ProtoMessage() {}

func (x *CreateEpisodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEpisodeResponse.ProtoReflect.Descriptor instead.
func (*CreateEpisodeResponse) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEpisodeResponse) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

type GetEpisodeByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EpisodeId     string                 `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEpisodeByIDRequest) Reset() {
	*x = GetEpisodeByIDRequest{}
	mi := &file_episode_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEpisodeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEpisodeByIDRequest) ProtoMessage() {}

func (x *GetEpisodeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEpisodeByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEpisodeByIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetEpisodeByIDRequest) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

type Episode struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	SeasonId      string                 `protobuf:"bytes,6,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	MediaInfo     *Media                 `protobuf:"bytes,7,opt,name=media_info,json=mediaInfo,proto3" json:"media_info,omitempty"`
	Order         int32                  `protobuf:"varint,8,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Episode) Reset() {
	*x = Episode{}
	mi := &file_episode_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Episode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episode) ProtoMessage() {}

func (x *Episode) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episode.ProtoReflect.Descriptor instead.
func (*Episode) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{3}
}

func (x *Episode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Episode) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Episode) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Episode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Episode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Episode) GetSeasonId() string {
	if x != nil {
		return x.SeasonId
	}
	return ""
}

func (x *Episode) GetMediaInfo() *Media {
	if x != nil {
		return x.MediaInfo
	}
	return nil
}

func (x *Episode) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type ListEpisodesBySeasonIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SeasonId      string                 `protobuf:"bytes,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEpisodesBySeasonIDRequest) Reset() {
	*x = ListEpisodesBySeasonIDRequest{}
	mi := &file_episode_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEpisodesBySeasonIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEpisodesBySeasonIDRequest) ProtoMessage() {}

func (x *ListEpisodesBySeasonIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEpisodesBySeasonIDRequest.ProtoReflect.Descriptor instead.
func (*ListEpisodesBySeasonIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListEpisodesBySeasonIDRequest) GetSeasonId() string {
	if x != nil {
		return x.SeasonId
	}
	return ""
}

type EpisodeList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*Episode             `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EpisodeList) Reset() {
	*x = EpisodeList{}
	mi := &file_episode_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EpisodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpisodeList) ProtoMessage() {}

func (x *EpisodeList) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpisodeList.ProtoReflect.Descriptor instead.
func (*EpisodeList) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{5}
}

func (x *EpisodeList) GetList() []*Episode {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateEpisodeByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EpisodeId     string                 `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MediaId       string                 `protobuf:"bytes,4,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEpisodeByIDRequest) Reset() {
	*x = UpdateEpisodeByIDRequest{}
	mi := &file_episode_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEpisodeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEpisodeByIDRequest) ProtoMessage() {}

func (x *UpdateEpisodeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEpisodeByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateEpisodeByIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEpisodeByIDRequest) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *UpdateEpisodeByIDRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateEpisodeByIDRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEpisodeByIDRequest) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

type ReorderEpisodesBySeasonIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SeasonId      string                 `protobuf:"bytes,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	EpisodeIds    []string               `protobuf:"bytes,2,rep,name=episode_ids,json=episodeIds,proto3" json:"episode_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReorderEpisodesBySeasonIDRequest) Reset() {
	*x = ReorderEpisodesBySeasonIDRequest{}
	mi := &file_episode_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReorderEpisodesBySeasonIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReorderEpisodesBySeasonIDRequest) ProtoMessage() {}

func (x *ReorderEpisodesBySeasonIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReorderEpisodesBySeasonIDRequest.ProtoReflect.Descriptor instead.
func (*ReorderEpisodesBySeasonIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{7}
}

func (x *ReorderEpisodesBySeasonIDRequest) GetSeasonId() string {
	if x != nil {
		return x.SeasonId
	}
	return ""
}

func (x *ReorderEpisodesBySeasonIDRequest) GetEpisodeIds() []string {
	if x != nil {
		return x.EpisodeIds
	}
	return nil
}

type DeleteAllEpisodesBySeasonIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SeasonId      string                 `protobuf:"bytes,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllEpisodesBySeasonIDRequest) Reset() {
	*x = DeleteAllEpisodesBySeasonIDRequest{}
	mi := &file_episode_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllEpisodesBySeasonIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllEpisodesBySeasonIDRequest) ProtoMessage() {}

func (x *DeleteAllEpisodesBySeasonIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllEpisodesBySeasonIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllEpisodesBySeasonIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAllEpisodesBySeasonIDRequest) GetSeasonId() string {
	if x != nil {
		return x.SeasonId
	}
	return ""
}

type DeleteEpisodeByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EpisodeId     string                 `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEpisodeByIDRequest) Reset() {
	*x = DeleteEpisodeByIDRequest{}
	mi := &file_episode_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEpisodeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEpisodeByIDRequest) ProtoMessage() {}

func (x *DeleteEpisodeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEpisodeByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteEpisodeByIDRequest) Descriptor() ([]byte, []int) {
	return file_episode_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEpisodeByIDRequest) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

type Media struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Path          string                 `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Type          string                 `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	MimeType      string                 `protobuf:"bytes,7,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size          int64                  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Media) Reset() {
	*x = Media{}
	mi := &file_episode_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_episode_service_proto_msgTypes[10]
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
	return file_episode_service_proto_rawDescGZIP(), []int{10}
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

func (x *Media) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
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

var File_episode_service_proto protoreflect.FileDescriptor

const file_episode_service_proto_rawDesc = "" +
	"\n" +
	"\x15episode_service.proto\x12\tepisodepb\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x86\x01\n" +
	"\x14CreateEpisodeRequest\x12\x1b\n" +
	"\tseason_id\x18\x01 \x01(\tR\bseasonId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x19\n" +
	"\bmedia_id\x18\x04 \x01(\tR\amediaId\"6\n" +
	"\x15CreateEpisodeResponse\x12\x1d\n" +
	"\n" +
	"episode_id\x18\x01 \x01(\tR\tepisodeId\"6\n" +
	"\x15GetEpisodeByIDRequest\x12\x1d\n" +
	"\n" +
	"episode_id\x18\x01 \x01(\tR\tepisodeId\"\xab\x02\n" +
	"\aEpisode\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12\x1b\n" +
	"\tseason_id\x18\x06 \x01(\tR\bseasonId\x12/\n" +
	"\n" +
	"media_info\x18\a \x01(\v2\x10.episodepb.MediaR\tmediaInfo\x12\x14\n" +
	"\x05order\x18\b \x01(\x05R\x05order\"<\n" +
	"\x1dListEpisodesBySeasonIDRequest\x12\x1b\n" +
	"\tseason_id\x18\x01 \x01(\tR\bseasonId\"5\n" +
	"\vEpisodeList\x12&\n" +
	"\x04list\x18\x01 \x03(\v2\x12.episodepb.EpisodeR\x04list\"\x8c\x01\n" +
	"\x18UpdateEpisodeByIDRequest\x12\x1d\n" +
	"\n" +
	"episode_id\x18\x01 \x01(\tR\tepisodeId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x19\n" +
	"\bmedia_id\x18\x04 \x01(\tR\amediaId\"`\n" +
	" ReorderEpisodesBySeasonIDRequest\x12\x1b\n" +
	"\tseason_id\x18\x01 \x01(\tR\bseasonId\x12\x1f\n" +
	"\vepisode_ids\x18\x02 \x03(\tR\n" +
	"episodeIds\"A\n" +
	"\"DeleteAllEpisodesBySeasonIDRequest\x12\x1b\n" +
	"\tseason_id\x18\x01 \x01(\tR\bseasonId\"9\n" +
	"\x18DeleteEpisodeByIDRequest\x12\x1d\n" +
	"\n" +
	"episode_id\x18\x01 \x01(\tR\tepisodeId\"\xfc\x01\n" +
	"\x05Media\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12\x12\n" +
	"\x04path\x18\x05 \x01(\tR\x04path\x12\x12\n" +
	"\x04type\x18\x06 \x01(\tR\x04type\x12\x1b\n" +
	"\tmime_type\x18\a \x01(\tR\bmimeType\x12\x12\n" +
	"\x04size\x18\b \x01(\x03R\x04size2\xf4\x04\n" +
	"\x0eEpisodeService\x12R\n" +
	"\rCreateEpisode\x12\x1f.episodepb.CreateEpisodeRequest\x1a .episodepb.CreateEpisodeResponse\x12F\n" +
	"\x0eGetEpisodeByID\x12 .episodepb.GetEpisodeByIDRequest\x1a\x12.episodepb.Episode\x12Z\n" +
	"\x16ListEpisodesBySeasonID\x12(.episodepb.ListEpisodesBySeasonIDRequest\x1a\x16.episodepb.EpisodeList\x12P\n" +
	"\x11UpdateEpisodeByID\x12#.episodepb.UpdateEpisodeByIDRequest\x1a\x16.google.protobuf.Empty\x12`\n" +
	"\x19ReorderEpisodesBySeasonID\x12+.episodepb.ReorderEpisodesBySeasonIDRequest\x1a\x16.google.protobuf.Empty\x12P\n" +
	"\x11DeleteEpisodeByID\x12#.episodepb.DeleteEpisodeByIDRequest\x1a\x16.google.protobuf.Empty\x12d\n" +
	"\x1bDeleteAllEpisodesBySeasonID\x12-.episodepb.DeleteAllEpisodesBySeasonIDRequest\x1a\x16.google.protobuf.EmptyB\x1fZ\x1dshared/pb/episodepb;episodepbb\x06proto3"

var (
	file_episode_service_proto_rawDescOnce sync.Once
	file_episode_service_proto_rawDescData []byte
)

func file_episode_service_proto_rawDescGZIP() []byte {
	file_episode_service_proto_rawDescOnce.Do(func() {
		file_episode_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_episode_service_proto_rawDesc), len(file_episode_service_proto_rawDesc)))
	})
	return file_episode_service_proto_rawDescData
}

var file_episode_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_episode_service_proto_goTypes = []any{
	(*CreateEpisodeRequest)(nil),               // 0: episodepb.CreateEpisodeRequest
	(*CreateEpisodeResponse)(nil),              // 1: episodepb.CreateEpisodeResponse
	(*GetEpisodeByIDRequest)(nil),              // 2: episodepb.GetEpisodeByIDRequest
	(*Episode)(nil),                            // 3: episodepb.Episode
	(*ListEpisodesBySeasonIDRequest)(nil),      // 4: episodepb.ListEpisodesBySeasonIDRequest
	(*EpisodeList)(nil),                        // 5: episodepb.EpisodeList
	(*UpdateEpisodeByIDRequest)(nil),           // 6: episodepb.UpdateEpisodeByIDRequest
	(*ReorderEpisodesBySeasonIDRequest)(nil),   // 7: episodepb.ReorderEpisodesBySeasonIDRequest
	(*DeleteAllEpisodesBySeasonIDRequest)(nil), // 8: episodepb.DeleteAllEpisodesBySeasonIDRequest
	(*DeleteEpisodeByIDRequest)(nil),           // 9: episodepb.DeleteEpisodeByIDRequest
	(*Media)(nil),                              // 10: episodepb.Media
	(*timestamppb.Timestamp)(nil),              // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                      // 12: google.protobuf.Empty
}
var file_episode_service_proto_depIdxs = []int32{
	11, // 0: episodepb.Episode.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: episodepb.Episode.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: episodepb.Episode.media_info:type_name -> episodepb.Media
	3,  // 3: episodepb.EpisodeList.list:type_name -> episodepb.Episode
	11, // 4: episodepb.Media.created_at:type_name -> google.protobuf.Timestamp
	11, // 5: episodepb.Media.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 6: episodepb.EpisodeService.CreateEpisode:input_type -> episodepb.CreateEpisodeRequest
	2,  // 7: episodepb.EpisodeService.GetEpisodeByID:input_type -> episodepb.GetEpisodeByIDRequest
	4,  // 8: episodepb.EpisodeService.ListEpisodesBySeasonID:input_type -> episodepb.ListEpisodesBySeasonIDRequest
	6,  // 9: episodepb.EpisodeService.UpdateEpisodeByID:input_type -> episodepb.UpdateEpisodeByIDRequest
	7,  // 10: episodepb.EpisodeService.ReorderEpisodesBySeasonID:input_type -> episodepb.ReorderEpisodesBySeasonIDRequest
	9,  // 11: episodepb.EpisodeService.DeleteEpisodeByID:input_type -> episodepb.DeleteEpisodeByIDRequest
	8,  // 12: episodepb.EpisodeService.DeleteAllEpisodesBySeasonID:input_type -> episodepb.DeleteAllEpisodesBySeasonIDRequest
	1,  // 13: episodepb.EpisodeService.CreateEpisode:output_type -> episodepb.CreateEpisodeResponse
	3,  // 14: episodepb.EpisodeService.GetEpisodeByID:output_type -> episodepb.Episode
	5,  // 15: episodepb.EpisodeService.ListEpisodesBySeasonID:output_type -> episodepb.EpisodeList
	12, // 16: episodepb.EpisodeService.UpdateEpisodeByID:output_type -> google.protobuf.Empty
	12, // 17: episodepb.EpisodeService.ReorderEpisodesBySeasonID:output_type -> google.protobuf.Empty
	12, // 18: episodepb.EpisodeService.DeleteEpisodeByID:output_type -> google.protobuf.Empty
	12, // 19: episodepb.EpisodeService.DeleteAllEpisodesBySeasonID:output_type -> google.protobuf.Empty
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_episode_service_proto_init() }
func file_episode_service_proto_init() {
	if File_episode_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_episode_service_proto_rawDesc), len(file_episode_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_episode_service_proto_goTypes,
		DependencyIndexes: file_episode_service_proto_depIdxs,
		MessageInfos:      file_episode_service_proto_msgTypes,
	}.Build()
	File_episode_service_proto = out.File
	file_episode_service_proto_goTypes = nil
	file_episode_service_proto_depIdxs = nil
}
