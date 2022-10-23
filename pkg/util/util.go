package util

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

type Mode = int

const (
	None Mode = iota
	Website
	Directory
	Binary
)

func OpenURL(url string) error {
	return exec.Command("firefox", "--new-tab", url).Run()
}

func Choose(choices map[string]interface{}, mode Mode) (string, Mode, error) {
	keys := []string{}
	for key := range choices {
		keys = append(keys, key)
	}

	choice, err := fuzzyfinder.Find(keys, func(idx int) string { return keys[idx] }, fuzzyfinder.WithPreviewWindow(func(idx, _, _ int) string {
		if idx == -1 {
			return "nil"
		}

		preview, err := json.MarshalIndent(choices[keys[idx]], "", " ")
		if err != nil {
			return "Could not load preview"
		}

		return string(preview)
	}))

	if err != nil {
		return "", None, err
	}

	chosenKey := keys[choice]
	val := choices[chosenKey]
	switch val.(type) {
	case string:
		return val.(string), mode, nil
	default:
		switch chosenKey {
		case "bin":
			mode = Binary
		case "directories":
			mode = Directory
		default:
			mode = Website
		}

		return Choose(val.(map[string]interface{}), mode)
	}
}

func LoadConfig(path string) (map[string]interface{}, error) {
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
