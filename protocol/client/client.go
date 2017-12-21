package client

import (
	"context"
	"log"

	pb "github.com/drayage/protocol"
)

func RMVolume(client pb.CommsProtoClient, v *pb.Volume) {
	status, err := client.RMVolume(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("RMVolume: %s", status)

}

func LSVolume(client pb.CommsProtoClient, v *pb.Volume) {
	status, err := client.LSVolume(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("LSVolume: %s", status)

}

func AddVolume(client pb.CommsProtoClient, v *pb.Volume) {
	status, err := client.AddVolume(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("AddVolume: %s", status)

}

func GetVolume(client pb.CommsProtoClient, v *pb.VolumeAndHost) {
	status, err := client.GetVolume(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("GetVolume: %s", status)

}

func VolumeFiles(client pb.CommsProtoClient, v *pb.Volume) {
	status, err := client.VolumeFiles(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("VolumeFiles: %s", status)

}
