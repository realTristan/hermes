package cache

import (
	"errors"

	Utils "github.com/realTristan/Hermes/utils"
)

/*
Fields:
  - wordCache: a map of words to an array of positions where the word appears in the text
  - maxWords: the maximum number of words to cache
  - maxSizeBytes: the maximum size of the text to cache in bytes
  - initialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	wordCache    map[string][]string
	maxWords     int
	maxSizeBytes int
}

// Get whether the full text cache is initialized
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

	// Check if the current size of the word cache is greater than the new max size
	if i, err := Utils.Size(c.ft.wordCache); err != nil {
		return err
	} else if i > maxBytes {
		return errors.New("the current size of the word cache is greater than the new max size")
	}

	// Set the maxSizeBytes field
	c.ft.maxSizeBytes = maxBytes

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

	// Check if the current size of the word cache is greater than the new max size
	if len(c.ft.wordCache) > maxWords {
		return errors.New("the current size of the word cache is greater than the new max size")
	}

	// Set the maxWords field
	c.ft.maxWords = maxWords

	// Return no error
	return nil
}

// Get the full text word cache
// This method is thread-safe.
func (c *Cache) FTWordCache() (map[string][]string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return nil, errors.New("full text cache not initialized")
	}

	// Copy the wordCache map
	var copy map[string][]string = make(map[string][]string, len(c.ft.wordCache))
	copy = c.ft.wordCache

	// Return the copy
	return copy, nil
}

// Get the word cache size in bytes
// This method is thread-safe.
func (c *Cache) FTWordCacheSize() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return -1, errors.New("full text cache not initialized")
	}

	// Return the size of the wordCache map
	return Utils.Size(c.ft.wordCache)
}

// Get the word cache length
// This method is thread-safe.
func (c *Cache) FTWordCacheLength() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the cache is initialized
	if c.ft == nil {
		return -1, errors.New("full text cache not initialized")
	}

	// Return the size of the wordCache map
	return len(c.ft.wordCache), nil
}
