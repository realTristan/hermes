package cache

import "fmt"

/*
Import map data to the FullText cache with Mutex Locking.

This function accepts a map of map[string]interface{} data and a schema, and imports the data to the FullText cache.

@Parameters:
  - data: a map[string]map[string]interface{} containing the data to be imported
  - schema: a map[string]bool representing the schema to be used for the imported data

@Returns:
  - error: an error object if an error occurs during the import process, or nil if the import is successful

Example Usage:

	data := map[string]map[string]interface{}{
			"1": map[string]interface{}{"text": "hello world"},
			"2": map[string]interface{}{"text": "foo bar"},
	}
	schema := map[string]bool{"text": true}
	err := ft.Import(data, schema)
*/
func (c *Cache) Import(data map[string]map[string]interface{}, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ft.loadCacheData(data, schema)
}

/*
ImportJson imports JSON data to the FullText cache with Mutex Locking.

This function accepts a JSON file path and a schema, and imports the data to the FullText cache.

@Parameters:
  - file: a string representing the path to the JSON file to be imported
  - schema: a map[string]bool representing the schema to be used for the imported data

@Returns:
  - error: an error object if an error occurs during the import process, or nil if the import is successful

Example Usage:

	schema := map[string]bool{"text": true}
	err := ft.ImportJson("data.json", schema)
*/
func (c *Cache) ImportJson(file string, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return fmt.Errorf("full text is not initialized")
	}

	// Read the JSON file and import the data
	if data, err := readJson(file); err != nil {
		return err
	} else {
		return c.ft.loadCacheData(data, schema)
	}
}
