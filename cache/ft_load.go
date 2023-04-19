package cache

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
Load the FullText cache with the provided map data and schema. This function populates the FullText word cache with the given map data,
based on the provided schema. Each item in the map represents a document or piece of content, and its keys are the field names while
the values are the corresponding field values. The schema parameter is a map of allowed field names. Only field names listed in this
map will be indexed in the FullText cache. Any field not listed in the schema will be ignored. The cache is a map where the keys
are individual words, and the values are lists of item keys that contain  that word. When a word is encountered, it is cleaned,
normalized and then added to the word cache. If the  word is already present in the cache, the corresponding item key is added to
the list. If the cache already contains the item key for the word, it will be ignored.

@param data: a map representing the data to be indexed.
The keys represent item keys while the values are  maps containing the actual field names and their corresponding values.

@param schema: a map representing the field names allowed to be indexed in the FullText cache. Only the
field names listed in this map will be added to the cache. Any field not listed in the schema will be ignored.

@return error if any errors occurred during the caching process. Otherwise, it returns nil.

Example usage:

	...
	data := make(map[string]map[string]interface{
		"item1": map[string]interface{}{"title": "This is a title", "body": "This is a long description of the item."},
		"item2": map[string]interface{}{"title": "Another title", "body": "This is another long description."},
	})
	schema := map[string]bool{"title": true, "body": true}
	if err := ft.loadCacheData(data, schema); err != nil {
		fmt.Printf("Error loading cache data: %s\n", err.Error())
	}
*/
func (ft *FullText) loadCacheData(data map[string]map[string]interface{}, schema map[string]bool) error {
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
					if ft.maxWords != -1 {
						if len(ft.wordCache) > ft.maxWords {
							return fmt.Errorf("full text cache key limit reached (%d/%d keys)", len(ft.wordCache), ft.maxWords)
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
						ft.wordCache[word] = []string{itemKey}
						continue
					}
					if containsString(ft.wordCache[word], itemKey) {
						continue
					}
					ft.wordCache[word] = append(ft.wordCache[word], itemKey)
				}
			}
		}
	}
	return nil
}
