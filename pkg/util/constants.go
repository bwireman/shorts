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
const Faves = "favorites"
const SeeYa = "👋 See ya"

var default_command = []string{"xdg-open"}

const quit = "❌ quit"
const back = "🔙 back"
const shorts_config_path = ".shorts"
const default_dir_val = "directories"
const default_bin_val = "bin"

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
