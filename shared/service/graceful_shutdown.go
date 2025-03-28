package service

import (
	"context"
	"net"
	"os"
	"os/signal"
	"shared/ports"
	"syscall"

	"google.golang.org/grpc"
)

func handleGracefulShutdown(server *grpc.Server, listener net.Listener, lg ports.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	lg.Info(context.Background(), "shutting down gracefully")
	server.GracefulStop()
	listener.Close()
}
