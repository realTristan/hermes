package cache

// Get the number of items stored in the cache.
// This function is thread-safe.
func (c *Cache) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.length()
}

// Get the number of items stored in the cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) length() int {
	return len(c.data)
}
