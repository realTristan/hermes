package hermes

import (
	"errors"

	utils "github.com/realTristan/hermes/utils"
)

// FullText is a struct that represents a full-text index for a cache of data.
// The index is used to enable full-text search on the data in the cache.
//
// Fields:
//   - storage (map[string]any): A map that stores the indices of the entries in the cache that contain each word in the full-text index. The keys of the map are strings that represent the words in the index, and the values are slices of integers that represent the indices of the entries in the cache that contain the word.
//   - indices (map[int]string): A map that stores the words in the full-text index. The keys of the map are integers that represent the indices of the words in the index, and the values are strings that represent the words.
//   - index (int): An integer that represents the current index of the full-text index. This is used to assign unique indices to new words as they are added to the index.
//   - maxSize (int): An integer that represents the maximum number of words that can be stored in the full-text index.
//   - maxBytes (int): An integer that represents the maximum size of the text that can be stored in the full-text index, in bytes.
//   - minWordLength (int): An integer that represents the minimum length of a word that can be stored in the full-text index.
type FullText struct {
	storage       map[string]any // either []int or int
	indices       map[int]string
	index         int
	maxSize       int
	maxBytes      int
	minWordLength int
}

// FTIsInitialized is a method of the Cache struct that returns a boolean value indicating whether the full-text index is initialized.
// If the full-text index is initialized, this method returns true. Otherwise, it returns false.
// This method is thread-safe.
//
// Parameters:
//   - None
//
// Returns:
//   - bool: A boolean value indicating whether the full-text index is initialized. If the full-text index is initialized, this method returns true. Otherwise, it returns false.
func (c *Cache) FTIsInitialized() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.ft != nil
}

// FTSetMaxBytes is a method of the Cache struct that sets the maximum size of the full-text index in bytes.
// If the full-text index is not initialized, this method returns an error.
// If the current size of the full-text index is greater than the new maximum size, this method returns an error.
// Otherwise, the maximum size of the full-text index is set to the specified value, and this method returns nil.
// This method is thread-safe.
//
// Parameters:
//   - maxBytes (int): An integer that represents the new maximum size of the full-text index, in bytes.
//
// Returns:
//   - error: An error object. If no error occurs, this will be nil.
func (c *Cache) FTSetMaxBytes(maxBytes int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return errors.New("full text not initialized")
	}

	// Check if the current size of the storage is the same as the new max size
	if c.ft.maxBytes == maxBytes {
		return nil
	}

	// Check if the current size of the storage is greater than the new max size
	if i, err := utils.Size(c.ft.storage); err != nil {
		return err
	} else if i > maxBytes {
		return errors.New("the current size of the full-text storage is greater than the new max size")
	}

	// Set the maxBytes field
	c.ft.maxBytes = maxBytes

	// Return no error
	return nil
}

// FTSetMaxSize is a method of the Cache struct that sets the maximum number of words in the full-text index.
// If the full-text index is not initialized, this method returns an error.
// If the current size of the full-text index is greater than the new maximum size, this method returns an error.
// Otherwise, the maximum number of words in the full-text index is set to the specified value, and this method returns nil.
// This method is thread-safe.
//
// Parameters:
//   - maxSize (int): An integer that represents the new maximum number of words in the full-text index.
//
// Returns:
//   - error: An error object. If no error occurs, this will be nil.
func (c *Cache) FTSetMaxSize(maxSize int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return errors.New("full text not initialized")
	}

	// Check if the current size of the storage is the same as the new max size
	if maxSize == c.ft.maxSize {
		return nil
	}

	// Check if the current size of the storage is greater than the new max size
	if len(c.ft.storage) > maxSize {
		return errors.New("the current size of the full-text storage is greater than the new max size")
	}

	// Set the maxSize field
	c.ft.maxSize = maxSize

	// Return no error
	return nil
}

// FTSetMinWordLength is a method of the Cache struct that sets the minimum word length for the full-text search.
// Parameters:
//   - minWordLength (int): An integer representing the minimum word length.
//
// Returns:
//   - error: An error if the full-text search is not initialized or if the new minimum word length is greater than the maximum word length.
func (c *Cache) FTSetMinWordLength(minWordLength int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return errors.New("full text not initialized")
	}

	// If they're the same
	if minWordLength == c.ft.minWordLength {
		return nil
	}

	// If the new min word length is greater than the max
	// word length, reset the ft
	if minWordLength > c.ft.minWordLength {
		return c.ftInit(c.ft.maxBytes, c.ft.maxSize, minWordLength)
	}

	// Set the minWordLength field
	c.ft.minWordLength = minWordLength

	// Iterate over the ft storage
	for word := range c.ft.storage {
		// Check if the word length is less than the min word length
		if len(word) < minWordLength {
			// Delete the word from the ft storage
			delete(c.ft.storage, word)
		}
	}

	// Return no error
	return nil
}

// FTStorage is a method of the Cache struct that returns a copy of the full-text index storage map.
// If the full-text index is not initialized, this method returns an error.
// Otherwise, a copy of the full-text index storage map is returned, and this method returns nil.
// This method is thread-safe.
//
// Returns:
//   - map[string][]int: A copy of the full-text index storage map.
//   - error: An error object. If no error occurs, this will be nil.
func (c *Cache) FTStorage() (map[string]any, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return nil, errors.New("full text is not initialized")
	}

	// Copy the storage map
	var copy map[string]any = c.ft.storage

	// Return the copy
	return copy, nil
}

// FTStorageSize is a method of the Cache struct that returns the size of the full-text index storage in bytes.
// If the full-text index is not initialized, this method returns an error.
// Otherwise, the size of the full-text index storage is returned as an integer, and this method returns nil.
// This method is thread-safe.
//
// Returns:
//   - int: An integer that represents the size of the full-text index storage in bytes.
//   - error: An error object. If no error occurs, this will be nil.
func (c *Cache) FTStorageSize() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the ft is initialized
	if c.ft == nil {
		return -1, errors.New("full text not initialized")
	}

	// Return the size of the storage map
	return utils.Size(c.ft.storage)
}

// FTStorageLength is a method of the Cache struct that returns the number of words in the full-text index storage.
// If the full-text index is not initialized, this method returns an error.
// Otherwise, the number of words in the full-text index storage is returned as an integer, and this method returns nil.
// This method is thread-safe.
//
// Returns:
//   - int: An integer that represents the number of words in the full-text index storage.
//   - error: An error object. If no error occurs, this will be nil.
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
