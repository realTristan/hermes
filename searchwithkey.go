package cache

import (
	"errors"
	"strings"
)

// SearchWithKey searches for all records containing the given query in the specified key column with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps containing the search results
//   - error: An error if the key, query or limit is invalid
func (c *Cache) SearchWithKey(sp SearchParams) ([]map[string]any, error) {
	switch {
	case len(sp.Key) == 0:
		return []map[string]any{}, errors.New("invalid key")
	case len(sp.Query) == 0:
		return []map[string]any{}, errors.New("invalid query")
	}

	// Set the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Search the data
	return c.searchWithKey(sp), nil
}

// searchWithKey searches for all records containing the given query in the specified key column with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps containing the search results
func (c *Cache) searchWithKey(sp SearchParams) []map[string]any {
	// Define variables
	var result []map[string]any = []map[string]any{}

	// Iterate over the query result
	for _, item := range c.data {
		for _, v := range item {
			if len(result) >= sp.Limit {
				return result
			}

			// Check if the value contains the query
			if v, ok := v.(string); ok {
				if strings.Contains(strings.ToLower(v), sp.Query) {
					result = append(result, item)
				}
			}
		}
	}

	// Return the result
	return result
}
