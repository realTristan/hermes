package cache

import (
	"errors"
	"strings"
)

// Search searches for a query by splitting the query into separate words and returning the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - query (string): The search query
//   - limit (int): The limit of search results to return
//   - strict (bool): A boolean to indicate whether the search should be strict or not
//   - schema (map[string]bool): A map containing the schema to search for
//
// Returns:
//   - []map[string]interface{}: A slice of maps containing the search results
//   - error: An error if the query or limit is invalid, or if the full-text index is not initialized
func (c *Cache) Search(query string, limit int, strict bool, schema map[string]bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the FT index is initialized
	if c.ft == nil {
		return []map[string]interface{}{}, errors.New("full-text is not initialized")
	}

	// Set the query to lowercase
	query = strings.ToLower(query)

	// Search for the query
	return c.search(query, limit, strict, schema)
}

// search searches for a query by splitting the query into separate words and returning the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - query (string): The search query
//   - limit (int): The limit of search results to return
//   - strict (bool): A boolean to indicate whether the search should be strict or not
//   - schema (map[string]bool): A map containing the schema to search for
//
// Returns:
//   - []map[string]interface{}: A slice of maps containing the search results
//   - error: An error if the query is invalid or if the smallest words array is not found in the cache
func (c *Cache) search(query string, limit int, strict bool, schema map[string]bool) ([]map[string]interface{}, error) {
	// Split the query into separate words
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	// Get the search result of the first word
	case len(words) == 1:
		return c.searchOneWord(words[0], limit, strict), nil
	}

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Variables for storing the smallest words array
	var (
		smallest      int = 0
		smallestIndex int = 0
	)

	// Check if the query is in the cache
	if v, ok := c.ft.storage[words[0]]; !ok {
		return []map[string]interface{}{}, errors.New("invalid query")
	} else {
		smallest = len(v)
	}

	// Find the smallest words array
	// Don't include the first or last words from the query
	for i := 1; i < len(words)-1; i++ {
		if v, ok := c.ft.storage[words[i]]; ok {
			if len(v) < smallest {
				smallest = len(v)
				smallestIndex = i
			}
		}
	}

	// Loop through the indices
	var keys []int = c.ft.storage[words[smallestIndex]]
	for i := 0; i < len(keys); i++ {
		for key, value := range c.data[c.ft.indices[keys[i]]] {
			if !schema[key] {
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if strings.Contains(strings.ToLower(v), query) {
					result = append(result, c.data[c.ft.indices[keys[i]]])
				}
			}
		}
	}

	// Return the result
	return result, nil
}
