package cache

import "errors"

// Clear the cache contents.
// This function is thread-safe.
// If the full-text index is initialized, clear it as well.
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clean()
}

// Clear the cache contents.
// This function is not thread-safe, and should only be called from
// an exported function.
// If the full-text index is initialized, clear it as well.
func (c *Cache) clean() {
	if c.ft != nil {
		c.ft.clean()
	}
	c.data = map[string]map[string]interface{}{}
}

// Clear the full-text cache contents.
// This function is thread-safe.
func (c *Cache) FTClean() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Verify that the full text is initialized
	if c.ft == nil {
		return errors.New("full text is not initialized")
	}

	// Clean the ft cache
	c.ft.clean()

	// Return no error
	return nil
}

// Clear the full-text cache contents.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) clean() {
	ft.wordCache = make(map[string][]int, ft.maxWords)
	ft.indicesCache = make(map[int]string)
}
