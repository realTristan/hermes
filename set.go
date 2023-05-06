package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// Set a value in the cache for the specified key.
// This function is thread-safe.
func (c *Cache) Set(key string, value map[string]interface{}, fullText bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value, fullText)
}

// Set a value in the cache for the specified key.
// This function is not thread-safe, and should only be called from
// an exported function.
// If fullText is true, set the value in the full-text cache as well.
func (c *Cache) set(key string, value map[string]interface{}, fullText bool) error {
	if _, ok := c.data[key]; ok {
		return fmt.Errorf("full text cache key already exists (%s). please delete it before setting it another value", key)
	}

	// Update the value in the FT cache
	if fullText && c.ft != nil {
		if err := c.ft.set(key, value); err != nil {
			return err
		}
	}

	// Update the value in the cache
	c.data[key] = value

	// Return nil for no error
	return nil
}

// Set a value in the full-text cache for the specified key.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) set(key string, value map[string]interface{}) error {
	// Create a copy of the existing word cache
	var (
		tempWordCache map[string][]int = ft.wordCache
		tempIndices   map[int]string   = ft.indicesCache
		tempCacheSize int              = ft.cacheSize
		tempKeys      map[string]int   = make(map[string]int)
	)

	// Loop through the json data
	for k, v := range tempIndices {
		tempKeys[v] = k
	}

	// Loop through the value
	for _, _v := range value {
		if v, ok := _v.(string); ok {
			// Clean the value
			v = strings.TrimSpace(v)
			v = Utils.RemoveDoubleSpaces(v)
			v = strings.ToLower(v)

			// Loop through the words
			for _, word := range strings.Split(v, " ") {
				if ft.maxWords > 0 {
					if len(tempWordCache) > ft.maxWords {
						return fmt.Errorf("full text cache key limit reached (%d/%d keys). set cancelled. cache reverted", len(tempWordCache), ft.maxWords)
					}
				}
				if ft.maxBytes > 0 {
					if cacheSize, err := Utils.Size(tempWordCache); err != nil {
						return err
					} else if cacheSize > ft.maxBytes {
						return fmt.Errorf("full text cache size limit reached (%d/%d bytes). set cancelled. cache reverted", cacheSize, ft.maxBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !Utils.IsAlphaNum(word):
					word = Utils.RemoveNonAlphaNum(word)
				}
				if _, ok := tempWordCache[word]; !ok {
					tempWordCache[word] = []int{tempCacheSize}
					continue
				}
				if v, ok := tempKeys[key]; !ok {
					tempIndices[tempCacheSize] = key
					tempKeys[key] = tempCacheSize
					tempCacheSize++
				} else {
					if Utils.ContainsInt(tempWordCache[word], v) {
						continue
					}
					tempWordCache[word] = append(tempWordCache[word], v)
				}
			}
		}
	}

	// Set the word cache to the temp map
	ft.wordCache = tempWordCache
	ft.indicesCache = tempIndices
	ft.cacheSize = tempCacheSize

	// Return nil for no errors
	return nil
}
