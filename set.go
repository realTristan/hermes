package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// Set is a method of the Cache struct that sets a value in the cache for the specified key.
// This function is thread-safe.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]interface{} representing the value to set.
//
// Returns:
//   - An error if the full-text cache key already exists. Otherwise, nil.
func (c *Cache) Set(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

// set is a method of the Cache struct that sets a value in the cache for the specified key.
// This function is not thread-safe, and should only be called from an exported function.
// If fullText is true, set the value in the full-text cache as well.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]interface{} representing the value to set.
//
// Returns:
//   - An error if the full-text cache key already exists. Otherwise, nil.
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

// ftSet is a method of the Cache struct that sets a value in the full-text cache for the specified key.
// This function is not thread-safe, and should only be called from an exported function.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]interface{} representing the value to set.
//
// Returns:
//   - An error if the full-text storage limit or byte-size limit is reached. Otherwise, nil.
func (c *Cache) ftSet(key string, value map[string]interface{}) error {
	// Create a copy of the existing full-text variables
	var (
		tempStorage      map[string]interface{} = c.ft.storage
		tempIndices      map[int]string         = c.ft.indices
		tempCurrentIndex int                    = c.ft.currentIndex
		tempKeys         map[string]int         = make(map[string]int)
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
		} else if _strv := ftFromMap(v); len(_strv) > 0 {
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
			if len(word) == 0 {
				continue
			}

			// Check if the storage limit has been reached
			if c.ft.maxLength > 0 {
				if len(tempStorage) > c.ft.maxLength {
					return fmt.Errorf("full-text storage limit reached (%d/%d keys). load cancelled", len(tempStorage), c.ft.maxLength)
				}
			}
			if c.ft.maxBytes > 0 {
				if cacheSize, err := Utils.Size(tempStorage); err != nil {
					return err
				} else if cacheSize > c.ft.maxBytes {
					return fmt.Errorf("full-text byte-size limit reached (%d/%d bytes). load cancelled", cacheSize, c.ft.maxBytes)
				}
			}

			// Trim the word
			word = Utils.TrimNonAlphaNum(word)
			var words []string = Utils.SplitByAlphaNum(word)

			// Loop through the words
			for i := 0; i < len(words); i++ {
				if len(words[i]) <= 3 {
					continue
				}
				if _, ok := tempStorage[words[i]]; !ok {
					tempStorage[words[i]] = []int{tempCurrentIndex}
					continue
				}
				if _, ok := tempStorage[words[i]].([]int); ok {
					if Utils.ContainsInt(tempStorage[words[i]].([]int), tempKeys[key]) {
						continue
					}
					tempStorage[words[i]] = append(tempStorage[words[i]].([]int), tempKeys[key])
				} else {
					tempStorage[words[i]] = []int{tempStorage[words[i]].(int), tempKeys[key]}
				}
			}
		}
	}

	// Iterate over the temp storage and set the values with len 1 to int
	for k, v := range tempStorage {
		if v, ok := v.([]int); ok && len(v) == 1 {
			tempStorage[k] = v[0]
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
