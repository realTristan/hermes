package cache

import (
	"errors"

	Utils "github.com/realTristan/Hermes/utils"
)

// FullText is a struct that represents a full-text index for a cache of data.
// The index is used to enable full-text search on the data in the cache.
//
// Fields:
//   - storage (map[string]any): A map that stores the indices of the entries in the cache that contain each word in the full-text index. The keys of the map are strings that represent the words in the index, and the values are slices of integers that represent the indices of the entries in the cache that contain the word.
//   - indices (map[int]string): A map that stores the words in the full-text index. The keys of the map are integers that represent the indices of the words in the index, and the values are strings that represent the words.
//   - currentIndex (int): An integer that represents the current index of the full-text index. This is used to assign unique indices to new words as they are added to the index.
//   - maxLength (int): An integer that represents the maximum number of words that can be stored in the full-text index.
//   - maxBytes (int): An integer that represents the maximum size of the text that can be stored in the full-text index, in bytes.
type FullText struct {
	storage      map[string]any // either []int or int
	indices      map[int]string
	currentIndex int
	maxLength    int
	maxBytes     int
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

// FTSetMaxLength is a method of the Cache struct that sets the maximum number of words in the full-text index.
// If the full-text index is not initialized, this method returns an error.
// If the current size of the full-text index is greater than the new maximum size, this method returns an error.
// Otherwise, the maximum number of words in the full-text index is set to the specified value, and this method returns nil.
// This method is thread-safe.
//
// Parameters:
//   - maxLength (int): An integer that represents the new maximum number of words in the full-text index.
//
// Returns:
//   - error: An error object. If no error occurs, this will be nil.
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
	return Utils.Size(c.ft.storage)
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
