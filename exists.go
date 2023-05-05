package cache

// Check if a key exists in the cache.
// This function is thread-safe.
func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.exists(key)
}

// Check if a key exists in the cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) exists(key string) bool {
	_, ok := c.data[key]
	return ok
}
