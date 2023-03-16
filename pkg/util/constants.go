package util

import (
	"fmt"
	"os"
	"path"
)

// paths
var FavesPath string
var ConfigPath string
var PathsPath string

// constants
var Faves = "favorites"

const quit = "❌ quit"
const back = "🔙 back"
const shorts_config_path = ".shorts"

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	FavesPath = path.Join(home, shorts_config_path, "faves.json")
	ConfigPath = path.Join(home, shorts_config_path, "config.json")
	PathsPath = path.Join(home, shorts_config_path, "paths.json")
}
