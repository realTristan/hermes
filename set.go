package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// Set a value in the cache for the specified key.
// This function is thread-safe.
func (c *Cache) Set(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

// Set a value in the cache for the specified key.
// This function is not thread-safe, and should only be called from
// an exported function.
// If fullText is true, set the value in the full-text cache as well.
func (c *Cache) set(key string, value map[string]interface{}) error {
	if _, ok := c.data[key]; ok {
		return fmt.Errorf("full-text cache key already exists (%s). please delete it before setting it another value", key)
	}

	// Update the value in the FT cache
	if c.ft != nil {
		if err := c.ftSet(key, value); err != nil {
			return err
		}
		return nil
	}

	// Update the value in the cache
	c.data[key] = value

	// Return nil for no error
	return nil
}

// Set a value in the full-text cache for the specified key.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) ftSet(key string, value map[string]interface{}) error {
	// Create a copy of the existing full-text variables
	var (
		tempStorage      map[string][]int = c.ft.storage
		tempIndices      map[int]string   = c.ft.indices
		tempCurrentIndex int              = c.ft.currentIndex
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
	for k, v := range value {
		var strv string
		// Check if the value is a WFT
		if wft, ok := v.(WFT); ok {
			strv = wft.value
		} else if _strv := fullTextMap(v); len(_strv) > 0 {
			strv = _strv
		} else {
			continue
		}

		// Set the key in the provided value to the string wft value
		value[k] = strv

		// Clean the string value
		strv = strings.TrimSpace(strv)
		strv = Utils.RemoveDoubleSpaces(strv)
		strv = strings.ToLower(strv)

		// Loop through the words
		for _, word := range strings.Split(strv, " ") {
			if c.ft.maxWords > 0 {
				if len(tempStorage) > c.ft.maxWords {
					return fmt.Errorf("full-text cache key limit reached (%d/%d keys). set cancelled. cache reverted", len(tempStorage), c.ft.maxWords)
				}
			}
			if c.ft.maxBytes > 0 {
				if cacheSize, err := Utils.Size(tempStorage); err != nil {
					return err
				} else if cacheSize > c.ft.maxBytes {
					return fmt.Errorf("full-text cache size limit reached (%d/%d bytes). set cancelled. cache reverted", cacheSize, c.ft.maxBytes)
				}
			}
			switch {
			case len(word) <= 1:
				continue
			case !Utils.IsAlphaNum(word):
				word = Utils.RemoveNonAlphaNum(word)
			}
			if _, ok := tempStorage[word]; !ok {
				tempStorage[word] = []int{tempKeys[key]}
				continue
			}
			if Utils.ContainsInt(tempStorage[word], tempKeys[key]) {
				continue
			}
			tempStorage[word] = append(tempStorage[word], tempKeys[key])
		}
	}

	// Set the full-text cache to the temp map
	c.ft.storage = tempStorage
	c.ft.indices = tempIndices
	c.ft.currentIndex = tempCurrentIndex

	// Update the cache
	c.data[key] = value

	// Return nil for no errors
	return nil
}
