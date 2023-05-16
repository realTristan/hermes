package cache

import (
	"errors"
	"strings"
)

// SearchWithKey searches for all records containing the given query in the specified key column with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - query (string): The search query to match against data in the FullText struct
//   - key (string): The key to search data on
//   - limit (int): The maximum number of results to be returned
//
// Returns:
//   - []map[string]interface{}: A slice of maps containing the search results
//   - error: An error if the key, query or limit is invalid
func (c *Cache) SearchWithKey(query string, key string, limit int) ([]map[string]interface{}, error) {
	switch {
	case len(key) == 0:
		return []map[string]interface{}{}, errors.New("invalid key")
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Set the query to lowercase
	query = strings.ToLower(query)

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Search the data
	return c.searchWithKey(query, key, limit), nil
}

// searchWithKey searches for all records containing the given query in the specified key column with a limit of results to return.
// Parameters:
//   - query (string): The search query to match against data in the FullText struct
//   - key (string): The key to search data on
//   - limit (int): The maximum number of results to be returned
//
// Returns:
//   - []map[string]interface{}: A slice of maps containing the search results
func (c *Cache) searchWithKey(query string, key string, limit int) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range c.data {
		for _, v := range item {
			if len(result) >= limit {
				return result
			}

			// Check if the value contains the query
			if v, ok := v.(string); ok {
				if strings.Contains(strings.ToLower(v), query) {
					result = append(result, item)
				}
			}
		}
	}

	// Return the result
	return result
}
