package cache

import "fmt"

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
func (c *Cache) UploadJson(file string, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil || !c.ft.isInitialized {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ft.uploadJson(file, schema)
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
func (c *Cache) UploadMap(data map[string]map[string]interface{}, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil || !c.ft.isInitialized {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ft.uploadMap(data, schema)
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
