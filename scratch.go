package main

/*
import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
	//"github.com/blackjack/webcam/ioctl"
)

// Open webcam device
func newCam(path string) (*Webcam, error) {
	dev, err := os.OpenFile(path, unix.O_RDONLY|unix.O_NONBLOCK, 0)
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
func ddSnap(savepath string) error {
func newBuff(w *Webcam, size uint) {
}

func (w *Webcam) GrabFrame() error {
	return nil
}

type Cam interface {
	newCam() error
	Close() error
}
type Webcam struct {
	devpath *os.File
	fd      uintptr
	buffers [][]byte
}

func (w *Webcam) Close() error {
	err := w.devpath.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	err = nil
	return err
}

func main() {
	var path = "/dev/video0"

	webcam, err := newCam(path)
	if err != nil {
		log.Println("There was an error instantiating webcam: %s", err)
	}
	defer func() {
		closeError := webcam.Close()
		if closeError != nil {
			log.Println("ERROR CLOSING DEVICE: %s", closeError)
			panic(closeError)
		}
	}()
	stub := webcam.GrabFrame()
	log.Println(stub)
	os.Exit(0)
}
*/
