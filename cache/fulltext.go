package cache

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

// Full Text struct
type FullText struct {
	mutex         *sync.RWMutex
	wordCache     map[string][]string
	data          map[string]map[string]interface{}
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}

// Upload json data to the FullText cache with Mutex Locking
func (ft *FullText) UploadJson(file string, schema map[string]bool) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.uploadJson(file, schema)
}

// Upload json data to the FullText cache
func (ft *FullText) uploadJson(file string, schema map[string]bool) error {
	// Read the json
	if data, err := readJson(file); err != nil {
		return err
	} else {
		return ft.loadCacheData(data, schema)
	}
}

// Upload map data to the FullText cache with Mutex Locking
func (ft *FullText) UploadMap(data map[string]map[string]interface{}, schema map[string]bool) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.uploadMap(data, schema)
}

// Upload map data to the FullText cache
func (ft *FullText) uploadMap(data map[string]map[string]interface{}, schema map[string]bool) error {
	return ft.loadCacheData(data, schema)
}

// Load the FullText cache
func (ft *FullText) loadCacheData(data map[string]map[string]interface{}, schema map[string]bool) error {
	// Loop through the json data
	for itemKey, itemValue := range data {
		// Loop through the map
		for key, value := range itemValue {
			// Check if the key is in the schema
			if !schema[key] {
				continue
			}

			// Check if the value is a string
			if v, ok := value.(string); ok {
				// Clean the value
				v = strings.TrimSpace(v)
				v = removeDoubleSpaces(v)
				v = strings.ToLower(v)

				// Loop through the words
				for _, word := range strings.Split(v, " ") {
					if ft.maxWords != -1 {
						if len(ft.wordCache) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(ft.wordCache), ft.maxWords)
						}
					}
					if ft.maxSizeBytes != -1 {
						var cacheSize int = int(unsafe.Sizeof(ft.wordCache))
						if cacheSize > ft.maxSizeBytes {
							return fmt.Errorf("full text cache size limit reached (%d/%d bytes)", cacheSize, ft.maxSizeBytes)
						}
					}
					switch {
					case len(word) <= 1:
						continue
					case !isAlphaNum(word):
						word = removeNonAlphaNum(word)
					}
					if _, ok := ft.wordCache[word]; !ok {
						ft.wordCache[word] = []string{itemKey}
						continue
					}
					if containsString(ft.wordCache[word], itemKey) {
						continue
					}
					ft.wordCache[word] = append(ft.wordCache[word], itemKey)
				}
			}
		}
	}
	return nil
}
