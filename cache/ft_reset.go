package cache

import "fmt"

/*
ResetFT initializes a new FullText cache with the given maximum number of words,
maximum size in bytes, and schema. It first acquires a write lock on the cache to
prevent concurrent access. It returns an error if the FullText cache has not been
previously initialized.

Arguments:
  - maxWords: An integer representing the maximum number of words that can be stored
    in the FullText cache.
  - maxSizeBytes: An integer representing the maximum size in bytes that the FullText
    cache can have.
  - schema: A map[string]bool representing the schema of the FullText cache.

Returns:
  - An error if the FullText cache has not been previously initialized or if the
    resetFT function fails to reset the cache.

Example usage:

	cache := InitCache()
	if err := cache.InitFT(100, 1024, map[string]bool{"field1": true}); err != nil {
		log.Fatalf("Error initializing FullText cache: %v", err)
	}
	// Reset the FullText cache
	if err := cache.ResetFT(200, 2048, map[string]bool{"field1": true, "field2": true}); err != nil {
		log.Fatalf("Error resetting FullText cache: %v", err)
	}
*/
func (c *Cache) FTReset(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ftInit(maxWords, maxSizeBytes, schema)
}

/*
ResetFTWithMap resets the FullText cache with a map and Mutex Locking.

This function takes a map of map[string]interface{}, maximum number of words to index, maximum size in bytes, and a schema as input arguments.
It acquires a lock on the cache using a mutex to ensure that only one goroutine can access the cache at a time.

@Parameters:
  - data (map[string]map[string]interface{}): The map containing the data to be imported
  - maxWords (int): The maximum number of words to index
  - maxSizeBytes (int): The maximum size in bytes
  - schema (map[string]bool): The schema to be used for the imported data

@Returns:
  - error: An error if the FullText cache is not initialized or if there is an error during the reset process
*/
func (c *Cache) FTResetWithMap(data map[string]map[string]interface{}, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ftInitWithMap(data, maxWords, maxSizeBytes, schema)
}

/*
ResetFTJson resets the FullText cache with Json and Mutex Locking.

This function takes a file path, maximum number of words to index, maximum size in bytes, and a schema as input arguments.
It acquires a lock on the cache using a mutex to ensure that only one goroutine can access the cache at a time.

@Parameters:
  - file (string): The file path of the JSON file containing the data to initialize the cache with.
  - maxWords (int): The maximum number of words to index.
  - maxSizeBytes (int): The maximum size in bytes.
  - schema (map[string]bool): A map of field names to boolean values, where the boolean value indicates whether the field should be indexed.

@Returns:
  - error: If the FullText cache has not been initialized, or if an error occurs during the reset and initialization process, it will return an error. Otherwise, it will return nil.
*/
func (c *Cache) FTResetWithJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil {
		return fmt.Errorf("full text is not initialized")
	}
	return c.ftInitWithJson(file, maxWords, maxSizeBytes, schema)
}
