package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/drayage/protocol"
	"github.com/drayage/volume"
)

type Server struct{}

func (s *Server) RMVolume(ctx context.Context, volumeName *pb.Volume) (status *pb.Status, err error) {
	err = volume.RemoveVolume(volumeName.Name)
	if err != nil {
		return status, err
	}
	return &pb.Status{"success"}, err
}

func (s *Server) LSVolume(volumeName *pb.Volume, stream pb.CommsProto_LSVolumeServer) error {
	// stream back values
	pbDockerVolume := pb.DockerVolume{}
	v, _ := volume.ListVolumes()
	log.Printf("\n result \n %v", v)
	// TODO add filter on driver
	for _, vol := range v {
		fmt.Printf("Name: %s\n MountPath: %s\n Driver: %s\n\n", vol.Name, vol.Mountpoint, vol.Driver)
		vn, sz, up, _ := volume.AnalyzeVolume(vol)
		szh := volume.SIByteFormat(uint64(sz))
		fmt.Printf("Analyzed Volume: %s %s, %v\n", vn, szh, up)
		// TODO Protobuf type modification needed
		pbDockerVolume.Name = vn
		pbDockerVolume.Size = szh
		pbDockerVolume.Updated = up.String()
		_ = stream.Send(&pbDockerVolume)
	}
	return nil

}

func (s *Server) AddVolume(ctx context.Context, volumeName *pb.Volume) (status *pb.Status, err error) {
	message := "success"
	v, err := volume.AddVolume(volumeName.Name)
	log.Printf("Volume name and mountpath %s %s\n", v.Name, v.Mountpoint)
	if err != nil {
		message = "failed"
	}
	// Docker VolumeCreate
	return &pb.Status{message}, err
}

func (s *Server) GetVolume(ctx context.Context, volumeAndHost *pb.VolumeAndHost) (*pb.Status, error) {
	// Drayage volume streamer
	return &pb.Status{"complete"}, nil
}

func (s *Server) VolumeFiles(volumeName *pb.Volume, stream pb.CommsProto_VolumeFilesServer) error {
	// Inspect Docker Volume VolumeInspect
	// Use root path with drayage GetFilesPathSize
	// stream files name back

	pbFile := pb.Files{""}

	path, err := volume.VolumePath(volumeName.Name)
	if err != nil {
		return err
	}
	log.Print(path)
	filelist, _, size, _ := volume.GetFilesPathSize(path)
	log.Printf("Filelist: \n%v\n TotalSize %d\n Human Readable Size %s %s\n", filelist, size,
		volume.IECByteFormat(uint64(size)), volume.SIByteFormat(uint64(size)))

	for filepath := range filelist {
		log.Print(filepath)
		pbFile.Path = filepath
		_ = stream.Send(&pbFile)
	}

	return nil
}
