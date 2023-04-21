package cache

import (
	"fmt"
	"strings"
)

/*
loadCacheData() loops through a map of data and extracts relevant information to populate the wordCache map in the FullText struct.

Parameters:
- data (map[string]map[string]interface{}): a map of data to be loaded into the cache
- schema (map[string]bool): a map representing the schema of the data; only keys in this map will be loaded into the cache

Returns:
- error: if an error occurs during the loading process, it is returned. Otherwise, returns nil.
*/
func (ft *FullText) loadCache(data map[string]map[string]interface{}, schema map[string]bool) (*FullText, error) {
	var temp *FullText = &FullText{
		wordCache:    make(map[string][]string, ft.maxWords),
		maxWords:     ft.maxWords,
		maxSizeBytes: ft.maxSizeBytes,
		initialized:  ft.initialized,
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
				v = removeDoubleSpaces(v)
				v = strings.ToLower(v)

				// Loop through the words
				for _, word := range strings.Split(v, " ") {
					if temp.maxWords != -1 {
						if len(temp.wordCache) > temp.maxWords {
							return ft, fmt.Errorf("full text cache key limit reached (%d/%d keys). load cancelled", len(temp.wordCache), temp.maxWords)
						}
					}
					if temp.maxSizeBytes != -1 {
						if cacheSize, err := size(temp.wordCache); err != nil {
							return ft, err
						} else if cacheSize > temp.maxSizeBytes {
							return ft, fmt.Errorf("full text cache size limit reached (%d/%d bytes). load cancelled", cacheSize, temp.maxSizeBytes)
						}
					}
					switch {
					case len(word) <= 1:
						continue
					case !isAlphaNum(word):
						word = removeNonAlphaNum(word)
					}
					if _, ok := temp.wordCache[word]; !ok {
						temp.wordCache[word] = []string{itemKey}
						continue
					}
					if containsString(temp.wordCache[word], itemKey) {
						continue
					}
					temp.wordCache[word] = append(temp.wordCache[word], itemKey)
				}
			}
		}
	}
	return temp, nil
}
