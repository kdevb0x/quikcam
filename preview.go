package main

import (
	"log"
	"os"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// local instance of gtk.Window
type window struct {
	gtk.Window
	callback
}

func (n *Window) Start()
func mainWindow(name string) (*gtk.Window, error) {
	gtk.Init(nil)
	win, err := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Printf(err)
		return nil, err
	}
	win.SetTitle(name)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})



