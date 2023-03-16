package main

import (
	"fmt"
	"os"

	util "github.com/bwireman/shorts/pkg/util"
)

func main() {
	conf, err := util.LoadConfig(util.ConfigPath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	paths, err := util.LoadPaths(util.PathsPath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	faves, err := util.GetFavorites(util.FavesPath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if len(faves) > 0 {
		paths[util.Faves] = faves
	}

	choice, mode, err := util.Choose(paths, paths, conf, []string{}, util.Website)
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
		util.OpenURL(choice, conf)
		fmt.Printf("echo %s", choice)
	case util.Quit:
		fmt.Print("echo 👋 See ya")
	default:
		fmt.Print("echo An error occurred")
		os.Exit(1)
	}
}
