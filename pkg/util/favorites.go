package util

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

type Favorite struct {
	Count int
	Url   string
}

func GetFavorites(path string) (map[string]interface{}, error) {
	payload, err := LoadFavorites(path)
	if err != nil {
		return nil, err
	}

	favorites := map[string]interface{}{}
	for key, fav := range payload {
		favorites[key] = fav.Url
	}

	return favorites, nil
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
	for idx, key := range map_keys {
		if idx < 10 {
			new_favorites[key] = favorites[key]
		}
	}

	bin, err := json.Marshal(new_favorites)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bin, 0)
}
