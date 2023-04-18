package cache

import "sync"

// Cache struct
type Cache struct {
	data  map[string]map[string]interface{}
	mutex *sync.RWMutex
	FT    *FullText
}

// Initialize the cache
func InitCache() *Cache {
	return &Cache{
		data:  map[string]map[string]interface{}{},
		mutex: &sync.RWMutex{},
		FT:    nil,
	}
}

// Clean the cache with Mutex Locking
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clean()
}

// Clean the cache
func (c *Cache) clean() {
	c.data = map[string]map[string]interface{}{}
}

// Set a value in the cache with Mutex Locking
func (c *Cache) Set(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

// Set a value in the cache
func (c *Cache) set(key string, value map[string]interface{}) error {
	// Update the value in the FT cache
	if c.FT != nil && c.FT.isInitialized {
		if err := c.FT.Set(key, value); err != nil {
			return err
		}
	}

	// Update the value in the cache
	c.data[key] = value

	// Return nil for no error
	return nil
}

// Get a value from the cache with Mutex Locking
func (c *Cache) Get(key string) map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.get(key)
}

// Get a value from the cache
func (c *Cache) get(key string) map[string]interface{} {
	return c.data[key]
}

// Delete a key from the cache with Mutex Locking
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.delete(key)
}

// Delete a key from the cache
func (c *Cache) delete(key string) {
	// Delete the key from the FT cache
	if c.FT != nil && c.FT.isInitialized {
		c.FT.Delete(key)
	}

	// Delete the key from the cache
	delete(c.data, key)
}

// Check if a key exists in the cache with Mutex Locking
func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.exists(key)
}

// Check if a key exists in the cache
func (c *Cache) exists(key string) bool {
	_, ok := c.data[key]
	return ok
}

// Get all the keys in the cache with Mutex Locking
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.keys()
}

// Get all the keys in the cache
func (c *Cache) keys() []string {
	keys := []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

// Get all the values in the cache with Mutex Locking
func (c *Cache) Values() []map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

// Get all the values in the cache
func (c *Cache) values() []map[string]interface{} {
	var values []map[string]interface{} = []map[string]interface{}{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}

// Get the length of the cache with Mutex Locking
func (c *Cache) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.length()
}

// Get the length of the cache
func (c *Cache) length() int {
	return len(c.data)
}
