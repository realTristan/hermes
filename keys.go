package hermes

// Keys is a method of the Cache struct that returns all the keys in the cache.
// This function is thread-safe.
//
// Returns:
//   - A slice of strings containing all the keys in the cache.
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.keys()
}

// keys is a method of the Cache struct that returns all the keys in the cache.
// This function is not thread-safe, and should only be called from an exported function.
//
// Returns:
//   - A slice of strings containing all the keys in the cache.
func (c *Cache) keys() []string {
	keys := make([]string, 0, len(c.data))
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}
