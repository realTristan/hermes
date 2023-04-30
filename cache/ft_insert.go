package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

/*
insertIntoWordCache() loops through a map of data and extracts relevant information to populate the wordCache map in the FullText struct.

Parameters:
- data (map[string]map[string]interface{}): a map of data to be loaded into the cache
- schema (map[string]bool): a map representing the schema of the data; only keys in this map will be loaded into the cache

Returns:
- error: if an error occurs during the loading process, it is returned. Otherwise, returns nil.
*/
func (ft *FullText) insertIntoWordCache(data map[string]map[string]interface{}, schema map[string]bool) error {
	// Create a copy of the existing word cache
	var temp map[string][]string = ft.wordCache

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
						if len(temp) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys). load cancelled", len(temp), ft.maxWords)
						}
					}
					if ft.maxSizeBytes > 0 {
						if cacheSize, err := Utils.Size(temp); err != nil {
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
					if _, ok := temp[word]; !ok {
						temp[word] = []string{itemKey}
						continue
					}
					if Utils.ContainsString(temp[word], itemKey) {
						continue
					}
					temp[word] = append(temp[word], itemKey)
				}
			}
		}
	}

	// Set the word cache to the temp map
	ft.wordCache = temp

	// Return nil for no errors
	return nil
}
