package streamer

import (
	"fmt"
	"io"
	"log"
	"net"
	"testing"
)

/*
var file = make(map[string][]byte)
file["data/test.file.yaml"] =  i*/
var filecontent = `---
receipt:     Oz-Ware Purchase Invoice
date:        2012-08-06
customer:
    first_name:   Dorothy
    family_name:  Gale

items:
    - part_no:   A4786
      descrip:   Water Bucket (Filled)
      price:     1.47
      quantity:  4

    - part_no:   E1628
      descrip:   High Heeled "Ruby" Slippers
      size:      8
      price:     133.7
      quantity:  1

bill-to:  &id001
    street: |
            123 Tornado Alley
            Suite 16
    city:   East Centerville:w

    state:  KS

ship-to:  *id001

specialDelivery:  >
    Follow the Yellow Brick
    Road to the Emerald City.
    Pay no attention to the
    man behind the curtain.
...`

func TestFileQueueCH(t *testing.T) {
	var fileQueue = make(chan map[string][]byte)
	var filemap = make(map[string][]byte)
	filemap["test"] = []byte(`1234 safds
`)
	go PostFileCH(filemap, fileQueue)
	b := ReadFileCH(fileQueue)
	for label, content := range b {
		fmt.Printf("%s\n%s", label, string(content))
	}
}

// Success is opening socket and streaming a file into it without errors
// TODO Rather then the tough md5 sum testing lets test official success on
//    - Check if file exists
//    - Check if bytes are between 600-700
func TestTCPFileReceiver(t *testing.T) {
	//TODO Relative path needed for portability
	filename := "/home/gspivey/Documents/Development/go/src/github.com/drayage/testdata/output/testfile.txt"
	go TCPFileReceiver(filename)

	// Open connection and send file to our TCPFileReceiver
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	write := io.Writer(conn)
	write.Write([]byte(filecontent))
	conn.Close()

	/*
		oursum := md5.Sum([]byte(filecontent))
		returnedFile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Returned File Content %s\n", string(returnedFile))
		fmt.Printf("Original File Content %s\n", filecontent)
		newsum := md5.Sum(returnedFile)
		if oursum != newsum {
			fmt.Printf("File sum values are different %x %x", oursum, newsum)
			panic("File sum values are different")
		} else {
			fmt.Printf("File sum values are %x %x", oursum, newsum)
		}
	*/
}

// You can also test negative case
