package hermes

import (
	"errors"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// SearchOneWord searches for a single word in the FullText struct's data and returns a list of maps containing the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid or if the full-text is not initialized.
func (c Cache) SearchOneWord(sp SearchParams) ([]map[string]any, error) {
	// If the query is empty, return an error
	if len(sp.Query) == 0 {
		return []map[string]any{}, errors.New("invalid query")
	}

	// If no limit is provided, set it to 10
	if sp.Limit == 0 {
		sp.Limit = 10
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the full-text is initialized
	if c.ft == nil {
		return []map[string]any{}, errors.New("full-text is not initialized")
	}

	// Search the data
	return c.searchOneWord(sp), nil
}

// searchOneWord searches for a single word in the FullText struct's data and returns a list of maps containing the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
func (c *Cache) searchOneWord(sp SearchParams) []map[string]any {
	// Set the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Define variables
	var result []map[string]any = []map[string]any{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if sp.Strict {
		return c.searchOneWordStrict(result, sp)
	}

	// Define a map to store the indices that have already been added
	var alreadyAdded map[int]int = map[int]int{}

	// Loop through the cache keys
	for k, v := range c.ft.storage {
		switch {
		case len(result) >= sp.Limit:
			return result
		case !Utils.Contains(k, sp.Query):
			continue
		}

		// Loop through the cache indices
		if index, ok := v.(int); ok {
			if _, ok := alreadyAdded[index]; ok {
				continue
			}
			result = append(result, c.data[c.ft.indices[index]])
			alreadyAdded[index] = 0
			continue
		}

		var indices []int = v.([]int)
		for j := 0; j < len(indices); j++ {
			if _, ok := alreadyAdded[indices[j]]; ok {
				continue
			}

			// Else, append the index to the result
			result = append(result, c.data[c.ft.indices[indices[j]]])
			alreadyAdded[indices[j]] = 0
		}
	}

	// Return the result
	return result
}

// searchOneWordStrict is a method of the Cache struct that searches for a single word in the cache and returns the results.
// This function is not thread-safe.
//
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - A slice of map[string]any representing the search results.
func (c *Cache) searchOneWordStrict(result []map[string]any, sp SearchParams) []map[string]any {
	// Check if the query is in the cache
	if _, ok := c.ft.storage[sp.Query]; !ok {
		return result
	}

	// If there's only one result
	if v, ok := c.ft.storage[sp.Query].(int); ok {
		return []map[string]any{c.data[c.ft.indices[v]]}
	}

	// Loop through the indices
	for i := 0; i < len(c.ft.storage[sp.Query].([]int)); i++ {
		if len(result) >= sp.Limit {
			return result
		}
		var (
			index int    = c.ft.storage[sp.Query].([]int)[i]
			key   string = c.ft.indices[index]
		)
		result = append(result, c.data[key])
	}

	// Return the result
	return result
}
