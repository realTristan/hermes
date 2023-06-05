package hermes

import (
	"errors"
	"strings"
)

// SearchValues searches for all records containing the given query in the specified schema with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid
func (c *Cache) SearchValues(sp SearchParams) ([]map[string]any, error) {
	// If the query is empty, return an error
	if len(sp.Query) == 0 {
		return []map[string]any{}, errors.New("invalid query")
	}

	// If no limit is provided, set it to 10
	if sp.Limit == 0 {
		sp.Limit = 10
	}

	// If no schema is provided, set it to all columns
	if len(sp.Schema) == 0 {
		sp.Schema = make(map[string]bool)
	}

	// Set the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Search the data
	return c.searchValues(sp), nil
}

// searchValues searches for all records containing the given query in the specified schema with a limit of results to return.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
func (c *Cache) searchValues(sp SearchParams) []map[string]any {
	// Define variables
	var result []map[string]any = []map[string]any{}

	// Iterate over the query result
	for _, item := range c.data {
		// Iterate over the keys and values for the data for that index
		for key, value := range item {
			switch {
			case len(result) >= sp.Limit:
				return result
			case !sp.Schema[key]:
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if strings.Contains(strings.ToLower(v), sp.Query) {
					result = append(result, item)
				}
			}
		}
	}

	// Return the result
	return result
}
