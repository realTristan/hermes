package cache

import (
	"fmt"
	"strings"
)

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
func (ft *FullText) set(key string, value map[string]interface{}) error {
	var temp map[string][]string = ft.wordCache

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Clean the value
			v = strings.TrimSpace(v)
			v = removeDoubleSpaces(v)
			v = strings.ToLower(v)

			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if ft.maxWords > 0 {
					if len(temp) > ft.maxWords {
						return fmt.Errorf("full text cache key limit reached (%d/%d keys). set cancelled. cache reverted", len(temp), ft.maxWords)
					}
				}
				if ft.maxSizeBytes > 0 {
					if cacheSize, err := size(temp); err != nil {
						return err
					} else if cacheSize > ft.maxSizeBytes {
						return fmt.Errorf("full text cache size limit reached (%d/%d bytes). set cancelled. cache reverted", cacheSize, ft.maxSizeBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !isAlphaNum(word):
					word = removeNonAlphaNum(word)
				}
				if _, ok := temp[word]; !ok {
					temp[word] = []string{key}
					continue
				}
				temp[word] = append(temp[word], key)
			}
		}
	}

	// Set the word cache to the temp map
	ft.wordCache = temp

	// Return nil for no errors
	return nil
}
