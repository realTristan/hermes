package cache

import "fmt"

// Set is a method of the Cache struct that sets a value in the cache for the specified key.
// This function is thread-safe.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]any representing the value to set.
//
// Returns:
//   - Error
func (c *Cache) Set(key string, value map[string]any) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

// set is a method of the Cache struct that sets a value in the cache for the specified key.
// This function is not thread-safe, and should only be called from an exported function.
// If fullText is true, set the value in the full-text cache as well.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]any representing the value to set.
//
// Returns:
//   - An error if the full-text cache key already exists. Otherwise, nil.
func (c *Cache) set(key string, value map[string]any) error {
	if _, ok := c.data[key]; ok {
		return fmt.Errorf("full-text cache key already exists (%s). delete it before setting it another value", key)
	}

	// Update the value in the FT cache
	if c.ft != nil {
		if err := c.ftSet(key, value); err != nil {
			return err
		}
	}

	// Update the value in the cache
	c.data[key] = value

	// Return nil for no error
	return nil
}

// ftSet is a method of the Cache struct that sets a value in the full-text cache for the specified key.
// This function is not thread-safe, and should only be called from an exported function.
//
// Parameters:
//   - key: A string representing the key to set the value for.
//   - value: A map[string]any representing the value to set.
//
// Returns:
//   - An error if the full-text storage limit or byte-size limit is reached. Otherwise, nil.
func (c *Cache) ftSet(key string, value map[string]any) error {
	var ts *TempStorage = NewTempStorage(c.ft)
	for k, v := range value {
		if ftv := WFTGetValue(v); len(ftv) == 0 {
			continue
		} else {
			// Update the value
			value[k] = ftv

			// Insert the value in the temp storage
			if err := ts.insert(c.ft, key, ftv); err != nil {
				return err
			}
		}
	}

	// Iterate over the temp storage and set the values with len 1 to int
	ts.cleanSingleArrays()

	// Set the full-text cache to the temp map
	ts.updateFullText(c.ft)

	// Return nil for no errors
	return nil
}
