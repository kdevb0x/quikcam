package main

import (
	"os"
	"bytes"
	"io"
	"bufio"
	"io/ioutil"
)

func newCam(path string) webcam, error {
	dev := os.Open(path)

	cam.

}
type Webcam interface {
	Scan

}
type webcam struct {
	devpath string
}

func main() {
	var path = "/dev/video0"
	cam := newCam(path)
}
