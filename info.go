package hermes

import (
	"errors"

	utils "github.com/realTristan/hermes/utils"
)

// Info is a method of the Cache struct that returns a map with the cache and full-text info.
// This method is thread-safe.
// An error is returned if the full-text index is not initialized.
//
// Returns:
//   - A map[string]any representing the cache and full-text info.
//   - An error if the full-text index is not initialized.
func (c *Cache) Info() (map[string]any, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.info()
}

// info is a method of the Cache struct that returns a map with the cache and full-text info.
// This method is not thread-safe, and should only be called from an exported function.
// An error is returned if the full-text index is not initialized.
//
// Returns:
//   - A map[string]any representing the cache and full-text info.
//   - An error if the full-text index is not initialized.
func (c *Cache) info() (map[string]any, error) {
	var info map[string]any = map[string]any{
		"keys": len(c.data),
	}

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return info, errors.New("full-text is not initialized")
	}

	// Add the full-text info to the map
	if size, err := utils.Size(c.ft.storage); err != nil {
		return info, err
	} else {
		// Add the full-text info to the map
		info["full-text"] = map[string]any{
			"keys":  len(c.ft.storage),
			"index": c.ft.index,
			"size":  size,
		}
	}

	// Return the info map
	return info, nil
}

// InfoForTesting is a method of the Cache struct that returns a map with the cache and full-text info for testing purposes.
// This method is thread-safe.
// An error is returned if the full-text index is not initialized.
//
// Returns:
//   - A map[string]any representing the cache and full-text info for testing purposes.
//   - An error if the full-text index is not initialized.
func (c *Cache) InfoForTesting() (map[string]any, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.infoForTesting()
}

// infoForTesting is a method of the Cache struct that returns a map with the cache and full-text info for testing purposes.
// This method is not thread-safe, and should only be called from an exported function.
// An error is returned if the full-text index is not initialized.
//
// Returns:
//   - A map[string]any representing the cache and full-text info for testing purposes.
//   - An error if the full-text index is not initialized.
func (c *Cache) infoForTesting() (map[string]any, error) {
	var info map[string]any = map[string]any{
		"keys": len(c.data),
		"data": c.data,
	}

	// Check if the cache full-text has been initialized
	if c.ft == nil {
		return info, errors.New("full-text is not initialized")
	}

	// Add the full-text info to the map
	if size, err := utils.Size(c.ft.storage); err != nil {
		return info, err
	} else {
		info["full-text"] = map[string]any{
			"keys":    len(c.ft.storage),
			"index":   c.ft.index,
			"size":    size,
			"storage": c.ft.storage,
			"indices": c.ft.indices,
		}
	}

	// Return the info map
	return info, nil
}
