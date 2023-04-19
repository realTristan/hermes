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

// Initialize the FT cache with a json file with Mutex Locking
func (c *Cache) InitFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
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
		wordCache:     map[string][]string{},
		data:          map[string]map[string]interface{}{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
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
		if _, ok := c.FT.data[k]; ok {
			c.FT.clean()
			return fmt.Errorf("the key: %s has already been imported via json file. unable to add it to the cache", k)
		}
		c.FT.data[k] = v
	}

	// Load the cache data
	if err := c.FT.loadCacheData(c.FT.data, schema); err != nil {
		c.FT.clean()
		return err
	}

	// Set the FT cache as initialized
	c.FT.isInitialized = true

	// Return the cache
	return nil
}

// Initialize the FT cache with Mutex Locking
func (c *Cache) InitFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
	return c.initFT(maxWords, maxSizeBytes, schema)
}

// Initialize the FT cache
func (c *Cache) initFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is already initialized, return an error
	if c.FT != nil && c.FT.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}

	// Initialize the FT struct
	c.FT = &FullText{
		mutex:         &sync.RWMutex{},
		wordCache:     map[string][]string{},
		data:          map[string]map[string]interface{}{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
	}

	// Convert the cache data into an array of maps
	for k, v := range c.data {
		c.FT.data[k] = v
	}

	// Load the cache data
	if err := c.FT.loadCacheData(c.FT.data, schema); err != nil {
		c.FT.clean()
		return err
	}

	// Set the FT cache as initialized
	c.FT.isInitialized = true

	// Return the cache
	return nil
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
	ft.wordCache = map[string][]string{}
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
					ft.wordCache[word] = []string{key}
					continue
				}
				ft.wordCache[word] = append(ft.wordCache[word], key)
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

	// Remove the key from the ft.wordCache
	for word, keys := range ft.wordCache {
		for i := 0; i < len(keys); i++ {
			if key != keys[i] {
				continue
			}

			// Remove the key from the ft.wordCache slice
			ft.wordCache[word] = append(ft.wordCache[word][:i], ft.wordCache[word][i+1:]...)
		}

		// If the ft.wordCache[word] is empty, remove it from the cache
		if len(ft.wordCache[word]) == 0 {
			delete(ft.wordCache, word)
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
