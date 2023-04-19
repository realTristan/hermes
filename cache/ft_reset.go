package cache

import "fmt"

// Reset the FullText cache with Mutex Locking
func (c *Cache) ResetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFT(maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache
func (c *Cache) resetFT(maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.FT == nil {
		return fmt.Errorf("full text cache not initialized")
	}

	// Lock the FT cache
	c.FT.mutex.Lock()
	defer c.FT.mutex.Unlock()

	// If the FT cache is not initialized, return an error
	if !c.FT.isInitialized {
		return fmt.Errorf("full text cache not initialized")
	}

	// Reset the FT cache
	c.FT.isInitialized = false
	return c.initFT(maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache with Json and Mutex Locking
func (c *Cache) ResetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.resetFTJson(file, maxWords, maxSizeBytes, schema)
}

// Reset the FullText cache with Json
func (c *Cache) resetFTJson(file string, maxWords int, maxSizeBytes int, schema map[string]bool) error {
	// If the FT cache is not initialized, return an error
	if c.FT == nil {
		return fmt.Errorf("FT cache not initialized")
	}

	// Lock the FT cache
	c.FT.mutex.Lock()
	defer c.FT.mutex.Unlock()

	// If the FT cache is not initialized, return an error
	if !c.FT.isInitialized {
		return fmt.Errorf("FT cache not initialized")
	}

	// Reset the FT cache
	c.FT.isInitialized = false
	return c.initFTJson(file, maxWords, maxSizeBytes, schema)
}
