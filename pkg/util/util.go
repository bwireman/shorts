package util

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
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

func makePreview(keys []string, choices map[string]interface{}, mode Mode) func(idx, _, _ int) string {
	return func(idx, _, _ int) string {
		if idx == -1 {
			return "nil"
		}

		chosenKey := keys[idx]
		if chosenKey == quit {
			return "👋 See ya"
		}

		if chosenKey == back {
			return "🔙"
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

func Choose(all_paths map[string]interface{}, choices map[string]interface{}, conf *Config, previous_keys []string, mode Mode) (string, Mode, error) {
	keys := []string{}
	for key := range choices {
		keys = append(keys, key)
	}

	keys = append(keys, quit)
	if len(previous_keys) > 0 {
		keys = append(keys, back)
	}

	choice, err := fuzzyfinder.Find(keys, func(idx int) string { return keys[idx] }, fuzzyfinder.WithPromptString("👉 "), fuzzyfinder.WithPreviewWindow(makePreview(keys, choices, mode)))
	if err != nil {
		return "", None, err
	}

	chosenKey := keys[choice]
	if chosenKey == quit {
		return "", Quit, nil
	}

	if chosenKey == back {
		prev_choice := all_paths
		if len(previous_keys) == 0 {
			previous_keys = []string{}
		} else {
			previous_keys = previous_keys[:len(previous_keys)-1]
		}

		for _, prev := range previous_keys {
			prev_choice = prev_choice[prev].(map[string]interface{})
		}

		return Choose(all_paths, prev_choice, conf, previous_keys, mode)
	}

	switch val := choices[chosenKey]; val.(type) {
	case string:
		as_string := val.(string)

		all_keys := append(previous_keys, chosenKey)
		if !slices.Contains(all_keys, Faves) {
			if err := UpdateFavorites(FavesPath, all_keys, as_string); err != nil {
				return "", Quit, err
			}
		}

		return as_string, mode, nil
	default:
		switch chosenKey {
		case conf.BinaryDirName:
			mode = Binary
		case conf.DirectoriesDirName:
			mode = Directory
		default:
			mode = Website
		}

		return Choose(all_paths, val.(map[string]interface{}), conf, append(previous_keys, chosenKey), mode)
	}
}
