package main

import (
	"flag"
	"log"

	pb "github.com/gspivey/drayage/protocol"
	clientpackage "github.com/gspivey/drayage/protocol/client"
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

	vt := &pb.Volume{"test"}
	v := &pb.Volume{"concourse-worker-pki"}
	vh := &pb.VolumeAndHost{"test", "host.name.domain"}
	clientpackage.AddVolume(client, vt)
	clientpackage.GetVolume(client, vh)
	clientpackage.LSVolume(client, v)
	clientpackage.RMVolume(client, vt)
	clientpackage.ProcesVolumeFiles(client, v)

}
