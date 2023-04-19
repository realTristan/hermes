package cache

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

/*
The FullText struct represents a full text search cache that stores words and the keys of the documents where they appear.
It uses a read-write mutex to ensure safe access to the cache data, a map to store the words and their associated keys, and a map to store the document data.
The maxWords and maxSizeBytes fields represent optional limits on the size of the cache. If either limit is reached, additional data will not be added to the cache.
When initialized, the isInitialized field is set to true to indicate that the cache is ready for use.

@Variables:

	mutex *sync.RWMutex: a reader/writer mutex used to synchronize access to the wordCache and data maps when reading and writing to them.

	wordCache map[string][]string: a map of words to a slice of keys for items that contain the word. The keys represent the unique identifiers for each item in the cache.

	data map[string]map[string]interface{}: a map of keys to a map of attributes and their corresponding values. The keys represent the unique identifiers for each item in the cache.

	maxWords int: the maximum number of unique keys that can be stored in the wordCache map. If the cache size exceeds this value, an error is returned when attempting to add a new item.

	maxSizeBytes int: the maximum size of the wordCache map in bytes. If the cache size exceeds this value, an error is returned when attempting to add a new item.

	isInitialized bool: a flag indicating whether the cache has been initialized or not. If the cache is not initialized, attempting to read or write to it will result in an error.
*/
type FullText struct {
	mutex         *sync.RWMutex
	wordCache     map[string][]string
	data          map[string]map[string]interface{}
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}

/*
Uploads JSON data to the FullText cache with Mutex Locking.

@Parameters:

	file: string, the path to the JSON file to upload to the cache.
	schema: map[string]bool, a map representing the schema of the JSON file where the key is the field name and the value is true if the field should be indexed in the cache, false otherwise.

@Returns:

	error, if there is an error uploading the JSON file to the cache.

Note:

	The JSON file must be an array of JSON objects where each object represents a key-value pair to add to the cache.
	The "id" field is reserved and should not be used as a field name in the schema.
*/
func (ft *FullText) UploadJson(file string, schema map[string]bool) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.uploadJson(file, schema)
}

/*
Uploads data from a json file to the FullText cache. The function reads the json file and loads the data into the cache.

@Arguments:
  - file (string): The path to the json file to read.
  - schema (map[string]bool): A map containing the schema definition for the data. The keys of the map represent the field names in the data and the values represent whether the field is indexed or not.

@Returns:
  - error: An error object if there was an error reading the json file or loading the data into the cache, or nil otherwise.

Example Usage:

	err := ft.uploadJson("data.json", map[string]bool{"title": true, "content": true})
	if err != nil {
		log.Fatalf("Error uploading json data to FullText cache: %v", err)
	}
*/
func (ft *FullText) uploadJson(file string, schema map[string]bool) error {
	// Read the json
	if data, err := readJson(file); err != nil {
		return err
	} else {
		return ft.loadCacheData(data, schema)
	}
}

/*
Upload map data to the FullText cache with Mutex Locking.

This function accepts a map of map[string]interface{} data and a schema, and uploads the data to the FullText cache.

@Parameters:
  - data: a map[string]map[string]interface{} containing the data to be uploaded
  - schema: a map[string]bool representing the schema to be used for the uploaded data

@Returns:
  - error: an error object if an error occurs during the upload process, or nil if the upload is successful

Example Usage:

	data := map[string]map[string]interface{}{
			"1": map[string]interface{}{"text": "hello world"},
			"2": map[string]interface{}{"text": "foo bar"},
	}
	schema := map[string]bool{"text": true}
	err := ft.UploadMap(data, schema)
*/
func (ft *FullText) UploadMap(data map[string]map[string]interface{}, schema map[string]bool) error {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	return ft.uploadMap(data, schema)
}

/*
Uploads map data to the FullText cache.

@Arguments:

	data: a map of maps with string keys and interfaces as values, representing the data to be uploaded.
	schema: a map with string keys and boolean values, representing the schema of the data to be uploaded.

@Returns:

	an error if the data could not be uploaded to the cache.

This function calls the 'loadCacheData' function, passing the provided data and schema as arguments.
*/
func (ft *FullText) uploadMap(data map[string]map[string]interface{}, schema map[string]bool) error {
	return ft.loadCacheData(data, schema)
}

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
