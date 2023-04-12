package Hermes

import (
	"encoding/json"
	"os"
	"regexp"
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

// SearchMultiple function with lock
func (c *Cache) SearchMultiple(query string, limit int, strict bool) []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchMultiple(query, limit, strict)
}

// Search for multiple words
func (c *Cache) _SearchMultiple(query string, limit int, strict bool) []map[string]string {
	// Split the query into words
	var words []string = strings.Split(strings.TrimSpace(query), " ")

	// If the words array is empty
	if len(words) == 0 {
		return []map[string]string{}
	}

	// If there's only one word, return the result
	if len(words) == 1 {
		var r, _ = c.Search(query, limit, strict)
		return r
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}
	var alreadyAdded []int = []int{}

	// Loop through the words and get the indices that are common
	for i := range words {
		// Search for the query inside the cache
		var queryResult, indices = c.Search(words[i], limit, strict)

		// Iterate over the result
		for j := range queryResult {
			// If the data's already been added
			if _ContainsInt(alreadyAdded, j) {
				continue
			}

			// Iterate over the keys and values for the json data for that index
			for _, v := range queryResult[j] {
				// If the data contains the query
				if strings.Contains(v, query) {
					result = append(result, queryResult[j])
					alreadyAdded = append(alreadyAdded, indices[j])
				}
			}
		}
	}

	// Return the result
	return result
}

// Search function with lock
func (c *Cache) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Search(query, limit, strict)
}

// Search for a single query
func (c *Cache) _Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]string{}, []int{}
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Create an array to store the indices that have already
	// been added to the result array
	var indices []int = []int{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := c.cache[query]; !ok {
			return result, indices
		}

		// Loop through the indices
		indices = c.cache[query]
		for i := range indices {
			result = append(result, c.json[indices[i]])
		}

		// Return the result
		return result, indices
	}

	// Loop through the cache keys
	for i := range c.keys {
	Switch:
		switch {
		// Check if the limit has been reached
		case len(result) >= limit:
			return result, indices

		// If the key doesn't start with the word
		case !strings.HasPrefix(c.keys[i], query):
			continue

		// Check if the key is shorter than the query
		case len(c.keys[i]) < len(query):
			continue

		// Check if the key is equal to the query
		case c.keys[i] == query:
			break Switch

		// Check if the key contains the query
		case !strings.Contains(c.keys[i], query):
			continue

		// Check if the index is already in the result
		case _ContainsInt(indices, i):
			continue
		}

		// Loop through the indices
		for j := range c.cache[c.keys[i]] {
			var index int = c.cache[c.keys[i]][j]

			// Else, append the index to the result
			result = append(result, c.json[index])
			indices = append(indices, index)
		}
	}

	// Return the result
	return result, indices
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

// Check if a string is all alphabetic
func _IsAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string
func _RemoveDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if an int is in an array
func _ContainsInt(array []int, value int) bool {
	for i := range array {
		if array[i] == value {
			return true
		}
	}
	return false
}
