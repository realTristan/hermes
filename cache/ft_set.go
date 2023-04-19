package cache

import (
	"fmt"
	"strings"
	"unsafe"
)

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
