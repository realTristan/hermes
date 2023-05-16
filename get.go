package cache

// Get is a method of the Cache struct that retrieves the value associated with the given key from the cache.
// This method is thread-safe.
//
// Parameters:
//   - key: A string representing the key to retrieve the value for.
//
// Returns:
//   - A map[string]interface{} representing the value associated with the given key in the cache.
func (c *Cache) Get(key string) map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.get(key)
}

// get is a method of the Cache struct that retrieves the value associated with the given key from the cache.
// This method is not thread-safe and should only be called from an exported function.
//
// Parameters:
//   - key: A string representing the key to retrieve the value for.
//
// Returns:
//   - A map[string]interface{} representing the value associated with the given key in the cache.
func (c *Cache) get(key string) map[string]interface{} {
	return c.data[key]
}
