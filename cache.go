package main

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

// Search
func (c *Cache) Search(word string, limit int, strict bool) []map[string]string {
	// Lock the cache
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Convert the word to lowercase
	word = strings.ToLower(word)

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the word is in the cache
		if _, ok := c.cache[word]; !ok {
			return result
		}

		// Loop through the indices
		var indices []int = c.cache[word]
		for i := 0; i < len(indices); i++ {
			result = append(result, c.json[indices[i]])
		}

		// Return the result
		return result
	}

	// Create an array to store the indices that have already
	// been added to the result array
	var indices []int = []int{}

	// Loop through the cache keys
	for i := 0; i < len(c.keys); i++ {
		switch {
		// Check if the limit has been reached
		case len(result) >= limit:
			return result

		// The word doesn't start with the same letter
		case c.keys[i][0] != word[0]:
			continue

		// Check if the key is shorter than the word
		case len(c.keys[i]) < len(word):
			continue

		// Check if the key is equal to the word
		case c.keys[i] == word:
			// Loop through the indices
			for j := 0; j < len(c.cache[c.keys[i]]); j++ {
				var index int = c.cache[c.keys[i]][j]

				// Else, append the index to the result
				result = append(result, c.json[index])
			}

			// Return the result
			return result

		// Check if the key contains the word
		case !strings.Contains(c.keys[i], word):
			continue

		// Check if the index is already in the result
		case ContainsInt(indices, i):
			continue
		}

		// Loop through the indices
		for j := 0; j < len(c.cache[c.keys[i]]); j++ {
			var index int = c.cache[c.keys[i]][j]

			// Else, append the index to the result
			result = append(result, c.json[index])
			indices = append(indices, index)
		}
	}

	// Return the result
	return result
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
		for _, runeValue := range item {
			// Convert the rune to a string
			var value string = string(runeValue)

			// Remove spaces from front and back
			value = strings.TrimSpace(value)

			// Remove double spaces
			value = RemoveDoubleSpaces(value)

			// Convert to lowercase
			value = strings.ToLower(value)

			// Split the string into an array
			var array []string = strings.Split(value, " ")

			// Loop through the array
			for _, word := range array {
				// If the word is not all alphabetic
				if !isAlphaNum(word) {
					continue
				}

				// If the key doesn't exist in the cache
				if _, ok := c.cache[word]; !ok {
					c.cache[word] = []int{i}
					c.keys = append(c.keys, word)
					continue
				}

				// If the index is already in the cache
				if ContainsInt(c.cache[word], i) {
					continue
				}
				c.cache[word] = append(c.cache[word], i)
			}
		}
	}
}

// Check if a string is all alphabetic
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string
func RemoveDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if an int is in an array
func ContainsInt(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string is in an array
func ContainsString(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
