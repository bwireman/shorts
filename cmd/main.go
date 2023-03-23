package main

import (
	"flag"
	"fmt"

	util "github.com/bwireman/shorts/pkg/util"
)

var skipToFaves = false

func main() {
	flag.BoolVar(&skipToFaves, "faves", false, "skip to the most commonly used shortcuts")
	flag.Parse()

	conf, err := util.LoadConfig(util.ConfigPath)
	util.MaybeExit(err)

	paths, err := util.LoadPaths(util.PathsPath)
	util.MaybeExit(err)

	faves, err := util.GetFavorites(util.FavesPath)
	util.MaybeExit(err)

	hasFaves := len(faves) > 0

	if hasFaves {
		paths[util.Faves] = faves
	}

	options := paths
	startingPath := []string{}
	if skipToFaves && hasFaves {
		startingPath = []string{util.Faves}
		options = paths[util.Faves].(map[string]interface{})
	}

	choice, mode, err := util.Choose(paths, options, conf, startingPath, util.Website)
	util.MaybeExit(err)

	switch mode {
	case util.Binary:
		fmt.Print(choice)
	case util.Directory:
		fmt.Printf("cd %s", choice)
	case util.Website:
		util.OpenURL(choice, conf)
		fmt.Printf("echo %s", choice)
	case util.Quit:
		fmt.Printf("echo %s", util.SeeYa)
	default:
		util.Exit("echo An error occurred")
	}
}
