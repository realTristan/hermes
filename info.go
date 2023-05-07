package cache

import (
	"errors"
	"fmt"

	Utils "github.com/realTristan/Hermes/utils"
)

// Return a string with the cache, and full-text info.
// This method is thread-safe.
// An error is returned if the full-text index is not initialized.
func (c *Cache) Info() (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.info()
}

// Return a string with the cache, and full-text info.
// This method is not thread-safe, and should only be called from
// an exported function.
// An error is returned if the full-text index is not initialized.
func (c *Cache) info() (string, error) {
	// The initial cache info string
	var s string = fmt.Sprintf("Cache Info:\n-----------\nNumber of keys: %d\n", len(c.data))

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return s, errors.New("full-text is not initialized")
	}

	// Append the full-text info to the cache info string
	s += "\nFull-Text Cache Info:\n-----------\n"
	if wordCacheSize, err := Utils.Size(c.ft.storage); err != nil {
		return s, err
	} else {
		s += fmt.Sprintf("Number of keys: %d\n", len(c.ft.storage))
		s += fmt.Sprintf("Full-text cache size (bytes): %d\n", wordCacheSize)
	}
	return s, nil
}

// Return a string with the cache, and full-text info.
// This method is not thread-safe, and should only be called from
// an exported function.
// An error is returned if the full-text index is not initialized.
func (c *Cache) InfoForTesting() (string, error) {
	// The initial cache info string
	var s string = fmt.Sprintf("Cache Info:\n-----------\nNumber of keys: %d\nData: %v\n", len(c.data), c.data)

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return s, errors.New("full-text is not initialized")
	}

	// Append the full-text info to the cache info string
	s += "\nFull-Text Cache Info:\n-----------\n"
	if wordCacheSize, err := Utils.Size(c.ft.storage); err != nil {
		return s, err
	} else {
		s += fmt.Sprintf("Number of keys: %d\n", len(c.ft.storage))
		s += fmt.Sprintf("Full-text cache: %v\n", c.ft.storage)
		s += fmt.Sprintf("Full-text cache size: %d\n", wordCacheSize)
	}
	return s, nil
}
