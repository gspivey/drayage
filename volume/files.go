package volume

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// We will use the FiveGiB limit for dynamic transfer limit
const (
	FiveGB  int64 = 5000000000 // Base10 - 1000
	FiveGiB int64 = 5368709000 // Base2 - 1024
)

//IECFormat prints bytes in the International Electrotechnical Commission format
func IECByteFormat(num_in uint64) string {
	suffix := "B" //just assume bytes
	num := float64(num_in)
	units := []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}
	for _, unit := range units {
		if num < 1024.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = (num / 1024)
	}
	return fmt.Sprintf("%.1f%s%s", num, "Yi", suffix)
}

//SIFormat prints bytes in the International System of Units format
func SIByteFormat(num_in uint64) string {
	suffix := "B" //just assume bytes
	num := float64(num_in)
	units := []string{"", "K", "M", "G", "T", "P", "E", "Z"}
	for _, unit := range units {
		if num < 1000.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = (num / 1000)
	}
	return fmt.Sprintf("%.1f%s%s", num, "Yi", suffix)
}

// Returns the names of the subdirectories (including their paths)
// that match the specified search pattern in the specified directory.
//func GetDirectories(root, pattern string) ([]string, error) {
func GetFiles(root string) ([]string, error) {
	dirs := make([]string, 0, 144)
	return dirs, filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			dirs = append(dirs, path)
			return nil
		}

		return nil
	})

}

func GetFilesPathSize(path string) (filelist map[string]int64, modifiedTime map[string]time.Time, totalsize int64, err error) {
	// composite type map needs to be initialized
	filelist = make(map[string]int64)
	modifiedTime = make(map[string]time.Time)
	return filelist, modifiedTime, totalsize, filepath.Walk(path, func(fpath string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			currentfile, err := os.Stat(fpath)
			if err != nil {
				return err
			}
			modifiedTime[fpath] = fi.ModTime()
			// get the size
			size := currentfile.Size()
			filelist[fpath] = size
			totalsize += size
			return nil
		}
		return nil

	})
}
