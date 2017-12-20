package streamer

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func PostFileCH(mappedFile map[string][]byte, ch chan map[string][]byte) {
	ch <- mappedFile
}

func ReadFileCH(channel chan map[string][]byte) (file map[string][]byte) {
	file = <-channel
	return file
}

func FileQueueCH() {
	var fileQueue = make(chan map[string][]byte)
	var filemap = make(map[string][]byte)
	filemap["test"] = []byte(`1234 safds`)
	go PostFileCH(filemap, fileQueue)
	b := ReadFileCH(fileQueue)
	for label, content := range b {
		fmt.Printf("%s\n %v", label, content)
	}
}

func TCPFileReceiver(file string) (err error) {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		return err
	}
	f, _ := os.Create(file)
	defer f.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go fileStream(conn, f) // handle connections concurrently
	}
	f.Close()
	return nil

}

func fileStream(c net.Conn, f *os.File) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		f.Write(input.Bytes())
		fmt.Printf("Receiveed content %s\n", input.Text())
	}
}

func DiskFileReader(file string) (content []byte, md5sum [16]byte, err error) {
	content, err = ioutil.ReadFile(file)
	if err != nil {
		return content, md5sum, err
	}
	fmt.Printf("Read content %s\n", string(content))
	md5sum = md5.Sum(content)
	fmt.Printf("MD5sum %x", md5sum)
	return content, md5sum, nil
}
