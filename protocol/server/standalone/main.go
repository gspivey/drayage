package main

import (
	"log"
	"net"

	pb "github.com/drayage/protocol"
	"github.com/drayage/protocol/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	commServer := grpc.NewServer()
	pb.RegisterCommsProtoServer(commServer, &server.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(commServer)
	if err := commServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
