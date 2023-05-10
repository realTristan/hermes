package cache

import (
	"errors"

	Utils "github.com/realTristan/Hermes/utils"
)

/*
Fields:
  - storage: a map of words to an array of indices where the word appears in the text
  - indices: a map of indices to words
  - currentIndex: the current position in the map of indices to words
  - maxLength: the maximum number of words in the storage
  - maxBytes: the maximum size of the text to storage in bytes
  - initialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	storage      map[string][]int
	indices      map[int]string
	currentIndex int
	maxLength    int
	maxBytes     int
}

// Get whether the full-text is initialized
// This method is thread-safe.
func (c *Cache) FTIsInitialized() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.ft != nil
}

// Set the max size in bytes of the full text storage
// This method is thread-safe.
func (c *Cache) FTSetMaxBytes(maxBytes int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return errors.New("full text not initialized")
	}

	// Check if the current size of the storage is greater than the new max size
	if i, err := Utils.Size(c.ft.storage); err != nil {
		return err
	} else if i > maxBytes {
		return errors.New("the current size of the full-text storage is greater than the new max size")
	}

	// Set the maxBytes field
	c.ft.maxBytes = maxBytes

	// Return no error
	return nil
}

// Set the maximum number of words in the full text storage
// This method is thread-safe.
func (c *Cache) FTSetMaxLength(maxLength int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return errors.New("full text not initialized")
	}

	// Check if the current size of the storage is greater than the new max size
	if len(c.ft.storage) > maxLength {
		return errors.New("the current size of the full-text storage is greater than the new max size")
	}

	// Set the maxLength field
	c.ft.maxLength = maxLength

	// Return no error
	return nil
}

// Get the full text storage
// This method is thread-safe.
func (c *Cache) FTStorage() (map[string][]int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return nil, errors.New("full text is not initialized")
	}

	// Copy the storage map
	var copy map[string][]int = c.ft.storage

	// Return the copy
	return copy, nil
}

// Get the full-text storage size in bytes
// This method is thread-safe.
func (c *Cache) FTStorageSize() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return -1, errors.New("full text not initialized")
	}

	// Return the size of the storage map
	return Utils.Size(c.ft.storage)
}

// Get the full-text storage length
// This method is thread-safe.
func (c *Cache) FTStorageLength() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return -1, errors.New("full text not initialized")
	}

	// Return the size of the storage map
	return len(c.ft.storage), nil
}
