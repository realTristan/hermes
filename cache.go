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
		fts:   nil,
	}
}

// Initialize the FTS cache with Mutex Locking
func (c *Cache) InitFTS(maxKeys int, maxSizeBytes int, keySettings map[string]bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.initFTS(maxKeys, maxSizeBytes, keySettings)
}

// Initialize the FTS cache
func (c *Cache) initFTS(maxKeys int, maxSizeBytes int, keySettings map[string]bool) error {
	// Convert the cache data into an array of maps
	var data []map[string]string = []map[string]string{}
	for _, key := range c.keys() {
		data = append(data, c.get(key))
	}
	c.fts = &FTS{
		mutex:        &sync.RWMutex{},
		cache:        map[string][]int{},
		keys:         []string{},
		json:         data,
		maxKeys:      maxKeys,
		maxSizeBytes: maxSizeBytes,
	}
	return c.fts.loadCacheJson(data, keySettings)
}

// Clean the cache with Mutex Locking
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clean()
}

// Clean the cache
func (c *Cache) clean() {
	c.data = map[string]map[string]string{}
}

// Set a value in the cache with Mutex Locking
func (c *Cache) Set(key string, value map[string]string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

// Set a value in the cache
func (c *Cache) set(key string, value map[string]string) error {
	c.data[key] = value

	// update the value in the FTS cache
	if c.fts != nil {
		return c.fts.set(key, value)
	}
	return nil
}

// Get a value from the cache with Mutex Locking
func (c *Cache) Get(key string) map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.get(key)
}

// Get a value from the cache
func (c *Cache) get(key string) map[string]string {
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
func (c *Cache) Values() []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

// Get all the values in the cache
func (c *Cache) values() []map[string]string {
	var values []map[string]string = []map[string]string{}
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

// Search the cache with Mutex Locking
func (c *Cache) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.fts.Search(query, limit, strict)
}

// Search the cache with spaces with Mutex Locking
func (c *Cache) SearchWithSpaces(query string, limit int, strict bool, keySettings map[string]bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.fts.SearchWithSpaces(query, limit, strict, keySettings)
}
