package hermes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"unsafe"
)

// Full Text struct
type FullText struct {
	mutex        *sync.RWMutex
	cache        map[string][]int
	words        []string
	data         []map[string]interface{}
	maxWords     int
	maxSizeBytes int
}

// Initialize the FT cache with a json file
func (c *Cache) InitFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	if c.FT != nil {
		return fmt.Errorf("FT cache already initialized")
	}

	// Initialize the FT cache
	c.FT = &FullText{
		mutex:        &sync.RWMutex{},
		cache:        map[string][]int{},
		words:        []string{},
		data:         []map[string]interface{}{},
		maxWords:     maxWords,
		maxSizeBytes: maxSizeBytes,
	}

	// Load the json cache
	c.FT.loadJson(file)
	if err := c.FT.loadCacheData(c.FT.data, schema); err != nil {
		return err
	}

	// Return the cache
	return nil
}

// Initialize the FT cache with Mutex Locking
func (c *Cache) InitFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.initFT(maxWords, maxSizeBytes, schema)
}

// Initialize the FT cache
func (c *Cache) initFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	if c.FT != nil {
		return fmt.Errorf("FT cache already initialized")
	}

	// Convert the cache data into an array of maps
	var data []map[string]interface{} = []map[string]interface{}{}
	for _, v := range c.data {
		data = append(data, v)
	}
	c.FT = &FullText{
		mutex:        &sync.RWMutex{},
		cache:        map[string][]int{},
		words:        []string{},
		data:         data,
		maxWords:     maxWords,
		maxSizeBytes: maxSizeBytes,
	}
	return c.FT.loadCacheData(data, schema)
}

// Set a value in the cache with Mutex Locking
func (ft *FullText) Set(key string, value map[string]interface{}) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.set(key, value)
}

// Set a value in the cache
func (ft *FullText) set(key string, value map[string]interface{}) error {
	ft.data = append(ft.data, value)

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if ft.maxWords != -1 {
					if len(ft.words) > ft.maxWords {
						return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(ft.words), ft.maxWords)
					}
				}
				if ft.maxSizeBytes != -1 {
					var cacheSize int = int(unsafe.Sizeof(ft.cache))
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
				if _, ok := ft.cache[word]; !ok {
					ft.cache[word] = []int{len(ft.data) - 1}
					ft.words = append(ft.words, word)
					continue
				}
				ft.cache[word] = append(ft.cache[word], len(ft.data)-1)
			}
		}
	}
	return nil
}

// Delete a value from the cache with Mutex Locking
func (ft *FullText) Delete(key string) {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.delete(key)
}

// Delete a word from the cache
func (ft *FullText) delete(key string) {
	// Remove the value from the keys
	for i, k := range ft.words {
		if k == key {
			ft.words = append(ft.words[:i], ft.words[i+1:]...)
			break
		}
	}

	// Remove the value from the cache
	delete(ft.cache, key)
}

// Clean the FullText cache with Mutex Locking
func (ft *FullText) Clean() {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.clean()
}

// Clean the FullText cache
func (ft *FullText) clean() {
	ft.cache = map[string][]int{}
	ft.words = []string{}
	ft.data = []map[string]interface{}{}
}

// Read the json data
func (ft *FullText) loadJson(file string) {
	var data, _ = os.ReadFile(file)
	json.Unmarshal(data, &ft.data)
}

// Load the FullText cache
func (ft *FullText) loadCacheData(data []map[string]interface{}, schema map[string]bool) error {
	// Loop through the json data
	for i, item := range data {
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
					if ft.maxWords != -1 {
						if len(ft.words) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(ft.words), ft.maxWords)
						}
					}
					if ft.maxSizeBytes != -1 {
						var cacheSize int = int(unsafe.Sizeof(ft.cache))
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
	}
	return nil
}
