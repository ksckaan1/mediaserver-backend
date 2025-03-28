// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: season_service.proto

package seasonpb

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
	SeasonService_CreateSeason_FullMethodName               = "/seasonpb.SeasonService/CreateSeason"
	SeasonService_GetSeasonByID_FullMethodName              = "/seasonpb.SeasonService/GetSeasonByID"
	SeasonService_ListSeasonsBySeriesID_FullMethodName      = "/seasonpb.SeasonService/ListSeasonsBySeriesID"
	SeasonService_UpdateSeasonByID_FullMethodName           = "/seasonpb.SeasonService/UpdateSeasonByID"
	SeasonService_ReorderSeasonsBySeriesID_FullMethodName   = "/seasonpb.SeasonService/ReorderSeasonsBySeriesID"
	SeasonService_DeleteSeasonByID_FullMethodName           = "/seasonpb.SeasonService/DeleteSeasonByID"
	SeasonService_DeleteAllSeasonsBySeriesID_FullMethodName = "/seasonpb.SeasonService/DeleteAllSeasonsBySeriesID"
)

// SeasonServiceClient is the client API for SeasonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeasonServiceClient interface {
	CreateSeason(ctx context.Context, in *CreateSeasonRequest, opts ...grpc.CallOption) (*CreateSeasonResponse, error)
	GetSeasonByID(ctx context.Context, in *GetSeasonByIDRequest, opts ...grpc.CallOption) (*Season, error)
	ListSeasonsBySeriesID(ctx context.Context, in *ListSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*SeasonList, error)
	UpdateSeasonByID(ctx context.Context, in *UpdateSeasonByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReorderSeasonsBySeriesID(ctx context.Context, in *ReorderSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSeasonByID(ctx context.Context, in *DeleteSeasonByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAllSeasonsBySeriesID(ctx context.Context, in *DeleteAllSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type seasonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSeasonServiceClient(cc grpc.ClientConnInterface) SeasonServiceClient {
	return &seasonServiceClient{cc}
}

func (c *seasonServiceClient) CreateSeason(ctx context.Context, in *CreateSeasonRequest, opts ...grpc.CallOption) (*CreateSeasonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSeasonResponse)
	err := c.cc.Invoke(ctx, SeasonService_CreateSeason_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) GetSeasonByID(ctx context.Context, in *GetSeasonByIDRequest, opts ...grpc.CallOption) (*Season, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Season)
	err := c.cc.Invoke(ctx, SeasonService_GetSeasonByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) ListSeasonsBySeriesID(ctx context.Context, in *ListSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*SeasonList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SeasonList)
	err := c.cc.Invoke(ctx, SeasonService_ListSeasonsBySeriesID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) UpdateSeasonByID(ctx context.Context, in *UpdateSeasonByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SeasonService_UpdateSeasonByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) ReorderSeasonsBySeriesID(ctx context.Context, in *ReorderSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SeasonService_ReorderSeasonsBySeriesID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) DeleteSeasonByID(ctx context.Context, in *DeleteSeasonByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SeasonService_DeleteSeasonByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seasonServiceClient) DeleteAllSeasonsBySeriesID(ctx context.Context, in *DeleteAllSeasonsBySeriesIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SeasonService_DeleteAllSeasonsBySeriesID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeasonServiceServer is the server API for SeasonService service.
// All implementations must embed UnimplementedSeasonServiceServer
// for forward compatibility.
type SeasonServiceServer interface {
	CreateSeason(context.Context, *CreateSeasonRequest) (*CreateSeasonResponse, error)
	GetSeasonByID(context.Context, *GetSeasonByIDRequest) (*Season, error)
	ListSeasonsBySeriesID(context.Context, *ListSeasonsBySeriesIDRequest) (*SeasonList, error)
	UpdateSeasonByID(context.Context, *UpdateSeasonByIDRequest) (*emptypb.Empty, error)
	ReorderSeasonsBySeriesID(context.Context, *ReorderSeasonsBySeriesIDRequest) (*emptypb.Empty, error)
	DeleteSeasonByID(context.Context, *DeleteSeasonByIDRequest) (*emptypb.Empty, error)
	DeleteAllSeasonsBySeriesID(context.Context, *DeleteAllSeasonsBySeriesIDRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSeasonServiceServer()
}

// UnimplementedSeasonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSeasonServiceServer struct{}

func (UnimplementedSeasonServiceServer) CreateSeason(context.Context, *CreateSeasonRequest) (*CreateSeasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeason not implemented")
}
func (UnimplementedSeasonServiceServer) GetSeasonByID(context.Context, *GetSeasonByIDRequest) (*Season, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeasonByID not implemented")
}
func (UnimplementedSeasonServiceServer) ListSeasonsBySeriesID(context.Context, *ListSeasonsBySeriesIDRequest) (*SeasonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSeasonsBySeriesID not implemented")
}
func (UnimplementedSeasonServiceServer) UpdateSeasonByID(context.Context, *UpdateSeasonByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeasonByID not implemented")
}
func (UnimplementedSeasonServiceServer) ReorderSeasonsBySeriesID(context.Context, *ReorderSeasonsBySeriesIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderSeasonsBySeriesID not implemented")
}
func (UnimplementedSeasonServiceServer) DeleteSeasonByID(context.Context, *DeleteSeasonByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSeasonByID not implemented")
}
func (UnimplementedSeasonServiceServer) DeleteAllSeasonsBySeriesID(context.Context, *DeleteAllSeasonsBySeriesIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllSeasonsBySeriesID not implemented")
}
func (UnimplementedSeasonServiceServer) mustEmbedUnimplementedSeasonServiceServer() {}
func (UnimplementedSeasonServiceServer) testEmbeddedByValue()                       {}

// UnsafeSeasonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeasonServiceServer will
// result in compilation errors.
type UnsafeSeasonServiceServer interface {
	mustEmbedUnimplementedSeasonServiceServer()
}

func RegisterSeasonServiceServer(s grpc.ServiceRegistrar, srv SeasonServiceServer) {
	// If the following call pancis, it indicates UnimplementedSeasonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SeasonService_ServiceDesc, srv)
}

func _SeasonService_CreateSeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).CreateSeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_CreateSeason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).CreateSeason(ctx, req.(*CreateSeasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_GetSeasonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeasonByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).GetSeasonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_GetSeasonByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).GetSeasonByID(ctx, req.(*GetSeasonByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_ListSeasonsBySeriesID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSeasonsBySeriesIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).ListSeasonsBySeriesID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_ListSeasonsBySeriesID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).ListSeasonsBySeriesID(ctx, req.(*ListSeasonsBySeriesIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_UpdateSeasonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSeasonByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).UpdateSeasonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_UpdateSeasonByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).UpdateSeasonByID(ctx, req.(*UpdateSeasonByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_ReorderSeasonsBySeriesID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderSeasonsBySeriesIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).ReorderSeasonsBySeriesID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_ReorderSeasonsBySeriesID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).ReorderSeasonsBySeriesID(ctx, req.(*ReorderSeasonsBySeriesIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_DeleteSeasonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSeasonByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).DeleteSeasonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_DeleteSeasonByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).DeleteSeasonByID(ctx, req.(*DeleteSeasonByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeasonService_DeleteAllSeasonsBySeriesID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllSeasonsBySeriesIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeasonServiceServer).DeleteAllSeasonsBySeriesID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeasonService_DeleteAllSeasonsBySeriesID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeasonServiceServer).DeleteAllSeasonsBySeriesID(ctx, req.(*DeleteAllSeasonsBySeriesIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SeasonService_ServiceDesc is the grpc.ServiceDesc for SeasonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeasonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seasonpb.SeasonService",
	HandlerType: (*SeasonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSeason",
			Handler:    _SeasonService_CreateSeason_Handler,
		},
		{
			MethodName: "GetSeasonByID",
			Handler:    _SeasonService_GetSeasonByID_Handler,
		},
		{
			MethodName: "ListSeasonsBySeriesID",
			Handler:    _SeasonService_ListSeasonsBySeriesID_Handler,
		},
		{
			MethodName: "UpdateSeasonByID",
			Handler:    _SeasonService_UpdateSeasonByID_Handler,
		},
		{
			MethodName: "ReorderSeasonsBySeriesID",
			Handler:    _SeasonService_ReorderSeasonsBySeriesID_Handler,
		},
		{
			MethodName: "DeleteSeasonByID",
			Handler:    _SeasonService_DeleteSeasonByID_Handler,
		},
		{
			MethodName: "DeleteAllSeasonsBySeriesID",
			Handler:    _SeasonService_DeleteAllSeasonsBySeriesID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "season_service.proto",
}
