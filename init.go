package cache

import (
	"errors"
	"fmt"
	"sync"

	Utils "github.com/realTristan/Hermes/utils"
)

// InitCache is a function that initializes a new Cache struct and returns a pointer to it.
//
// Returns:
//   - A pointer to a new Cache struct.
func InitCache() *Cache {
	return &Cache{
		data:  make(map[string]map[string]any),
		mutex: &sync.RWMutex{},
		ft:    nil,
	}
}

// Initialize the full-text for the cache
// This method is thread-safe.
// If the full-text index is already initialized, an error is returned.
//
// Parameters:
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: If the full-text is already initialized.
func (c *Cache) FTInit(maxLength int, maxBytes int, minWordLength int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Verify that the ft has already been initialized
	if c.ft != nil {
		return errors.New("full-text already initialized")
	}

	// Initialize the FT
	return c.ftInit(maxLength, maxBytes, minWordLength)
}

// Initialize the full-text for the cache.
// This method is not thread-safe, and should only be called from an exported function.
//
// Parameters:
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: From full-text cache insertion.
func (c *Cache) ftInit(maxLength int, maxBytes int, minWordLength int) error {
	// Initialize the FT struct
	var ft *FullText = &FullText{
		storage:       make(map[string]any),
		indices:       make(map[int]string),
		currentIndex:  0,
		maxLength:     maxLength,
		maxBytes:      maxBytes,
		minWordLength: minWordLength,
	}

	// Load the cache data
	if err := ft.insert(c.data); err != nil {
		return err
	}

	// Update the cache full-text
	c.ft = ft

	// Return no error
	return nil
}

// Initialize the full-text index for the cache with a map.
// This method is thread-safe.
// If the full-text index is already initialized, an error is returned.
//
// Parameters:
// - data: the data to initialize the full-text index with.
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: If the full-text is already initialized.
func (c *Cache) FTInitWithMap(data map[string]map[string]any, maxLength int, maxBytes int, minWordLength int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Verify that the cache is already initialized
	if c.ft != nil {
		return errors.New("full-text cache already initialized")
	}

	// Initialize the FT cache
	return c.ftInitWithMap(data, maxLength, maxBytes, minWordLength)
}

// Initialize the full-text for the cache with a map.
// This method is not thread-safe, and should only be called from an exported function.
//
// Parameters:
// - data: the data to initialize the full-text index with.
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: From full-text storage insertion or if a key from the data already exists in the cache
func (c *Cache) ftInitWithMap(data map[string]map[string]any, maxLength int, maxBytes int, minWordLength int) error {
	// Initialize the FT struct
	var ft *FullText = &FullText{
		storage:       make(map[string]any),
		indices:       make(map[int]string),
		currentIndex:  0,
		maxLength:     maxLength,
		maxBytes:      maxBytes,
		minWordLength: minWordLength,
	}

	// Iterate over the cache keys and add them to the data
	for k := range c.data {
		if _, ok := data[k]; ok {
			return fmt.Errorf("key %s already exists in cache", k)
		}
		data[k] = c.data[k]
	}

	// Insert the data into the ft storage
	if err := ft.insert(data); err != nil {
		return err
	}

	// Update the cache varoables
	c.data = data
	c.ft = ft

	// Return no error
	return nil
}

// Initialize the full-text for the cache with a JSON file.
// This method is thread-safe.
// If the full-text index is already initialized, an error is returned.
//
// Parameters:
// - file: the path to the JSON file to initialize the full-text index with.
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: If the full-text is already initialized.
func (c *Cache) FTInitWithJson(file string, maxLength int, maxBytes int, minWordLength int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Verify that the ft cache is initialized
	if c.ft != nil {
		return errors.New("full-text cache already initialized")
	}

	// Initialize the FT
	return c.ftInitWithJson(file, maxLength, maxBytes, minWordLength)
}

// Initialize the full-text for the cache with a JSON file.
// This method is not thread-safe, and should only be called from an exported function.
//
// Parameters:
// - file: the path to the JSON file to initialize the full-text index with.
// - maxLength: the maximum number of words to store in the full-text index.
// - maxBytes: the maximum size, in bytes, of the full-text index.
//
// Returns:
// - error: Json file read error, or init with map error.
func (c *Cache) ftInitWithJson(file string, maxLength int, maxBytes int, minWordLength int) error {
	if data, err := Utils.ReadJson[map[string]map[string]any](file); err != nil {
		return err
	} else {
		return c.ftInitWithMap(data, maxLength, maxBytes, minWordLength)
	}
}
