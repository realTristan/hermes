package hermes

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
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
func InitJson(file string) *FTS {
	var fts *FTS = &FTS{
		mutex: &sync.RWMutex{},
		cache: map[string][]int{},
		keys:  []string{},
		json:  []map[string]string{},
	}

	// Load the json data
	fts._ReadJson(file)

	// Load the cache
	fts._LoadCacheJson(fts.json)

	// Return the cache
	return fts
}

// Set a value in the cache
func (fts *FTS) Set(key string, value map[string]string) {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts._Set(key, value)
}

// Set a value in the cache
func (fts *FTS) _Set(key string, value map[string]string) {
	fts.json = append(fts.json, value)
	// Loop through the value
	for _, v := range value {
		// Loop through the words
		for _, word := range strings.Split(v, " ") {
			switch {
			case len(word) <= 1:
				continue
			case !_IsAlphaNum(word):
				continue
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
	fts._Clean()
}

// Clean the FTS cache
func (fts *FTS) _Clean() {
	fts.cache = map[string][]int{}
	fts.keys = []string{}
	fts.json = []map[string]string{}
}

// Reset the FTS cache
func (fts *FTS) Reset(fileName string) {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts._Reset(fileName)
}

// Reset the FTS cache
func (fts *FTS) _Reset(fileName string) {
	fts.cache = map[string][]int{}
	fts.keys = []string{}
	fts.json = []map[string]string{}
	fts._ReadJson(fileName)
	fts._LoadCacheJson(fts.json)
}

// Read the json data
func (fts *FTS) _ReadJson(fileName string) {
	var data, _ = os.ReadFile(fileName)
	json.Unmarshal(data, &fts.json)
}

// Load the FTS cache
func (fts *FTS) _LoadCacheJson(json []map[string]string) error {
	// Loop through the json data
	for i, item := range json {
		if reflect.TypeOf(item).Kind() != reflect.Map {
			return errors.New("invalid data type. values inside json array must be: map[string]string")
		}

		for _, value := range item {
			if reflect.TypeOf(value).Kind() != reflect.String {
				return errors.New("invalid data type. values inside json map must be: string")
			}

			// Clean the value
			value = strings.TrimSpace(value)
			value = _RemoveDoubleSpaces(value)
			value = strings.ToLower(value)

			// Loop through the words
			for _, word := range strings.Split(value, " ") {
				switch {
				case len(word) <= 1:
					continue
				case !_IsAlphaNum(word):
					continue
				}
				if _, ok := fts.cache[word]; !ok {
					fts.cache[word] = []int{i}
					fts.keys = append(fts.keys, word)
					continue
				}
				if !_ContainsInt(fts.cache[word], i) {
					continue
				}
				fts.cache[word] = append(fts.cache[word], i)
			}
		}
	}
	return nil
}
