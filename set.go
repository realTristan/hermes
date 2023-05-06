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
		return fmt.Errorf("full-text cache key already exists (%s). please delete it before setting it another value", key)
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
	// Create a copy of the existing full-text variables
	var (
		tempCache        map[string][]int = ft.cache
		tempIndices      map[int]string   = ft.indices
		tempCurrentIndex int              = ft.currentIndex
		tempKeys         map[string]int   = make(map[string]int)
	)

	// Loop through the json data
	for k, v := range tempIndices {
		tempKeys[v] = k
	}

	// Check if the key is in the temp keys map. If not, add it.
	if _, ok := tempKeys[key]; !ok {
		tempIndices[tempCurrentIndex] = key
		tempKeys[key] = tempCurrentIndex
		tempCurrentIndex++
	}

	// Loop through the provided value
	for _, v := range value {
		if strv, ok := v.(string); ok {
			// Clean the string value
			strv = strings.TrimSpace(strv)
			strv = Utils.RemoveDoubleSpaces(strv)
			strv = strings.ToLower(strv)

			// Loop through the words
			for _, word := range strings.Split(strv, " ") {
				if ft.maxWords > 0 {
					if len(tempCache) > ft.maxWords {
						return fmt.Errorf("full-text cache key limit reached (%d/%d keys). set cancelled. cache reverted", len(tempCache), ft.maxWords)
					}
				}
				if ft.maxBytes > 0 {
					if cacheSize, err := Utils.Size(tempCache); err != nil {
						return err
					} else if cacheSize > ft.maxBytes {
						return fmt.Errorf("full-text cache size limit reached (%d/%d bytes). set cancelled. cache reverted", cacheSize, ft.maxBytes)
					}
				}
				switch {
				case len(word) <= 1:
					continue
				case !Utils.IsAlphaNum(word):
					word = Utils.RemoveNonAlphaNum(word)
				}
				if _, ok := tempCache[word]; !ok {
					tempCache[word] = []int{tempCurrentIndex}
					continue
				}
				if Utils.ContainsInt(tempCache[word], tempKeys[key]) {
					continue
				}
				tempCache[word] = append(tempCache[word], tempKeys[key])
			}
		}
	}

	// Set the full-text cache to the temp map
	ft.cache = tempCache
	ft.indices = tempIndices
	ft.currentIndex = tempCurrentIndex

	// Return nil for no errors
	return nil
}
