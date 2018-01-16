package client

import (
	"context"
	"io"
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
	stream, err := client.LSVolume(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Print("LSVolume: \n")
	for {
		dockerVolume, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(dockerVolume.Name)
		log.Println(dockerVolume.Updated)
		log.Println(dockerVolume.Size)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}

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

func ProcesVolumeFiles(client pb.CommsProtoClient, v *pb.Volume) {
	stream, err := client.VolumeFiles(context.Background(), v)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}
