package hermes

// Length is a method of the Cache struct that returns the number of items stored in the cache.
// This function is thread-safe.
//
// Returns:
//   - An integer representing the number of items stored in the cache.
func (c *Cache) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.length()
}

// length is a method of the Cache struct that returns the number of items stored in the cache.
// This function is not thread-safe, and should only be called from an exported function.
//
// Returns:
//   - An integer representing the number of items stored in the cache.
func (c *Cache) length() int {
	return len(c.data)
}
