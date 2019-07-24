package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gspivey/drayage/volume"
)

//WriteToTar tars any found volumes
func WriteToTar(name string) (err error) {
	volumes, err := volume.ListVolumes()
	findVolume := false

	for _, vol := range volumes {
		if name != "" {
			if vol.Name != name {
				continue
			}
		}
		findVolume = true

		dir, err := os.Open(vol.Mountpoint)
		if err != nil {
			return fmt.Errorf("can't open %v (did you use sudo?) : %v", vol.Name, err)
		}
		defer dir.Close()

		files, err := dir.Readdir(0)
		if err != nil {
			return fmt.Errorf("can't read directory : %v", err)
		}

		destfile := vol.Name + ".tgz"

		//
		tarfile, err := os.Create(destfile)
		defer tarfile.Close()

		var fileWriter io.WriteCloser = tarfile

		fileWriter = gzip.NewWriter(tarfile) //gzip filter
		defer fileWriter.Close()

		tarfileWriter := tar.NewWriter(fileWriter) //tar filter
		defer tarfileWriter.Close()

		for _, fileInfo := range files {

			if fileInfo.IsDir() {
				continue
			}

			file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
			if err != nil {
				return fmt.Errorf("can't open file %v : %v", fileInfo.Name(), err)
			}
			defer file.Close()

			header := new(tar.Header)
			header.Name = file.Name()
			header.Size = fileInfo.Size()
			header.Mode = int64(fileInfo.Mode())
			header.ModTime = fileInfo.ModTime()

			err = tarfileWriter.WriteHeader(header)

			_, err = io.Copy(tarfileWriter, file)
		}
	}
	if findVolume == false {
		fmt.Println("volume(s) not found")
	}
	return nil
}
