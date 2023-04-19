package cache

import (
	"fmt"
	"sync"
)

// Initialize the FT cache with a json file with Mutex Locking
func (c *Cache) InitFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}

// Initialize the FT cache with a json file
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

// Initialize the FT cache with Mutex Locking
func (c *Cache) InitFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.FT != nil {
		c.FT.mutex.Lock()
		defer c.FT.mutex.Unlock()
	}
	return c.initFT(maxWords, maxSizeBytes, schema)
}

// Initialize the FT cache
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
