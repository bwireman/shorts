package main

import (
	"fmt"
	"os"
	"path"

	util "github.com/bwireman/shorts/pkg/util"
)

func main() {
	home := os.Getenv("HOME")
	conf, err := util.LoadConfig(path.Join(home, ".shorts", "config.json"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	paths, err := util.LoadPaths(path.Join(home, ".shorts", "paths.json"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	faves_path := path.Join(home, ".shorts", "faves.json")
	faves, err := util.GetFavorites(faves_path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	faves_map := map[string]interface{}{}

	for _, f := range faves {
		faves_map[f] = f
	}

	paths["faves"] = faves_map

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
		util.UpdateFavorites(faves_path, choice)
		util.OpenURL(choice, conf)
		fmt.Printf("echo %s", choice)
	case util.Quit:
		fmt.Print("echo 👋 See ya")
	default:
		fmt.Print("echo An error occurred")
		os.Exit(1)
	}
}
