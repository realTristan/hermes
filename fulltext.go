package hermes

import (
	"fmt"
	"strings"
	"sync"
)

// Full Text struct
type FullText struct {
	mutex *sync.RWMutex
	cache map[string][]int
	words []string
	data  []map[string]string
}

// InitCache function
func InitJson(file string, schema map[string]bool) (*FullText, error) {
	var ft *FullText = &FullText{
		mutex: &sync.RWMutex{},
		cache: map[string][]int{},
		words: []string{},
		data:  []map[string]string{},
	}

	// Read the json data
	if data, err := readJson(file); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		ft.data = data
	}

	// Load the cache data
	ft.loadCacheData(ft.data, schema)

	// Return the cache
	return ft, nil
}

// Load the ft cache
func (ft *FullText) loadCacheData(data []map[string]string, schema map[string]bool) error {
	// Loop through the data
	for i, item := range data {
		// Loop through the map
		for key, value := range item {
			if !schema[key] {
				continue
			}

			// Clean the value
			value = strings.TrimSpace(value)
			value = removeDoubleSpaces(value)
			value = strings.ToLower(value)

			// Loop through the words
			for _, word := range strings.Split(value, " ") {
				switch {
				case len(word) <= 1:
					continue
				case !isAlphaNum(word):
					word = removeNonAlphaNum(word)
				}
				if _, ok := ft.cache[word]; !ok {
					ft.cache[word] = []int{i}
					ft.words = append(ft.words, word)
					continue
				}
				if containsInt(ft.cache[word], i) {
					continue
				}
				ft.cache[word] = append(ft.cache[word], i)
			}
		}
	}
	return nil
}
