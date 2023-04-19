package hermes

import (
	"strings"
	"sync"
)

/*
FullText is a struct that represents a full text search cache. It has the following fields:
- mutex (*sync.RWMutex): a pointer to a read-write mutex used to synchronize access to the cache
- wordCache (map[string][]int): a map where the keys are words and the values are arrays of integers representing the indices of the data items that contain the word
- words ([]string): a slice of strings representing all the unique words in the cache
- data ([]map[string]string): a slice of maps representing the data items in the cache, where the keys are the names of the fields and the values are the field values

Example Usage:

	ft := FullText{
		mutex:     &sync.RWMutex{},
		wordCache: map[string][]int{},
		words:     []string{},
		data:      []map[string]string{},
	}

fmt.Println(ft)
*/
type FullText struct {
	mutex     *sync.RWMutex
	wordCache map[string][]int
	words     []string
	data      []map[string]string
}

/*
InitMap initializes a FullText object by loading data into the wordCache field, based on the specified schema.
It takes in two parameters:
  - data ([]map[string]string): a slice of maps representing the data to be loaded into the wordCache field
  - schema (map[string]bool): a map representing the schema for the data, where the keys are the names of the fields and
    the values are booleans indicating whether the field should be included in the cache

It returns a pointer to the initialized FullText object and an error, if any.

Example Usage:

	data := []map[string]string{
		{"id": "1", "name": "John Doe", "description": "Software engineer"},
		{"id": "2", "name": "Jane Smith", "description": "Data analyst"},
	}

schema := map[string]bool{"name": true, "description": true}
ft, err := InitMap(data, schema)

	if err != nil {
		fmt.Println("Error initializing FullText object:", err)
		return
	}

fmt.Println("FullText object successfully initialized:", ft)
*/
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

/*
InitJson initializes a FullText object by loading data from a JSON file into the wordCache field, based on the specified schema.
It takes in two parameters:
  - file (string): the path of the JSON file containing the data to be loaded into the wordCache field
  - schema (map[string]bool): a map representing the schema for the data, where the keys are the names of the fields and the values are booleans
    indicating whether the field should be included in the cache

It returns a pointer to the initialized FullText object and an error, if any.

Example Usage:
schema := map[string]bool{"name": true, "description": true}
ft, err := InitJson("data.json", schema)

	if err != nil {
		fmt.Println("Error initializing FullText object:", err)
		return
	}

fmt.Println("FullText object successfully initialized:", ft)
*/
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

/*
loadCacheData is a method of the FullText type that loads data into the wordCache field based on the provided schema.
It takes in two parameters:
  - data ([]map[string]string): an array of maps representing the data to be loaded into the wordCache field
  - schema (map[string]bool): a map representing the schema for the data, where the keys are the names of the fields and the values are booleans
    indicating whether the field should be included in the cache

It returns an error, if any.

Example Usage:
schema := map[string]bool{"name": true, "description": true}

	data := []map[string]string{
		{"name": "apple", "description": "a round fruit"},
		{"name": "orange", "description": "a citrus fruit"},
	}

ft := FullText{}
err := ft.loadCacheData(data, schema)

	if err != nil {
		fmt.Println("Error loading data into wordCache:", err)
		return
	}

fmt.Println("Data successfully loaded into wordCache")
*/
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
