package volume

import (
	"fmt"
	"testing"
)

func TestGetFilesPathSize(t *testing.T) {
	//create tmp dir with some files and tree of files
	// write known bytes to each file
	// pass tmp dir to FilesInPath
	// verify names and size
	//filelist, size, err := GetFilesPathSize("/home/gspivey/Documents/Development/go/src/github.com/drayage")
	// TODO using relative path for portability until real test is complete
	filelist, _, size, err := GetFilesPathSize(".")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Filelist: \n%v\n TotalSize %d\n Human Readable Size %s %s\n", filelist, size, IECByteFormat(uint64(size)), SIByteFormat(uint64(size)))
}

func TestVolumeExpiriment(t *testing.T) {
	VolumeExpiriment()
}
