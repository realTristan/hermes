package hermes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"unsafe"
)

// Full Text Search struct
type FTS struct {
	mutex        *sync.RWMutex
	cache        map[string][]int
	keys         []string
	json         []map[string]interface{}
	maxKeys      int
	maxSizeBytes int
}

// InitCache function
func InitJson(file string, maxKeys int, maxSizeBytes int, schema map[string]bool) (*FTS, error) {
	var fts *FTS = &FTS{
		mutex:        &sync.RWMutex{},
		cache:        map[string][]int{},
		keys:         []string{},
		json:         []map[string]interface{}{},
		maxKeys:      maxKeys,
		maxSizeBytes: maxSizeBytes,
	}

	// Load the json cache
	fts.loadJson(file)
	if err := fts.loadCacheJson(fts.json, schema); err != nil {
		return nil, err
	}

	// Return the cache
	return fts, nil
}

// Set a value in the cache with Mutex Locking
func (fts *FTS) Set(key string, value map[string]interface{}) error {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	return fts.set(key, value)
}

// Set a value in the cache
func (fts *FTS) set(key string, value map[string]interface{}) error {
	fts.json = append(fts.json, value)
	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if fts.maxKeys != -1 {
					if len(fts.keys) > fts.maxKeys {
						return fmt.Errorf("fts cache key limit reached (%d/%d keys)", len(fts.keys), fts.maxKeys)
					}
				}
				if fts.maxSizeBytes != -1 {
					var cacheSize int = int(unsafe.Sizeof(fts.cache))
					if cacheSize > fts.maxSizeBytes {
						return fmt.Errorf("fts cache size limit reached (%d/%d bytes)", cacheSize, fts.maxSizeBytes)
					}
				}
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
	return nil
}

// Delete a word from the cache with Mutex Locking
func (fts *FTS) Delete(key string) {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts.delete(key)
}

// Delete a word from the cache
func (fts *FTS) delete(key string) {
	// Remove the value from the keys
	for i, k := range fts.keys {
		if k == key {
			fts.keys = append(fts.keys[:i], fts.keys[i+1:]...)
			break
		}
	}

	// Remove the value from the cache
	delete(fts.cache, key)
}

// Clean the FTS cache with Mutex Locking
func (fts *FTS) Clean() {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	fts.clean()
}

// Clean the FTS cache
func (fts *FTS) clean() {
	fts.cache = map[string][]int{}
	fts.keys = []string{}
	fts.json = []map[string]interface{}{}
}

// Reset the FTS cache with Mutex Locking
func (fts *FTS) Reset(file string, schema map[string]bool) error {
	fts.mutex.Lock()
	defer fts.mutex.Unlock()
	return fts.reset(file, schema)
}

// Reset the FTS cache
func (fts *FTS) reset(file string, schema map[string]bool) error {
	var newFts, err = InitJson(file, fts.maxKeys, fts.maxSizeBytes, schema)
	if err != nil {
		return err
	}
	*fts = *newFts
	return nil
}

// Read the json data
func (fts *FTS) loadJson(file string) {
	var data, _ = os.ReadFile(file)
	json.Unmarshal(data, &fts.json)
}

// Load the FTS cache
func (fts *FTS) loadCacheJson(json []map[string]interface{}, schema map[string]bool) error {
	// Loop through the json data
	for i, item := range json {
		// Loop through the map
		for key, value := range item {
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
					if fts.maxKeys != -1 {
						if len(fts.keys) > fts.maxKeys {
							return fmt.Errorf("fts cache key limit reached (%d/%d keys)", len(fts.keys), fts.maxKeys)
						}
					}
					if fts.maxSizeBytes != -1 {
						var cacheSize int = int(unsafe.Sizeof(fts.cache))
						if cacheSize > fts.maxSizeBytes {
							return fmt.Errorf("fts cache size limit reached (%d/%d bytes)", cacheSize, fts.maxSizeBytes)
						}
					}
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
	}
	return nil
}
