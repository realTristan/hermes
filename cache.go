package hermes

import "sync"

// Cache struct
type Cache struct {
	data  map[string]map[string]string
	mutex *sync.RWMutex
	fts   *FTS
}

// Initialize the cache
func InitCache() *Cache {
	return &Cache{
		data:  map[string]map[string]string{},
		mutex: &sync.RWMutex{},
		fts: &FTS{
			mutex: &sync.RWMutex{},
			cache: map[string][]int{},
			keys:  []string{},
			json:  []map[string]string{},
		},
	}
}

// Initialize the FTS cache
func (c *Cache) InitFTS() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._InitFTS()
}
func (c *Cache) _InitFTS() {
	// Convert the cache data into an array of maps
	var jsonArray []map[string]string = []map[string]string{}
	for _, key := range c._Keys() {
		jsonArray = append(jsonArray, c._Get(key))
	}
	c.fts = &FTS{
		mutex: &sync.RWMutex{},
		cache: map[string][]int{},
		keys:  []string{},
		json:  jsonArray,
	}
	c.fts._LoadCacheJson(jsonArray)
}

// Clean the cache
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._Clean()
}
func (c *Cache) _Clean() {
	c.data = map[string]map[string]string{}
}

// Set a value in the cache
func (c *Cache) Set(key string, value map[string]string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._Set(key, value)
}
func (c *Cache) _Set(key string, value map[string]string) {
	c.data[key] = value

	// update the value in the FTS cache
	if c.fts != nil {
		c.fts._Set(key, value)
	}
}

// Get a value from the cache
func (c *Cache) Get(key string) map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Get(key)
}
func (c *Cache) _Get(key string) map[string]string {
	return c.data[key]
}

// Delete a key from the cache
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._Delete(key)
}
func (c *Cache) _Delete(key string) {
	delete(c.data, key)
}

// Check if a key exists in the cache
func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Exists(key)
}
func (c *Cache) _Exists(key string) bool {
	_, ok := c.data[key]
	return ok
}

// Get all the keys in the cache
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Keys()
}
func (c *Cache) _Keys() []string {
	keys := []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

// Get all the values in the cache
func (c *Cache) Values() []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Values()
}
func (c *Cache) _Values() []map[string]string {
	var values []map[string]string = []map[string]string{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}

// Get the length of the cache
func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Len()
}
func (c *Cache) _Len() int {
	return len(c.data)
}

// Search the cache
func (c *Cache) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.fts.Search(query, limit, strict)
}

// Search the cache with spaces
func (c *Cache) SearchWithSpaces(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.fts.SearchWithSpaces(query, limit, strict, keys)
}
