package cache

import "fmt"

/*
ResetFT initializes a new FullText cache with the given maximum number of words,
maximum size in bytes, and schema. It first acquires a write lock on the cache to
prevent concurrent access, and then calls the resetFT function to reset the cache.
It returns an error if the FullText cache has not been previously initialized or if
the resetFT function fails to reset the cache.

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
func (c *Cache) ResetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFT(maxWords, maxSizeBytes, schema)
}

/*
resetFT resets the FullText cache. It initializes a new instance of the FullText struct with the specified maxWords,
maxSizeBytes, and schema parameters and loads the cache data into the new instance.
If the FullText cache is not already initialized, an error will be returned.

@Arguments:
  - maxWords (int): the maximum number of words to cache
  - maxSizeBytes (int): the maximum size of the cache in bytes
  - schema (map[string]bool): a map containing the cache schema, where the key is the field name and the value is a boolean indicating whether the field should be indexed for searching

@Returns:
  - error: an error if the FullText cache is not initialized or there is an issue initializing the cache with the specified parameters
*/
func (c *Cache) resetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.ft == nil {
		return fmt.Errorf("full text is not initialized")
	}

	// Reset the FT cache
	c.ft.isInitialized = false
	return c.initFT(maxWords, maxSizeBytes, schema)
}

/*
ResetFTJson resets the FullText cache with Json and Mutex Locking.

This function takes a file path, maximum number of words to index, maximum size in bytes, and a schema as input arguments.
It acquires a lock on the cache using a mutex to ensure that only one goroutine can access the cache at a time.
Then it calls the `resetFTJson` method to reset the cache and initialize it with new data from the specified JSON file.
If an error occurs during the process, it will return an error.

@Parameters:
  - file (string): The file path of the JSON file containing the data to initialize the cache with.
  - maxWords (int): The maximum number of words to index.
  - maxSizeBytes (int): The maximum size in bytes.
  - schema (map[string]bool): A map of field names to boolean values, where the boolean value indicates whether the field should be indexed.

@Returns:
  - error: If the FullText cache has not been initialized, or if an error occurs during the reset and initialization process, it will return an error. Otherwise, it will return nil.
*/
func (c *Cache) ResetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFTJson(file, maxWords, maxSizeBytes, schema)
}

/*
resetFTJson resets the FullText cache using data from a JSON file.
If the FT cache is not initialized, an error is returned.
The FT cache is locked during the reset process using a mutex.

Args:
  - file (string): The path to the JSON file containing the cache data.
  - maxWords (int): The maximum number of words to cache per document.
  - maxSizeBytes (int): The maximum size of the cache in bytes.
  - schema (map[string]bool): A map of the cache schema.

Returns:
  - error: An error if the FT cache is not initialized, or an error occurs during initialization.
*/
func (c *Cache) resetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.ft == nil {
		return fmt.Errorf("full text is not initialized")
	}

	// Reset the FT cache
	c.ft.isInitialized = false
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}
