package util

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

func GetFavorites(path string) (map[string]interface{}, error) {
	payload, err := LoadFavorites(path)
	if err != nil {
		return nil, err
	}

	favorites := map[string]interface{}{}
	for k, v := range payload {
		favorites[k] = v.Url
	}

	return favorites, nil
}

type Favorite struct {
	Count int
	Url   string
}

func UpdateFavorites(path string, keys []string, newest string) error {
	favorites, err := LoadFavorites(path)
	if err != nil {
		return err
	}

	key_path := strings.Join(keys, " |> ")
	if val, ok := favorites[key_path]; ok {
		favorites[key_path] = Favorite{Url: newest, Count: val.Count + 1}
	} else {
		favorites[key_path] = Favorite{Url: newest, Count: 1}
	}

	map_keys := maps.Keys(favorites)
	sort.SliceStable(map_keys, func(i, j int) bool {
		return favorites[map_keys[i]].Count < favorites[map_keys[j]].Count
	})

	new_favorites := map[string]Favorite{}
	for i, k := range map_keys {
		if i < 10 {
			new_favorites[k] = favorites[k]
		}
	}

	bin, err := json.Marshal(new_favorites)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bin, 0)
}
