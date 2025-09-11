package main

import (
	"context"
	"log"
	"net"

	echov1 "github.com/lorette-m/GoLearn/gRPC_Learn/gen/go/echo/v1"
	"google.golang.org/grpc"
)

type server struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	log.Printf("Получен запрос: %s", req.GetMessage())
	return &echov1.EchoResponse{Message: "Эхо: " + req.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	s := grpc.NewServer()
	echov1.RegisterEchoServiceServer(s, &server{})

	log.Println("gRPC сервер на :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
