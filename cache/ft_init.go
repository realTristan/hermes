package cache

import (
	"fmt"
	"sync"
)

/*
This function is used to initialize the FT cache with a JSON file with Mutex Locking. It takes the file path,
the maximum number of words, the maximum size in bytes, and a schema map as arguments. It locks the cache's mutex
to prevent concurrent access while calling the initFTJson method. If the cache's FT field is not nil,
it locks the FT's mutex as well to prevent concurrent access while initializing the FT cache. Once the
initFTJson method returns, it unlocks both mutexes.

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
func (c *Cache) InitFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}

/*
This function is used to initialize the FT cache with a JSON file. It takes the file path,
the maximum number of words, the maximum size in bytes, and a schema map as arguments. If the FT cache is already initialized,
it returns an error indicating that the FT cache has already been initialized. Otherwise, it initializes the FT cache by creating a
new instance of the FullText struct and setting its fields accordingly. It then reads the JSON file and assigns its contents to the
data field of the FT cache. It checks if there are any keys in the JSON file that are already in the regular cache and returns an
error if any are found. It then loads the cache data using the loadCacheData method of the FullText struct and sets the isInitialized field to true.

@Parameters:

	file (string): The path to the JSON file.
	maxWords (int): The maximum number of words to cache.
	maxSizeBytes (int): The maximum size in bytes of the cache.
	schema (map[string]bool): A map that specifies which keys in the JSON file to cache.

@Returns:

	error: An error indicating if there was a problem initializing the FT cache with the JSON file.

Example usage:

	err := c.initFTJson("data.json", 10000, 1048576, map[string]bool{"title": true, "description": true})
	// Initializes the FT cache with the data in "data.json" with a maximum of 10000 words and a maximum size of 1MB,
	// caching only the "title" and "description" keys from the JSON file.

Note that this function does not use mutex locking to prevent concurrent access, as it is intended to
be used within the context of the Cache struct, which already has its own mutex locking.
*/
func (c *Cache) initFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is already initialized, return an error
	if c.FT != nil && c.FT.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}

	// Initialize the FT cache
	c.FT = &FullText{
		mutex:         &sync.RWMutex{},
		wordCache:     map[string][]string{},
		data:          map[string]map[string]interface{}{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
	}

	// Read the json
	if data, err := readJson(file); err != nil {
		return err
	} else {
		c.FT.data = data
	}

	// Iterate over the regular cache data and
	// add it to the FT cache data
	for k, v := range c.data {
		if _, ok := c.FT.data[k]; ok {
			c.FT.clean()
			return fmt.Errorf("the key: %s has already been imported via json file. unable to add it to the cache", k)
		}
		c.FT.data[k] = v
	}

	// Load the cache data
	if err := c.FT.loadCacheData(c.FT.data, schema); err != nil {
		c.FT.clean()
		return err
	}

	// Set the FT cache as initialized
	c.FT.isInitialized = true

	// Return the cache
	return nil
}

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
func (c *Cache) InitFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
	return c.initFT(maxWords, maxSizeBytes, schema)
}

/*
InitFT initializes the FT cache with the given parameters. It converts the cache
data into an array of maps and loads it into the FT cache. If the FT cache is already initialized, it returns an error.

@Parameters:

	maxWords: the maximum number of words to cache.
	maxSizeBytes: the maximum size of the cache data in bytes.
	schema: a map[string]bool where the keys represent the fields that should be indexed, and the values represent whether the field should be treated as a text field or a numerical field.

@Returns:

	An error if the FT cache is already initialized, or if there was an error loading the cache data into the FT cache. Otherwise, returns nil.
*/
func (c *Cache) initFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is already initialized, return an error
	if c.FT != nil && c.FT.isInitialized {
		return fmt.Errorf("full text cache already initialized")
	}

	// Initialize the FT struct
	c.FT = &FullText{
		mutex:         &sync.RWMutex{},
		wordCache:     map[string][]string{},
		data:          map[string]map[string]interface{}{},
		maxWords:      maxWords,
		maxSizeBytes:  maxSizeBytes,
		isInitialized: false,
	}

	// Convert the cache data into an array of maps
	for k, v := range c.data {
		c.FT.data[k] = v
	}

	// Load the cache data
	if err := c.FT.loadCacheData(c.FT.data, schema); err != nil {
		c.FT.clean()
		return err
	}

	// Set the FT cache as initialized
	c.FT.isInitialized = true

	// Return the cache
	return nil
}
