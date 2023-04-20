package cache

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
Set a value in the FullText cache with Mutex Locking

This method locks the FullText cache mutex, sets the given value for the specified key in the cache, and then unlocks the mutex.

@Parameters:

	key: A string representing the key for the cache entry
	value: A map[string]interface{} representing the value to be stored in the cache

@Returns:

	error: Returns an error if the FullText cache is not initialized or the cache size limit is reached

Example:

	cache := InitCache()
	cache.InitFT(100, 1000000, map[string]bool{"title": true, "body": true})
	err := cache.ft.Set("document123", map[string]interface{}{"title": "Document Title", "body": "Document Body"})
	if err != nil {
		log.Fatal(err)
	}
*/
func (c *Cache) SetFT(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil || !c.ft.isInitialized {
		return fmt.Errorf("full text is not initialized")
	}
	return c.setFT(key, value)
}

/*
setFT sets the value of a key in the cache and updates the full text search index.

Parameters:
- key: the key to set.
- value: the value to set.

Returns an error if the key already exists in the cache or if the full text cache limit has been reached.
*/
func (c *Cache) setFT(key string, value map[string]interface{}) error {
	// If the key already exists, return an error
	if _, ok := c.data[key]; ok {
		return fmt.Errorf("full text cache key already exists (%s). please delete it before setting it another value", key)
	}

	// Add the value to the cache
	c.data[key] = value

	// Add the provided key to the keys slice
	c.ft.keys = append(c.ft.keys, key)
	var index int = len(c.ft.keys) - 1

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Clean the value
			v = strings.TrimSpace(v)
			v = removeDoubleSpaces(v)
			v = strings.ToLower(v)

			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if c.ft.maxWords != -1 {
					if len(c.ft.keys) > c.ft.maxWords {
						return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(c.ft.keys), c.ft.maxWords)
					}
				}
				if c.ft.maxSizeBytes != -1 {
					var cacheSize int = int(unsafe.Sizeof(c.ft.wordCache))
					if cacheSize > c.ft.maxSizeBytes {
						return fmt.Errorf("full text cache size limit reached (%d/%d bytes)", cacheSize, c.ft.maxSizeBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !isAlphaNum(word):
					word = removeNonAlphaNum(word)
				}
				if _, ok := c.ft.wordCache[word]; !ok {
					c.ft.wordCache[word] = []int{index}
					continue
				}
				c.ft.wordCache[word] = append(c.ft.wordCache[word], index)
			}
		}
	}
	return nil
}
