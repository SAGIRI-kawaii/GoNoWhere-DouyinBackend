package upload

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUploadFile(t *testing.T) {
	data, err := ioutil.ReadFile("./test.mp4")
	url, err := UploadVideo(&data, 3214432254531)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
