package volume

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func VolumePath(volumeName string) (mountPath string, err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	emptyArgs := filters.Args{}
	//
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	if err != nil {
		return "", err
	}
	//log.Printf("volume list body content \n%s\n", volumeListBody)
	v := volumeListBody.Volumes
	for _, vol := range v {
		//log.Printf("volume Metadata \nName: %s\n MountPath: %s\n", vol.Name, vol.Mountpoint)
		if vol.Name == volumeName {
			mountPath = vol.Mountpoint
			return mountPath, nil
		}
	}
	return "", nil

}

func AddVolume(volumeName string) (volume types.Volume, err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return volume, err
	}
	options := volumetypes.VolumeCreateBody{}
	options.Name = volumeName
	volume, err = cli.VolumeCreate(context.Background(), options)
	return volume, err
}

func RemoveVolume(volumeName string) (err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = cli.VolumeRemove(context.Background(), volumeName, false)
	return err

}

func highestTime(modifiedTime map[string]time.Time) (highest time.Time) {
	for _, t := range modifiedTime {
		if t.After(highest) {
			highest = t
		}
	}
	return
}

func AnalyzeVolume(volume *types.Volume) (name string, size int64, updated time.Time, err error) {
	_, timeMap, size, err := GetFilesPathSize(volume.Mountpoint)
	updated = highestTime(timeMap)

	return volume.Name, size, updated, err
}

func ListVolumes() (volumeList []*types.Volume, err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return volumeList, err
	}

	emptyArgs := filters.Args{}
	//
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	if err != nil {
		return volumeList, err
	}
	volumeList = volumeListBody.Volumes
	return volumeList, err
}

func VolumeExpiriment() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	emptyArgs := filters.Args{}
	//
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	fmt.Printf("volume list body content \n%v\n", volumeListBody)
	fmt.Print("Volume's:\n")
	v := volumeListBody.Volumes
	// TODO add filter on driver
	for _, vol := range v {
		fmt.Printf("Name: %s\n MountPath: %s\n Driver: %s\n\n", vol.Name, vol.Mountpoint, vol.Driver)
		vn, sz, up, _ := AnalyzeVolume(vol)
		szh := SIByteFormat(uint64(sz))
		fmt.Printf("Analyzed Volume: %s %s, %v\n", vn, szh, up)
	}
}
