package hermes

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

// Full Text Search struct
type FTS struct {
	mutex *sync.RWMutex
	cache map[string][]int
	keys  []string
	json  []map[string]string
}

// InitCache function
func InitJson(file string, keySettings map[string]bool) *FTS {
	var fts *FTS = &FTS{
		mutex: &sync.RWMutex{},
		cache: map[string][]int{},
		keys:  []string{},
		json:  []map[string]string{},
	}

	// Load the json cache
	fts.loadJson(file)
	fts.loadCacheJson(fts.json, keySettings)

	// Return the cache
	return fts
}

// Set a value in the cache
func (fts *FTS) Set(key string, value map[string]string) {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts.set(key, value)
}

// Set a value in the cache
func (fts *FTS) set(key string, value map[string]string) {
	fts.json = append(fts.json, value)
	// Loop through the value
	for _, v := range value {
		// Loop through the words
		for _, word := range strings.Split(v, " ") {
			switch {
			case len(word) <= 1:
				continue
			case !isAlphaNum(word):
				word = removeNonAlphaNum(word)
			}
			if _, ok := fts.cache[word]; !ok {
				fts.cache[word] = []int{len(fts.json) - 1}
				fts.keys = append(fts.keys, word)
				continue
			}
			fts.cache[word] = append(fts.cache[word], len(fts.json)-1)
		}
	}
}

// Clean the FTS cache
func (fts *FTS) Clean() {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts.clean()
}

// Clean the FTS cache
func (fts *FTS) clean() {
	fts.cache = map[string][]int{}
	fts.keys = []string{}
	fts.json = []map[string]string{}
}

// Reset the FTS cache
func (fts *FTS) Reset(file string, keySettings map[string]bool) {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts.reset(file, keySettings)
}

// Reset the FTS cache
func (fts *FTS) reset(file string, keySettings map[string]bool) {
	fts.cache = map[string][]int{}
	fts.keys = []string{}
	fts.json = []map[string]string{}
	fts.loadJson(file)
	fts.loadCacheJson(fts.json, keySettings)
}

// Read the json data
func (fts *FTS) loadJson(file string) {
	var data, _ = os.ReadFile(file)
	json.Unmarshal(data, &fts.json)
}

// Load the FTS cache
func (fts *FTS) loadCacheJson(json []map[string]string, keySettings map[string]bool) error {
	// Loop through the json data
	for i, item := range json {
		// Loop through the map
		for key, value := range item {
			if !keySettings[key] {
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
				if _, ok := fts.cache[word]; !ok {
					fts.cache[word] = []int{i}
					fts.keys = append(fts.keys, word)
					continue
				}
				if containsInt(fts.cache[word], i) {
					continue
				}
				fts.cache[word] = append(fts.cache[word], i)
			}
		}
	}
	return nil
}
