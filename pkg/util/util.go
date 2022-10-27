package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

type Mode = int

const (
	None Mode = iota
	Website
	Directory
	Binary
	Quit
)

func OpenURL(url string, conf *Config) error {
	fullArgs := append(conf.BrowserCommand, url)

	cmd := fullArgs[0]
	args := fullArgs[1:]

	return exec.Command(cmd, args...).Run()
}

const quit = "quit"

func Choose(choices map[string]interface{}, conf *Config, mode Mode) (string, Mode, error) {
	keys := []string{}
	for key := range choices {
		keys = append(keys, key)
	}
	keys = append(keys, quit)

	choice, err := fuzzyfinder.Find(keys, func(idx int) string { return keys[idx] }, fuzzyfinder.WithPreviewWindow(func(idx, _, _ int) string {
		if idx == -1 {
			return "nil"
		}

		chosenKey := keys[idx]
		if chosenKey == quit {
			return "See ya"
		}

		val := choices[chosenKey]

		switch val.(type) {

		case string:
			emoji:= "🚀"

			switch mode {
			case Directory:
				emoji = "📁" 
			case Binary:
				emoji = "🖲️" 
			default:
			}
			
			return fmt.Sprintf("%s %s", emoji, val)

		default:
			preview, err := json.MarshalIndent(val, "", " ")
			if err != nil {
				return "Could not load preview"
			}
			return string(preview)
		}

	}))

	if err != nil {
		return "", None, err
	}

	chosenKey := keys[choice]
	if chosenKey == quit {
		return "", Quit, nil
	}

	val := choices[chosenKey]
	switch val.(type) {
	case string:
		return val.(string), mode, nil
	default:
		switch chosenKey {
		case conf.BinaryDirName:
			mode = Binary
		case conf.DirectoriesDirName:
			mode = Directory
		default:
			mode = Website
		}

		return Choose(val.(map[string]interface{}), conf, mode)
	}
}

func LoadPaths(path string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

type Config struct {
	BrowserCommand     []string `json:"BrowserCommand,omitempty"`
	BinaryDirName      string   `json:"BinaryDirName,omitempty"`
	DirectoriesDirName string   `json:"DirectoriesDirName,omitempty"`
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

var default_command = []string{"xdg-open"}

const default_dir_val = "bin"
const default_bin_val = "directories"

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
	err = json.Unmarshal(content, &conf)
	if err != nil {
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
