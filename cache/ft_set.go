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
	err := cache.FT.Set("document123", map[string]interface{}{"title": "Document Title", "body": "Document Body"})
	if err != nil {
		log.Fatal(err)
	}
*/
func (c *Cache) SetFT(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.setFT(key, value)
}

/*
The Set function adds a key-value pair to the FullText cache. The cache is locked with a mutex to ensure thread-safety.
If the key already exists, an error is returned.
If the maxWords or maxSizeBytes are exceeded, an error is returned.
The value is cleaned and tokenized before being added to the cache.

@Parameters:

	key (string): The key to be added to the cache
	value (map[string]interface{}): The value to be added to the cache

@Returns:

	error: If an error occurs, it is returned. If the operation was successful, nil is returned.

Example usage:

	cache := InitCache()
	schema := map[string]bool{
		"content": true,
		"title":   true,
	}
	maxWords := 1000
	maxSizeBytes := 1024 * 1024 // 1MB

	// Add data to cache
	err := cache.ResetFT(maxWords, maxSizeBytes, schema)
	if err != nil {
			log.Fatalf("Error resetting FullText cache: %s", err)
	}

	// Set a key-value pair in the cache
	data := map[string]interface{}{
			"content": "This is some example content",
			"title":   "Example",
	}
	err = cache.FT.Set("example_key", data)
	if err != nil {
			log.Fatalf("Error setting FullText cache: %s", err)
	}
*/
func (c *Cache) setFT(key string, value map[string]interface{}) error {
	// If the key already exists, return an error
	if _, ok := c.data[key]; ok {
		return fmt.Errorf("full text cache key already exists (%s). please delete it before setting it another value", key)
	}

	// Add the value to the cache
	c.data[key] = value

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Clean the value
			v = strings.TrimSpace(v)
			v = removeDoubleSpaces(v)
			v = strings.ToLower(v)

			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if c.FT.maxWords != -1 {
					if len(c.FT.wordCache) > c.FT.maxWords {
						return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(c.FT.wordCache), c.FT.maxWords)
					}
				}
				if c.FT.maxSizeBytes != -1 {
					var cacheSize int = int(unsafe.Sizeof(c.FT.wordCache))
					if cacheSize > c.FT.maxSizeBytes {
						return fmt.Errorf("full text cache size limit reached (%d/%d bytes)", cacheSize, c.FT.maxSizeBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !isAlphaNum(word):
					word = removeNonAlphaNum(word)
				}
				if _, ok := c.FT.wordCache[word]; !ok {
					c.FT.wordCache[word] = []string{key}
					continue
				}
				c.FT.wordCache[word] = append(c.FT.wordCache[word], key)
			}
		}
	}
	return nil
}
