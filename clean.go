package hermes

import "errors"

// Clean is a method of the Cache struct that clears the cache contents.
// If the full-text index is initialized, it is also cleared.
// This method is thread-safe.
//
// Parameters:
//   - None
//
// Returns:
//   - None
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clean()
}

// clean is a method of the Cache struct that clears the cache contents.
// If the full-text index is initialized, it is also cleared.
// This method is not thread-safe and should only be called from an exported function.
//
// Parameters:
//   - None
//
// Returns:
//   - None
func (c *Cache) clean() {
	if c.ft != nil {
		c.ft.clean()
	}
	c.data = map[string]map[string]any{}
}

// FTClean is a method of the Cache struct that clears the full-text cache contents.
// If the full-text index is not initialized, this method returns an error.
// Otherwise, the full-text index storage and indices are cleared, and this method returns nil.
// This method is thread-safe.
//
// Returns:
//   - error: An error object. If no error occurs, this will be nil.
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

// clean is a method of the FullText struct that clears the full-text index storage and indices.
// This method initializes a new empty storage map and indices map with the maximum length specified in the FullText struct.
//
// Parameters:
//   - None
//
// Returns:
//   - None
func (ft *FullText) clean() {
	ft.storage = make(map[string]any)
	ft.indices = make(map[int]string)
}
