package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/robmerrell/comandante"
)

var (
	savepath string
)

func snapJpg(savepath string) error {
	cmd := exec.Command("dd", "if=/dev/video0", "of="+savepath, "bs=1M", "count=1")
	fmt.Println("saving jpeg to ", savepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("There was an error saving jpg", err)
		return err
	}
	return nil
}
func Snap() error {
	err := snapJpg(savepath)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	bin := comandante.New("quikcam", "A command-line webcam utility.")
	bin.IncludeHelp()
	snapshot := comandante.NewCommand("snapshot", "snap a single jpg image", Snap)
	snapshot.Documentation = `
	snapshot takes a single webcam image snapshot in jpg format
	`
	snapshot.FlagInit = func(fs *flag.FlagSet) {
		fs.StringVar(&savepath, "o", "~/pic.jpg", "output path (default is ~/pic.jpg")
	}
	bin.RegisterCommand(snapshot)
	if err := bin.Run(); err != nil {
		panic(err)
	}
}
