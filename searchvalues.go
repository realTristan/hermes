package cache

import (
	"errors"
	"strings"
)

// SearchValues searches for all records containing the given query in the specified schema with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - query (string): The search query to match against data in the FullText struct
//   - limit (int): The maximum number of results to be returned
//   - schema (map[string]bool): A map of column names in the FullText struct's data that should be searched. The boolean value indicates whether the column should be searched or not.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid
func (c *Cache) SearchValues(query string, limit int, schema map[string]bool) ([]map[string]interface{}, error) {
	switch {
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
	return c.searchValues(query, limit, schema), nil
}

// searchValues searches for all records containing the given query in the specified schema with a limit of results to return.
// Parameters:
//   - query (string): The search query to match against data in the FullText struct
//   - limit (int): The maximum number of results to be returned
//   - schema (map[string]bool): A map of column names in the FullText struct's data that should be searched. The boolean value indicates whether the column should be searched or not.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
func (c *Cache) searchValues(query string, limit int, schema map[string]bool) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range c.data {
		// Iterate over the keys and values for the data for that index
		for key, value := range item {
			switch {
			case len(result) >= limit:
				return result
			case !schema[key]:
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if strings.Contains(strings.ToLower(v), query) {
					result = append(result, item)
				}
			}
		}
	}

	// Return the result
	return result
}
