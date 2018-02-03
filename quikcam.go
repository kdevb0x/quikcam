package main

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
	//"github.com/blackjack/webcam/ioctl"
)

// Open webcam device
func newCam(path string) (*Webcam, error) {
	dev, err := os.OpenFile(path, unix.O_RDWR|unix.O_NONBLOCK, 0)
	fd := dev.Fd()
	if err != nil || fd < 0 {
		log.Fatal("There was an error opening the device: %s", err)
		return nil, err
	}
	cam := new(Webcam)
	cam.devpath = dev
	cam.fd = fd
	return cam, nil
	//buff := make([][]byte, dev)

}
func (w *Webcam) GrabFrame() error {
	return nil
}

type Cam interface {
	GrabFrame() error
}
type Webcam struct {
	devpath *os.File
	fd      uintptr
	buffers [][]byte
}

func main() {
	var path = "/dev/video0"
	webcam, err := newCam(path)
	if err != nil {
		log.Println("There was an error instantiating webcam: %s", err)
	}
	stub := webcam.GrabFrame()
	log.Println(stub)
	os.Exit(0)
}
