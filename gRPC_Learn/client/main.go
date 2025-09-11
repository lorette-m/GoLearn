package main

import (
	"context"
	"log"
	"time"

	echov1 "github.com/lorette-m/GoLearn/gRPC_Learn/gen/go/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // для продакшена используйте TLS
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	client := echov1.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &echov1.EchoRequest{Message: "Привет, gRPC!"})
	if err != nil {
		log.Fatalf("Echo error: %v", err)
	}

	log.Printf("Ответ сервера: %s", resp.GetMessage())
}
