package cache

// Get all the values in the cache.
// This function is thread-safe.
func (c *Cache) Values() []map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

// Get all the values in the cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) values() []map[string]interface{} {
	var values []map[string]interface{} = []map[string]interface{}{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}
