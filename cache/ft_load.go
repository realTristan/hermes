package cache

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
loadCacheData() loops through a map of data and extracts relevant information to populate the wordCache map in the FullText struct.

Parameters:
- data (map[string]map[string]interface{}): a map of data to be loaded into the cache
- schema (map[string]bool): a map representing the schema of the data; only keys in this map will be loaded into the cache

Returns:
- error: if an error occurs during the loading process, it is returned. Otherwise, returns nil.
*/
func (ft *FullText) loadCacheData(data map[string]map[string]interface{}, schema map[string]bool) error {
	// Loop through the json data
	var i int = 0
	for _, itemValue := range data {
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
				v = removeDoubleSpaces(v)
				v = strings.ToLower(v)

				// Loop through the words
				for _, word := range strings.Split(v, " ") {
					if ft.maxWords != -1 {
						if len(ft.keys) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(ft.keys), ft.maxWords)
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
						ft.wordCache[word] = []int{i}
						continue
					}
					if containsInt(ft.wordCache[word], i) {
						continue
					}
					ft.wordCache[word] = append(ft.wordCache[word], i)
				}
			}
		}
		i++
	}
	return nil
}
