package hermes

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

// Cache struct
type Cache struct {
	mutex *sync.RWMutex
	cache map[string][]int
	keys  []string
	json  []map[string]string
}

// InitCache function
func InitCache(jsonFile string) *Cache {
	var cache Cache = Cache{
		mutex: &sync.RWMutex{},
		cache: map[string][]int{},
		keys:  []string{},
		json:  []map[string]string{},
	}

	// Load the json data
	cache._LoadJson(jsonFile)

	// Load the cache
	cache._LoadCache()

	// Return the cache
	return &cache
}

// Clean the cache
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._Clean()
}

// Clean the cache
func (c *Cache) _Clean() {
	c.cache = map[string][]int{}
	c.keys = []string{}
	c.json = []map[string]string{}
}

// Reset the cache
func (c *Cache) Reset(fileName string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c._Reset(fileName)
}

// Reset the cache
func (c *Cache) _Reset(fileName string) {
	c.cache = map[string][]int{}
	c.keys = []string{}
	c.json = []map[string]string{}
	c._LoadJson(fileName)
	c._LoadCache()
}

// Load the cache json data
func (c *Cache) _LoadJson(fileName string) {
	var data, _ = os.ReadFile(fileName)
	json.Unmarshal(data, &c.json)
}

// Load the cache
func (c *Cache) _LoadCache() {
	// Loop through the json data
	for i, item := range c.json {
		for _, value := range item {
			// Remove spaces from front and back
			value = strings.TrimSpace(value)

			// Remove double spaces
			value = _RemoveDoubleSpaces(value)

			// Convert to lowercase
			value = strings.ToLower(value)

			// Loop through the words
			for _, word := range strings.Split(value, " ") {
				// Make sure the word is not empty
				if len(word) <= 1 {
					continue
				}

				// If the word is not all alphabetic
				if !_IsAlphaNum(word) {
					continue
				}

				// If the key doesn't exist in the cache
				if _, ok := c.cache[word]; !ok {
					c.cache[word] = []int{i}
					c.keys = append(c.keys, word)
					continue
				}

				// If the index is already in the cache
				if _ContainsInt(c.cache[word], i) {
					continue
				}

				// Append the index to the cache
				c.cache[word] = append(c.cache[word], i)
			}
		}
	}
}
