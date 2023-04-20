package cache

import (
	"fmt"
)

/*
Initialize the FT cache with Mutex Locking.

This function initializes the FullText cache with a given maximum number of words,
maximum size in bytes, and a schema for the cache data. It also adds a mutex lock to
prevent data races when accessing the cache.

@Parameters:

	maxWords (int): The maximum number of words to store in the cache.
	maxSizeBytes (int): The maximum size of the cache in bytes.
	schema (map[string]bool): A schema for the cache data.

Returns:

	error: An error if the FullText cache is already initialized or an error occurs during initialization.
*/
func (c *Cache) InitFT(maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft != nil && c.ft.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}
	return c.initFT(maxWords, maxSizeBytes, schema)
}

/*
The `initFT` method initializes the FullText index for the cache.

	The method initializes the FullText struct and loads the cache data. It then sets the FullText index as
	initialized and returns the cache. If the FullText index is already initialized, the method returns an error.

	Parameters:
	- maxWords: the maximum number of words to store in the FullText index.
	- maxSizeBytes: the maximum size, in bytes, of the FullText index.
	- schema: a map of field names to boolean values indicating whether the field should be indexed in the FullText index.

	Returns:
	- error: an error if the FullText index is already initialized, or if there is an error loading the cache data.

	Example usage:
	cache := &Cache{}
	err := cache.initFT(1000, 1024, map[string]bool{"title": true, "description": true})
	if err != nil {
			log.Fatalf("Error initializing FullText index: %v", err)
	}
*/
func (c *Cache) initFT(maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	// Initialize the FT struct
	c.ft = &FullText{
		wordCache:     make(map[string][]int),
		keys:          []string{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
	}

	// Load the cache data
	if err := c.ft.loadCacheData(c.data, schema); err != nil {
		c.ft.clean()
		return err
	}

	// Set the keys array
	for k := range c.data {
		c.ft.keys = append(c.ft.keys, k)
	}

	// Set the FT cache as initialized
	c.ft.isInitialized = true
	return nil
}

/*
This function is used to initialize the FT cache with a JSON file with Mutex Locking. It takes the file path,
the maximum number of words, the maximum size in bytes, and a schema map as arguments. It locks the cache's mutex
to prevent concurrent access while calling the initFTJson method.

@Parameters:

	filePath (string): The path to the JSON file to initialize the FT cache with.
	maxWords (int): The maximum number of words to cache.
	maxSizeBytes (int): The maximum size in bytes of the cache.
	schema (map[string]bool): A map that specifies which keys in the JSON file to cache.

Returns:

	error: An error if the FT cache is already initialized or an error occurs during initialization.
*/
func (c *Cache) InitFTWithMap(data map[string]map[string]interface{}, maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft != nil && c.ft.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}
	return c.initFTWithMap(data, maxWords, maxSizeBytes, schema)
}

/*
This function is used to initialize the FT cache with a map with Mutex Locking. It takes the map, the maximum number of words,
the maximum size in bytes, and a schema map as arguments. It locks the cache's mutex to prevent concurrent access while calling
the initFTMap method.

@Parameters:

	data (map[string]map[string]interface{}): The map to initialize the FT cache with.
	maxWords (int): The maximum number of words to cache.
	maxSizeBytes (int): The maximum size in bytes of the cache.
	schema (map[string]bool): A map that specifies which keys in the JSON file to cache.

Returns:

	error: An error if the FT cache is already initialized or an error occurs during initialization.
*/
func (c *Cache) initFTWithMap(data map[string]map[string]interface{}, maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	// Initialize the FT struct
	c.ft = &FullText{
		wordCache:     make(map[string][]int),
		keys:          []string{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
	}

	// Iterate over the cache keys and add them to the data
	for k := range c.data {
		if _, ok := data[k]; ok {
			return fmt.Errorf("key %s already exists in cache", k)
		}
		data[k] = c.data[k]
	}

	// Load the cache data
	if err := c.ft.loadCacheData(data, schema); err != nil {
		c.ft.clean()
		return err
	}

	// Set the data
	c.data = data
	for k := range data {
		c.ft.keys = append(c.ft.keys, k)
	}

	// Set the FT cache as initialized
	c.ft.isInitialized = true
	return nil
}

/*
This function is used to initialize the FT cache with a JSON file with Mutex Locking. It takes the file path,
the maximum number of words, the maximum size in bytes, and a schema map as arguments. It locks the cache's mutex
to prevent concurrent access while calling the initFTJson method.

@Parameters:

	file (string): The path to the JSON file.
	maxWords (int): The maximum number of words to cache.
	maxSizeBytes (int): The maximum size in bytes of the cache.
	schema (map[string]bool): A map that specifies which keys in the JSON file to cache.

@Returns:

	error: An error indicating if there was a problem initializing the FT cache with the JSON file.

Example usage:

	err := c.InitFTJson("data.json", 10000, 1048576, map[string]bool{"title": true, "description": true})
	// Initializes the FT cache with the data in "data.json" with a maximum of 10000 words and a maximum size of 1MB,
	// caching only the "title" and "description" keys from the JSON file, and using Mutex Locking to prevent concurrent access.

Note that this function is a wrapper around the initFTJson method that adds mutex locking to prevent concurrent access.
*/
func (c *Cache) InitFTWithJson(file string, maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft != nil && c.ft.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}
	return c.initFTWithJson(file, maxWords, maxSizeBytes, schema)
}

/*
This function is used to initialize the FT cache with a JSON file. It takes the file path, the maximum number of words,
the maximum size in bytes, and a schema map as arguments. It locks the cache's mutex to prevent concurrent access while calling
the initFTJson method.

@Parameters:

	file (string): The path to the JSON file.
	maxWords (int): The maximum number of words to cache.
	maxSizeBytes (int): The maximum size in bytes of the cache.
	schema (map[string]bool): A map that specifies which keys in the JSON file to cache.

@Returns:

	error: An error indicating if there was a problem initializing the FT cache with the JSON file.
*/
func (c *Cache) initFTWithJson(file string, maxWords int, maxSizeBytes uintptr, schema map[string]bool) error {
	if data, err := readJson(file); err != nil {
		return err
	} else {
		return c.initFTWithMap(data, maxWords, maxSizeBytes, schema)
	}
}
