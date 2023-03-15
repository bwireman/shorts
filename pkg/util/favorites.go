package util

import (
	"encoding/json"
	"os"

	"sort"

	"golang.org/x/exp/maps"
)

func GetFavorites(path string) ([]string, error) {
	payload, err := LoadFavorites(path)
	if err != nil {
		return nil, err
	}

	return maps.Keys(payload), nil
}

func UpdateFavorites(path, newest string) error {
	favorites, err := LoadFavorites(path)
	if err != nil {
		return err
	}

	if val, ok := favorites[newest]; ok {
		favorites[newest] = val + 1
	} else {
		favorites[newest] = 1
	}

	keys := maps.Keys(favorites)

	sort.SliceStable(keys, func(i, j int) bool {
		return favorites[keys[i]] < favorites[keys[j]]
	})

	new_favorites := map[string]int{}
	for i, k := range keys {
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
