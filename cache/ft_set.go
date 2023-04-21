package cache

import (
	"fmt"
	"strings"
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
func (c *Cache) FTSet(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return fmt.Errorf("full text is not initialized")
	}

	// Set the key-value pair in the cache
	if new, err := c.ft.set(key, value); err != nil {
		return err
	} else {
		c.ft = new
	}

	// Return no error
	return nil
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
	err = cache.ft.Set("example_key", data)
	if err != nil {
			log.Fatalf("Error setting FullText cache: %s", err)
	}
*/
func (ft *FullText) set(key string, value map[string]interface{}) (*FullText, error) {
	// Initialize a new full text object
	var temp *FullText = &FullText{
		maxWords:     ft.maxWords,
		maxSizeBytes: ft.maxSizeBytes,
		wordCache:    ft.wordCache,
		initialized:  ft.initialized,
	}

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Clean the value
			v = strings.TrimSpace(v)
			v = removeDoubleSpaces(v)
			v = strings.ToLower(v)

			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if temp.maxWords != -1 {
					if len(temp.wordCache) > temp.maxWords {
						return ft, fmt.Errorf("full text cache key limit reached (%d/%d keys). set cancelled. cache reverted", len(temp.wordCache), temp.maxWords)
					}
				}
				if temp.maxSizeBytes != -1 {
					if cacheSize, err := size(temp.wordCache); err != nil {
						return ft, err
					} else if cacheSize > temp.maxSizeBytes {
						return ft, fmt.Errorf("full text cache size limit reached (%d/%d bytes). set cancelled. cache reverted", cacheSize, temp.maxSizeBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !isAlphaNum(word):
					word = removeNonAlphaNum(word)
				}
				if _, ok := temp.wordCache[word]; !ok {
					temp.wordCache[word] = []string{key}
					continue
				}
				temp.wordCache[word] = append(temp.wordCache[word], key)
			}
		}
	}
	return temp, nil
}
