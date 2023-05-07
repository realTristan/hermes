package nocache

import (
	"strings"
	"sync"

	Utils "github.com/realTristan/Hermes/utils"
)

/*
FullText is a struct that represents a full text search cache. It has the following fields:
- mutex (*sync.RWMutex): a pointer to a read-write mutex used to synchronize access to the cache
- wordCache (map[string][]int): a map where the keys are words and the values are arrays of integers representing the indices of the data items that contain the word
- words ([]string): a slice of strings representing all the unique words in the cache
- data ([]map[string]interface{}): a slice of maps representing the data items in the cache, where the keys are the names of the fields and the values are the field values
*/
type FullText struct {
	mutex     *sync.RWMutex
	wordCache map[string][]int
	words     []string
	data      []map[string]interface{}
}

// Initialize the full-text cache with the provided data.
// This function is thread safe.
func InitWithMap(data []map[string]interface{}) (*FullText, error) {
	var ft *FullText = &FullText{
		mutex:     &sync.RWMutex{},
		wordCache: make(map[string][]int),
		words:     []string{},
		data:      []map[string]interface{}{},
	}

	// Load the cache data
	if err := ft.insert(data); err != nil {
		return nil, err
	}

	// Set the data
	ft.data = data

	// Return the full-text variable
	return ft, nil
}

// Initialize the full-text cache with the provided json file.
// This function is thread safe.
func InitWithJson(file string) (*FullText, error) {
	// Read the json data
	if data, err := Utils.ReadSliceJson(file); err != nil {
		return nil, err
	} else {
		return InitWithMap(data)
	}
}

// Insert data into the full-text cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) insert(data []map[string]interface{}) error {
	// Loop through the data
	for i, item := range data {
		// Loop through the map
		for key, value := range item {
			// Get the string value
			var strv string
			if _strv := fullTextMap(value); len(_strv) > 0 {
				strv = _strv
			} else {
				continue
			}

			// Clean the value
			strv = strings.TrimSpace(strv)
			strv = Utils.RemoveDoubleSpaces(strv)
			strv = strings.ToLower(strv)

			// Loop through the words
			for _, word := range strings.Split(strv, " ") {
				switch {
				case len(word) <= 1:
					continue
				case !Utils.IsAlphaNum(word):
					word = Utils.RemoveNonAlphaNum(word)
				}
				if _, ok := ft.wordCache[word]; !ok {
					ft.wordCache[word] = []int{i}
					ft.words = append(ft.words, word)
					continue
				}
				if Utils.ContainsInt(ft.wordCache[word], i) {
					continue
				}
				ft.wordCache[word] = append(ft.wordCache[word], i)
			}

			// Set the value
			data[i][key] = strv
		}
	}
	return nil
}
