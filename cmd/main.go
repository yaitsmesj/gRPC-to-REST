package main

import (
	"log"
	"net"

	pb "github.com/yaitsmesj/gRPC-to-REST/proto"
	"github.com/yaitsmesj/gRPC-to-REST/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	s := server.UserServer{}
	pb.RegisterUserServiceServer(grpcServer, &s)

	log.Println("Starting Server...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
