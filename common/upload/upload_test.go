package upload

import (
	"fmt"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fi, _ := os.Stat("./test.mp4")
	url, err := UploadVideo(file, fi.Size(), 3254531)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
