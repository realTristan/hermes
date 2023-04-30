package cache

import "errors"

// Add a value from the regular cache to the full text cache
// This only works if the key already exists in the cache, but
// not in the full text cache
func (c *Cache) FTAdd(key string) error {
	if !c.exists(key) {
		return errors.New("key does not exist in the cache")
	}
	return c.ft.set(key, c.data[key])
}
