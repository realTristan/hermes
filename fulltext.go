package hermes

import (
	"strings"
	"sync"
)

// Full Text struct
type FullText struct {
	mutex     *sync.RWMutex
	wordCache map[string][]int
	words     []string
	data      []map[string]string
}

// InitMap function
func InitMap(data []map[string]string, schema map[string]bool) (*FullText, error) {
	var ft *FullText = &FullText{
		mutex:     &sync.RWMutex{},
		wordCache: map[string][]int{},
		words:     []string{},
		data:      []map[string]string{},
	}

	// Load the cache data
	if err := ft.loadCacheData(data, schema); err != nil {
		return nil, err
	}

	// Return the cache
	return ft, nil
}

// InitJson function
func InitJson(file string, schema map[string]bool) (*FullText, error) {
	var ft *FullText = &FullText{
		mutex:     &sync.RWMutex{},
		wordCache: map[string][]int{},
		words:     []string{},
		data:      []map[string]string{},
	}

	// Read the json data
	if data, err := readJson(file); err != nil {
		return nil, err
	} else {
		ft.data = data
	}

	// Load the cache data
	if err := ft.loadCacheData(ft.data, schema); err != nil {
		return nil, err
	}

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
				if _, ok := ft.wordCache[word]; !ok {
					ft.wordCache[word] = []int{i}
					ft.words = append(ft.words, word)
					continue
				}
				if containsInt(ft.wordCache[word], i) {
					continue
				}
				ft.wordCache[word] = append(ft.wordCache[word], i)
			}
		}
	}
	return nil
}
