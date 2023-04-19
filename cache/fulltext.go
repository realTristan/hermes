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
	cache         map[string][]string
	words         []string
	data          map[string]map[string]interface{}
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}

// Initialize the FT cache with a json file with Mutex Locking
func (c *Cache) InitFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}

// Initialize the FT cache with a json file
func (c *Cache) initFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is already initialized, return an error
	if c.FT != nil && c.FT.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}

	// Initialize the FT cache
	c.FT = &FullText{
		mutex:         &sync.RWMutex{},
		cache:         map[string][]string{},
		words:         []string{},
		data:          map[string]map[string]interface{}{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: true,
	}

	// Read the json
	if data, err := readJson(file); err != nil {
		return err
	} else {
		c.FT.data = data
	}

	// Iterate over the regular cache data and
	// add it to the FT cache data
	for k, v := range c.data {
		c.FT.data[k] = v
	}

	// Load the cache data
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
	// If the FT cache is already initialized, return an error
	if c.FT != nil && c.FT.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}

	// Convert the cache data into an array of maps
	var data map[string]map[string]interface{} = map[string]map[string]interface{}{}
	for k, v := range c.data {
		data[k] = v
	}
	c.FT = &FullText{
		mutex:         &sync.RWMutex{},
		cache:         map[string][]string{},
		words:         []string{},
		data:          data,
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: true,
	}
	return c.FT.loadCacheData(data, schema)
}

// Reset the FullText cache with Mutex Locking
func (c *Cache) ResetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFT(maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache
func (c *Cache) resetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.FT == nil {
		return fmt.Errorf("full text cache not initialized")
	}

	// Lock the FT cache
	c.FT.mutex.Lock()
	defer c.FT.mutex.Unlock()

	// If the FT cache is not initialized, return an error
	if !c.FT.isInitialized {
		return fmt.Errorf("full text cache not initialized")
	}

	// Reset the FT cache
	c.FT.isInitialized = false
	return c.initFT(maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache with Json and Mutex Locking
func (c *Cache) ResetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFTJson(file, maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache with Json
func (c *Cache) resetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.FT == nil {
		return fmt.Errorf("FT cache not initialized")
	}

	// Lock the FT cache
	c.FT.mutex.Lock()
	defer c.FT.mutex.Unlock()

	// If the FT cache is not initialized, return an error
	if !c.FT.isInitialized {
		return fmt.Errorf("FT cache not initialized")
	}

	// Reset the FT cache
	c.FT.isInitialized = false
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}

// Clean the FullText cache with Mutex Locking
func (ft *FullText) Clean() {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.clean()
}

// Clean the FullText cache
func (ft *FullText) clean() {
	ft.cache = map[string][]string{}
	ft.words = []string{}
	ft.data = map[string]map[string]interface{}{}
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

// Set a value in the cache with Mutex Locking
func (ft *FullText) Set(key string, value map[string]interface{}) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.set(key, value)
}

// Set a value in the cache
func (ft *FullText) set(key string, value map[string]interface{}) error {
	// If the key already exists, return an error
	if _, ok := ft.data[key]; ok {
		return fmt.Errorf("full text cache key already exists (%s)", key)
	}

	// Add the value to the cache
	ft.data[key] = value

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
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
					ft.cache[word] = []string{key}
					ft.words = append(ft.words, word)
					continue
				}
				ft.cache[word] = append(ft.cache[word], key)
			}
		}
	}
	return nil
}

// Delete a value from the cache with Mutex Locking
func (ft *FullText) Delete(word string) {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.delete(word)
}

// Delete a word from the cache
func (ft *FullText) delete(key string) {
	// Remove the key from the ft.data
	delete(ft.data, key)

	// Remove the key from the ft.cache
	for word, keys := range ft.cache {
		for i, _key := range keys {
			if key != _key {
				continue
			}

			// Remove the key from the ft.cache slice
			ft.cache[word] = append(ft.cache[word][:i], ft.cache[word][i+1:]...)
		}

		// If the ft.cache slice is empty, remove the word from the ft.cache
		if len(ft.cache[word]) == 0 {
			delete(ft.cache, word)

			// Remove the word from the ft.words
			for i, _word := range ft.words {
				if word != _word {
					continue
				}
				ft.words = append(ft.words[:i], ft.words[i+1:]...)
				break
			}
		}
	}
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
						ft.cache[word] = []string{itemKey}
						ft.words = append(ft.words, word)
						continue
					}
					if containsString(ft.cache[word], itemKey) {
						continue
					}
					ft.cache[word] = append(ft.cache[word], itemKey)
				}
			}
		}
	}
	return nil
}
