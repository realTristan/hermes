package hermes

// Values is a method of the Cache struct that gets all the values in the cache.
// This function is thread-safe.
//
// Returns:
//   - A slice of map[string]any representing all the values in the cache.
func (c *Cache) Values() []map[string]any {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

// values is a method of the Cache struct that returns all the values in the cache.
// This function is not thread-safe, and should only be called from an exported function.
//
// Returns:
//   - A slice of map[string]any representing all the values in the cache.
func (c *Cache) values() []map[string]any {
	var values []map[string]any = []map[string]any{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}
