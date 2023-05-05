package cache

// Get all the keys in the cache.
// This function is thread-safe.
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.keys()
}

// Get all the keys in the cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) keys() []string {
	var keys []string = []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}
