package main

import (
	"flag"
	"log"

	pb "github.com/drayage/protocol"
	clientpackage "github.com/drayage/protocol/client"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCommsProtoClient(conn)

	v := &pb.Volume{"test"}
	vh := &pb.VolumeAndHost{"test", "host.name.domain"}
	clientpackage.AddVolume(client, v)
	clientpackage.GetVolume(client, vh)
	clientpackage.LSVolume(client, v)
	clientpackage.RMVolume(client, v)
	clientpackage.VolumeFiles(client, v)

}
