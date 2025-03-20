package main

import (
	"context"
	"movie_service/internal/port"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type loggerMiddleWare struct {
	logger port.Logger
}

func (l *loggerMiddleWare) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	resp, err := handler(ctx, req)
	statusCode := codes.OK
	if err != nil {
		statusCode = status.Code(err)
	}

	l.logger.Info(ctx, "request accepted",
		"method", info.FullMethod,
		"duration", time.Since(startTime).String(),
		"status", statusCode,
		"metadata", md,
		"request", req,
	)
	return resp, err
}

// LoggingStreamInterceptor, gelen gRPC stream isteklerini loglamak için bir stream interceptor
func (l *loggerMiddleWare) streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	startTime := time.Now()
	ctx := ss.Context()
	md, _ := metadata.FromIncomingContext(ctx)
	wrappedStream := &wrappedServerStream{
		ServerStream: ss,
		info:         info,
	}
	err := handler(srv, wrappedStream)

	statusCode := codes.OK
	if err != nil {
		statusCode = status.Code(err)
	}
	l.logger.Info(ctx, "request accepted",
		"method", info.FullMethod,
		"duration", time.Since(startTime).String(),
		"status", statusCode,
		"metadata", md,
	)
	return err
}

type wrappedServerStream struct {
	grpc.ServerStream
	info *grpc.StreamServerInfo
}
