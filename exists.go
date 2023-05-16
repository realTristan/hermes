package cache

// Exists is a method of the Cache struct that checks if a key exists in the cache.
// This method is thread-safe.
//
// Parameters:
//   - key: A string representing the key to check for existence in the cache.
//
// Returns:
//   - A boolean value indicating whether the key exists in the cache or not.
func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.exists(key)
}

// exists is a method of the Cache struct that checks if a key exists in the cache.
// This method is not thread-safe and should only be called from an exported function.
//
// Parameters:
//   - key: A string representing the key to check for existence in the cache.
//
// Returns:
//   - A boolean value indicating whether the key exists in the cache or not.
func (c *Cache) exists(key string) bool {
	_, ok := c.data[key]
	return ok
}
