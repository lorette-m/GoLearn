package server

import (
	"context"

	echov1 "github.com/lorette-m/gRPC_Learn/gen/go/echo/v1"
	_ "google.golang.org/grpc"
)

type server struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echov1.EchoRequest)
