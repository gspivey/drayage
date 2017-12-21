package main

import (
	"context"
	"log"
	"net"

	pb "github.com/drayage/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) RMVolume(ctx context.Context, volumeName *pb.Volume) (*pb.Status, error) {
	// docker VolumeRemove
	return &pb.Status{"success"}, nil
}

func (s *server) LSVolume(volumeName *pb.Volume, stream pb.CommsProto_LSVolumeServer) error {
	// Docker VolumeList
	//    VolumeInspect each volume within the list
	// return volume info found from docker API volume inspect
	/*
		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	*/
	return nil

}

func (s *server) AddVolume(ctx context.Context, volumeName *pb.Volume) (*pb.Status, error) {
	// Docker VolumeCreate
	return &pb.Status{"success"}, nil
}

func (s *server) GetVolume(ctx context.Context, volumeAndHost *pb.VolumeAndHost) (*pb.Status, error) {
	// Drayage volume streamer
	return &pb.Status{"complete"}, nil
}

func (s *server) VolumeFiles(volumeName *pb.Volume, stream pb.CommsProto_VolumeFilesServer) error {
	// Inspect Docker Volume VolumeInspect
	// Use root path with drayage GetFilesPathSize
	// stream files name back
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	commServer := grpc.NewServer()
	pb.RegisterCommsProtoServer(commServer, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(commServer)
	if err := commServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
