package main

import (
	"fmt"
	"os"

	util "github.com/bwireman/shorts/pkg/util"
)

func main() {
	home := os.Getenv("HOME")
	conf, err := util.LoadConfig(fmt.Sprintf("%s/.shorts/config.json", home))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	choice, mode, err := util.Choose(conf, util.Website)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	switch mode {
	case util.Binary:
		fmt.Print(choice)
	case util.Directory:
		fmt.Printf("cd %s", choice)
	case util.Website:
		util.OpenURL(choice)
		fmt.Printf("echo %s", choice)
	default:
		fmt.Print("echo 'An error occurred'")
		os.Exit(1)
	}
}
