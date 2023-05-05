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
						if len(tempWordCache) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys). load cancelled", len(tempWordCache), ft.maxWords)
						}
					}
					if ft.maxSizeBytes > 0 {
						if cacheSize, err := Utils.Size(tempWordCache); err != nil {
							return err
						} else if cacheSize > ft.maxSizeBytes {
							return fmt.Errorf("full text cache size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxSizeBytes)
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
					if v, ok := tempKeys[itemKey]; !ok {
						tempIndices[tempCacheSize] = itemKey
						tempKeys[itemKey] = tempCacheSize
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
	}

	// Set the word cache to the temp map
	ft.wordCache = tempWordCache
	ft.indicesCache = tempIndices
	ft.cacheSize = tempCacheSize

	// Return nil for no errors
	return nil
}
