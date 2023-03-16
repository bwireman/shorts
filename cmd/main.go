package main

import (
	"fmt"

	util "github.com/bwireman/shorts/pkg/util"
)

func main() {
	conf, err := util.LoadConfig(util.ConfigPath)
	util.MaybeExit(err)

	paths, err := util.LoadPaths(util.PathsPath)
	util.MaybeExit(err)

	faves, err := util.GetFavorites(util.FavesPath)
	util.MaybeExit(err)

	if len(faves) > 0 {
		paths[util.Faves] = faves
	}

	choice, mode, err := util.Choose(paths, paths, conf, []string{}, util.Website)
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
