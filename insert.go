package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// Insert a value in the full-text cache for the specified key.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) insert(data map[string]map[string]interface{}, schema map[string]bool) error {
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

	// Loop through the json data
	for cacheKey, cacheValue := range data {
		// Check if the key is in the temp keys map. If not, add it.
		if _, ok := tempKeys[cacheKey]; !ok {
			tempIndices[tempCurrentIndex] = cacheKey
			tempKeys[cacheKey] = tempCurrentIndex
			tempCurrentIndex++
		}

		// Loop through the map
		for k, v := range cacheValue {
			// Check if the key is in the schema
			if !schema[k] {
				continue
			}

			// Check if the value is a string
			if strv, ok := v.(string); ok {
				// Clean the string value
				strv = strings.TrimSpace(strv)
				strv = Utils.RemoveDoubleSpaces(strv)
				strv = strings.ToLower(strv)

				// Loop through the words
				for _, word := range strings.Split(strv, " ") {
					if ft.maxWords > 0 {
						if len(tempCache) > ft.maxWords {
							return fmt.Errorf("full-text cache key limit reached (%d/%d keys). load cancelled", len(tempCache), ft.maxWords)
						}
					}
					if ft.maxBytes > 0 {
						if cacheSize, err := Utils.Size(tempCache); err != nil {
							return err
						} else if cacheSize > ft.maxBytes {
							return fmt.Errorf("full-text cache size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxBytes)
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
					if Utils.ContainsInt(tempCache[word], tempKeys[cacheKey]) {
						continue
					}
					tempCache[word] = append(tempCache[word], tempKeys[cacheKey])
				}
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
