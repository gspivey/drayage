package archive

import (
	"fmt"
	"testing"
)

func TestWriteToTar(t *testing.T) {
	name := "gtest"
	err := WriteToTar(name)
	if err != nil {
		fmt.Println(err)
	}
}
