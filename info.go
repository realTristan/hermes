package cache

import (
	"errors"
	"fmt"

	Utils "github.com/realTristan/Hermes/utils"
)

// Return a string with the cache, and full-text info.
// This method is thread-safe.
// An error is returned if the full-text index is not initialized.
func (c *Cache) InfoString() (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.infoString()
}

// Return a string with the cache, and full-text info.
// This method is not thread-safe, and should only be called from
// an exported function.
// An error is returned if the full-text index is not initialized.
func (c *Cache) infoString() (string, error) {
	// The initial cache info string
	var s string = fmt.Sprintf("Cache Info:\n-----------\nNumber of keys: %d\n", len(c.data))

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return s, errors.New("full-text is not initialized")
	}

	// Append the full-text info to the cache info string
	s += "\nFull-Text Cache Info:\n-----------\n"
	if size, err := Utils.Size(c.ft.storage); err != nil {
		return s, err
	} else {
		s += fmt.Sprintf("Current Index: %d\n", c.ft.currentIndex)
		s += fmt.Sprintf("Full-text storage size (bytes): %d\n", size)
		s += fmt.Sprintf("Number of keys: %d\n", len(c.ft.storage))
	}
	return s, nil
}

// Return a string with the cache, and full-text info.
// This method is thread-safe.
// An error is returned if the full-text index is not initialized.
func (c *Cache) InfoStringForTesting() (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.infoStringForTesting()
}

// Return a string with the cache, and full-text info.
// This method is not thread-safe, and should only be called from
// an exported function.
// An error is returned if the full-text index is not initialized.
func (c *Cache) infoStringForTesting() (string, error) {
	// The initial cache info string
	var s string = fmt.Sprintf("Cache Info:\n-----------\nNumber of keys: %d\nData: %v\n", len(c.data), c.data)

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return s, errors.New("full-text is not initialized")
	}

	// Append the full-text info to the cache info string
	s += "\nFull-Text Cache Info:\n-----------\n"
	if size, err := Utils.Size(c.ft.storage); err != nil {
		return s, err
	} else {
		s += fmt.Sprintf("Number of keys: %d\n", len(c.ft.storage))
		s += fmt.Sprintf("Full-text storage: %v\n", c.ft.storage)
		s += fmt.Sprintf("Full-text storage size: %d\n", size)
		s += fmt.Sprintf("Indices: %v\n", c.ft.indices)
		s += fmt.Sprintf("Current Index: %d\n", c.ft.currentIndex)
	}
	return s, nil
}
