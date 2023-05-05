package cache

// This function retrieves the value associated with the given key from the cache.
// This function is thread-safe.
func (c *Cache) Get(key string) map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.get(key)
}

// This function retrieves the value associated with the given key from the cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) get(key string) map[string]interface{} {
	return c.data[key]
}
