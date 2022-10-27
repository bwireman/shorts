package util

import (
	"encoding/json"
	"fmt"
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

const quit = "quit"

func OpenURL(url string, conf *Config) error {
	fullArgs := append(conf.BrowserCommand, url)
	cmd := fullArgs[0]
	args := fullArgs[1:]
	return exec.Command(cmd, args...).Run()
}

func makePreview(keys []string, choices map[string]interface{}, mode Mode) func(idx, _, _ int) string {
	return func(idx, _, _ int) string {
		if idx == -1 {
			return "nil"
		}

		chosenKey := keys[idx]
		if chosenKey == quit {
			return "See ya"
		}

		switch val := choices[chosenKey]; val.(type) {
		case string:
			emoji := "🚀"

			switch mode {
			case Directory:
				emoji = "📁"
			case Binary:
				emoji = "🖥️"
			}

			return fmt.Sprintf("%s %s", emoji, val)

		default:
			preview, err := json.MarshalIndent(val, "", " ")
			if err != nil {
				return "Could not load preview"
			}
			return string(preview)
		}

	}
}

func Choose(choices map[string]interface{}, conf *Config, mode Mode) (string, Mode, error) {
	keys := []string{}
	for key := range choices {
		keys = append(keys, key)
	}
	keys = append(keys, quit)

	choice, err := fuzzyfinder.Find(keys, func(idx int) string { return keys[idx] }, fuzzyfinder.WithPreviewWindow(makePreview(keys, choices, mode)))
	if err != nil {
		return "", None, err
	}

	chosenKey := keys[choice]
	if chosenKey == quit {
		return "", Quit, nil
	}

	switch val := choices[chosenKey]; val.(type) {
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
