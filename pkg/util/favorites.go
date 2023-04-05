package util

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

type Favorite struct {
	Count   int
	Url     string
	LastUse int64
}

func sortFaves(favorites map[string]Favorite) {
	map_keys := maps.Keys(favorites)
	sort.SliceStable(map_keys, func(i, j int) bool {
		l := favorites[map_keys[i]]
		r := favorites[map_keys[j]]

		if l.Count == r.Count {
			return l.LastUse < r.LastUse
		} else {
			return l.Count < r.Count
		}
	})
}

func GetFavorites(path string) (map[string]interface{}, error) {
	payload, err := LoadFavorites(path)
	if err != nil {
		return nil, err
	}

	sortFaves(payload)
	map_keys := maps.Keys(payload)
	favorites := map[string]interface{}{}
	for idx, key := range map_keys {
		if idx < 15 {
			favorites[key] = payload[key].Url
		}
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
		favorites[key_path] = Favorite{Url: newest, Count: val.Count + 1, LastUse: time.Now().Unix()}
	} else {
		favorites[key_path] = Favorite{Url: newest, Count: 1, LastUse: time.Now().Unix()}
	}

	bin, err := json.Marshal(favorites)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bin, 0)
}
