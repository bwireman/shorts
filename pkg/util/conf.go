package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	BrowserCommand     []string `json:"BrowserCommand,omitempty"`
	BinaryDirName      string   `json:"BinaryDirName,omitempty"`
	DirectoriesDirName string   `json:"DirectoriesDirName,omitempty"`
}

var default_command = []string{"xdg-open"}

const default_dir_val = "bin"
const default_bin_val = "directories"

func LoadPaths(path string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	if err = json.Unmarshal(content, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func LoadConfig(path string) (*Config, error) {
	if !checkFileExists(path) {
		return &Config{
			BrowserCommand:     default_command,
			BinaryDirName:      "bin",
			DirectoriesDirName: "directories",
		}, nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := new(Config)
	if err := json.Unmarshal(content, &conf); err != nil {
		return nil, err
	}

	if len(conf.BrowserCommand) == 0 {
		conf.BrowserCommand = default_command
	}

	if len(conf.BinaryDirName) == 0 {
		conf.BinaryDirName = default_bin_val
	}

	if len(conf.DirectoriesDirName) == 0 {
		conf.DirectoriesDirName = default_dir_val
	}

	return conf, nil
}
