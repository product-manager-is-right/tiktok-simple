package test

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

func TestA(t *testing.T) {

	coverData, _ := readFrameAsJpeg("http://120.25.2.146:9000/tiktok/videos/test.mp4")
	pictureReader := bytes.NewReader(coverData)
	print(pictureReader)
}
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if jpeg.Encode(buf, img, nil) != nil {
		return nil, err
	}

	return buf.Bytes(), err
}
