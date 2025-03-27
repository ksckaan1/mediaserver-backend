// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: episode_service.proto

package episodepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EpisodeService_CreateEpisode_FullMethodName               = "/episodepb.EpisodeService/CreateEpisode"
	EpisodeService_GetEpisodeByID_FullMethodName              = "/episodepb.EpisodeService/GetEpisodeByID"
	EpisodeService_ListEpisodesBySeasonID_FullMethodName      = "/episodepb.EpisodeService/ListEpisodesBySeasonID"
	EpisodeService_UpdateEpisodeByID_FullMethodName           = "/episodepb.EpisodeService/UpdateEpisodeByID"
	EpisodeService_ReorderEpisodesBySeasonID_FullMethodName   = "/episodepb.EpisodeService/ReorderEpisodesBySeasonID"
	EpisodeService_DeleteEpisodeByID_FullMethodName           = "/episodepb.EpisodeService/DeleteEpisodeByID"
	EpisodeService_DeleteAllEpisodesBySeasonID_FullMethodName = "/episodepb.EpisodeService/DeleteAllEpisodesBySeasonID"
)

// EpisodeServiceClient is the client API for EpisodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EpisodeServiceClient interface {
	CreateEpisode(ctx context.Context, in *CreateEpisodeRequest, opts ...grpc.CallOption) (*CreateEpisodeResponse, error)
	GetEpisodeByID(ctx context.Context, in *GetEpisodeByIDRequest, opts ...grpc.CallOption) (*Episode, error)
	ListEpisodesBySeasonID(ctx context.Context, in *ListEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*EpisodeList, error)
	UpdateEpisodeByID(ctx context.Context, in *UpdateEpisodeByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReorderEpisodesBySeasonID(ctx context.Context, in *ReorderEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEpisodeByID(ctx context.Context, in *DeleteEpisodeByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAllEpisodesBySeasonID(ctx context.Context, in *DeleteAllEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type episodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEpisodeServiceClient(cc grpc.ClientConnInterface) EpisodeServiceClient {
	return &episodeServiceClient{cc}
}

func (c *episodeServiceClient) CreateEpisode(ctx context.Context, in *CreateEpisodeRequest, opts ...grpc.CallOption) (*CreateEpisodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEpisodeResponse)
	err := c.cc.Invoke(ctx, EpisodeService_CreateEpisode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) GetEpisodeByID(ctx context.Context, in *GetEpisodeByIDRequest, opts ...grpc.CallOption) (*Episode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Episode)
	err := c.cc.Invoke(ctx, EpisodeService_GetEpisodeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) ListEpisodesBySeasonID(ctx context.Context, in *ListEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*EpisodeList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EpisodeList)
	err := c.cc.Invoke(ctx, EpisodeService_ListEpisodesBySeasonID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) UpdateEpisodeByID(ctx context.Context, in *UpdateEpisodeByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EpisodeService_UpdateEpisodeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) ReorderEpisodesBySeasonID(ctx context.Context, in *ReorderEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EpisodeService_ReorderEpisodesBySeasonID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) DeleteEpisodeByID(ctx context.Context, in *DeleteEpisodeByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EpisodeService_DeleteEpisodeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) DeleteAllEpisodesBySeasonID(ctx context.Context, in *DeleteAllEpisodesBySeasonIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EpisodeService_DeleteAllEpisodesBySeasonID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EpisodeServiceServer is the server API for EpisodeService service.
// All implementations must embed UnimplementedEpisodeServiceServer
// for forward compatibility.
type EpisodeServiceServer interface {
	CreateEpisode(context.Context, *CreateEpisodeRequest) (*CreateEpisodeResponse, error)
	GetEpisodeByID(context.Context, *GetEpisodeByIDRequest) (*Episode, error)
	ListEpisodesBySeasonID(context.Context, *ListEpisodesBySeasonIDRequest) (*EpisodeList, error)
	UpdateEpisodeByID(context.Context, *UpdateEpisodeByIDRequest) (*emptypb.Empty, error)
	ReorderEpisodesBySeasonID(context.Context, *ReorderEpisodesBySeasonIDRequest) (*emptypb.Empty, error)
	DeleteEpisodeByID(context.Context, *DeleteEpisodeByIDRequest) (*emptypb.Empty, error)
	DeleteAllEpisodesBySeasonID(context.Context, *DeleteAllEpisodesBySeasonIDRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEpisodeServiceServer()
}

// UnimplementedEpisodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEpisodeServiceServer struct{}

func (UnimplementedEpisodeServiceServer) CreateEpisode(context.Context, *CreateEpisodeRequest) (*CreateEpisodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEpisode not implemented")
}
func (UnimplementedEpisodeServiceServer) GetEpisodeByID(context.Context, *GetEpisodeByIDRequest) (*Episode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEpisodeByID not implemented")
}
func (UnimplementedEpisodeServiceServer) ListEpisodesBySeasonID(context.Context, *ListEpisodesBySeasonIDRequest) (*EpisodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEpisodesBySeasonID not implemented")
}
func (UnimplementedEpisodeServiceServer) UpdateEpisodeByID(context.Context, *UpdateEpisodeByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEpisodeByID not implemented")
}
func (UnimplementedEpisodeServiceServer) ReorderEpisodesBySeasonID(context.Context, *ReorderEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderEpisodesBySeasonID not implemented")
}
func (UnimplementedEpisodeServiceServer) DeleteEpisodeByID(context.Context, *DeleteEpisodeByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEpisodeByID not implemented")
}
func (UnimplementedEpisodeServiceServer) DeleteAllEpisodesBySeasonID(context.Context, *DeleteAllEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllEpisodesBySeasonID not implemented")
}
func (UnimplementedEpisodeServiceServer) mustEmbedUnimplementedEpisodeServiceServer() {}
func (UnimplementedEpisodeServiceServer) testEmbeddedByValue()                        {}

// UnsafeEpisodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EpisodeServiceServer will
// result in compilation errors.
type UnsafeEpisodeServiceServer interface {
	mustEmbedUnimplementedEpisodeServiceServer()
}

func RegisterEpisodeServiceServer(s grpc.ServiceRegistrar, srv EpisodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedEpisodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EpisodeService_ServiceDesc, srv)
}

func _EpisodeService_CreateEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).CreateEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_CreateEpisode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).CreateEpisode(ctx, req.(*CreateEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_GetEpisodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEpisodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).GetEpisodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_GetEpisodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).GetEpisodeByID(ctx, req.(*GetEpisodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_ListEpisodesBySeasonID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEpisodesBySeasonIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).ListEpisodesBySeasonID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_ListEpisodesBySeasonID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).ListEpisodesBySeasonID(ctx, req.(*ListEpisodesBySeasonIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_UpdateEpisodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEpisodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).UpdateEpisodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_UpdateEpisodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).UpdateEpisodeByID(ctx, req.(*UpdateEpisodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_ReorderEpisodesBySeasonID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderEpisodesBySeasonIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).ReorderEpisodesBySeasonID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_ReorderEpisodesBySeasonID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).ReorderEpisodesBySeasonID(ctx, req.(*ReorderEpisodesBySeasonIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_DeleteEpisodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEpisodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).DeleteEpisodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_DeleteEpisodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).DeleteEpisodeByID(ctx, req.(*DeleteEpisodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_DeleteAllEpisodesBySeasonID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllEpisodesBySeasonIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).DeleteAllEpisodesBySeasonID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpisodeService_DeleteAllEpisodesBySeasonID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).DeleteAllEpisodesBySeasonID(ctx, req.(*DeleteAllEpisodesBySeasonIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EpisodeService_ServiceDesc is the grpc.ServiceDesc for EpisodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EpisodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "episodepb.EpisodeService",
	HandlerType: (*EpisodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEpisode",
			Handler:    _EpisodeService_CreateEpisode_Handler,
		},
		{
			MethodName: "GetEpisodeByID",
			Handler:    _EpisodeService_GetEpisodeByID_Handler,
		},
		{
			MethodName: "ListEpisodesBySeasonID",
			Handler:    _EpisodeService_ListEpisodesBySeasonID_Handler,
		},
		{
			MethodName: "UpdateEpisodeByID",
			Handler:    _EpisodeService_UpdateEpisodeByID_Handler,
		},
		{
			MethodName: "ReorderEpisodesBySeasonID",
			Handler:    _EpisodeService_ReorderEpisodesBySeasonID_Handler,
		},
		{
			MethodName: "DeleteEpisodeByID",
			Handler:    _EpisodeService_DeleteEpisodeByID_Handler,
		},
		{
			MethodName: "DeleteAllEpisodesBySeasonID",
			Handler:    _EpisodeService_DeleteAllEpisodesBySeasonID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "episode_service.proto",
}
