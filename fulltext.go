package cache

import (
	"errors"

	Utils "github.com/realTristan/Hermes/utils"
)

/*
Fields:
  - cache: a map of words to an array of indices where the word appears in the text
  - indices: a map of indices to words
  - currentIndex: the current position in the map of indices to words
  - maxWords: the maximum number of words to cache
  - maxBytes: the maximum size of the text to cache in bytes
  - initialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	cache        map[string][]int
	indices      map[int]string
	currentIndex int
	maxWords     int
	maxBytes     int
}

// Get whether the full-text cache is initialized
// This method is thread-safe.
func (c *Cache) FTIsInitialized() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.ft != nil
}

// Set the max size in bytes of the full text cache
// This method is thread-safe.
func (c *Cache) FTSetMaxBytes(maxBytes int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return errors.New("full text cache not initialized")
	}

	// Check if the current size of the cache is greater than the new max size
	if i, err := Utils.Size(c.ft.cache); err != nil {
		return err
	} else if i > maxBytes {
		return errors.New("the current size of the full-text cache is greater than the new max size")
	}

	// Set the maxBytes field
	c.ft.maxBytes = maxBytes

	// Return no error
	return nil
}

// Set the max number of words in the full text cache
// This method is thread-safe.
func (c *Cache) FTSetMaxWords(maxWords int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return errors.New("full text cache not initialized")
	}

	// Check if the current size of the cache is greater than the new max size
	if len(c.ft.cache) > maxWords {
		return errors.New("the current size of the full-text cache is greater than the new max size")
	}

	// Set the maxWords field
	c.ft.maxWords = maxWords

	// Return no error
	return nil
}

// Get the full text cache
// This method is thread-safe.
func (c *Cache) FTCache() (map[string][]int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return nil, errors.New("full text cache not initialized")
	}

	// Copy the cache map
	var copy map[string][]int = make(map[string][]int, len(c.ft.cache))
	copy = c.ft.cache

	// Return the copy
	return copy, nil
}

// Get the full-text cache size in bytes
// This method is thread-safe.
func (c *Cache) FTCacheSize() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return -1, errors.New("full text cache not initialized")
	}

	// Return the size of the cache map
	return Utils.Size(c.ft.cache)
}

// Get the full-text cache length
// This method is thread-safe.
func (c *Cache) FTCacheLength() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return -1, errors.New("full text cache not initialized")
	}

	// Return the size of the cache map
	return len(c.ft.cache), nil
}
