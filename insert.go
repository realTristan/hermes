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
	for itemKey, itemValue := range data {
		// Loop through the map
		for key, value := range itemValue {
			// Check if the key is in the schema
			if !schema[key] {
				continue
			}

			// Check if the value is a string
			if v, ok := value.(string); ok {
				// Clean the value
				v = strings.TrimSpace(v)
				v = Utils.RemoveDoubleSpaces(v)
				v = strings.ToLower(v)

				// Loop through the words
				for _, word := range strings.Split(v, " ") {
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
					if v, ok := tempKeys[itemKey]; !ok {
						tempIndices[tempCurrentIndex] = itemKey
						tempKeys[itemKey] = tempCurrentIndex
						tempCurrentIndex++
					} else {
						if Utils.ContainsInt(tempCache[word], v) {
							continue
						}
						tempCache[word] = append(tempCache[word], v)
					}
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
