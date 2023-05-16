package cache

import (
	"errors"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// SearchOneWord searches for a single word in the FullText struct's data and returns a list of maps containing the search results.
// Parameters:
//   - query (string): The search query to match against data in the FullText struct
//   - limit (int): The maximum number of results to be returned. If set to 0, there is no limit to the number of results returned.
//   - strict (bool): Determines whether the search should be strict or not. If set to true, only exact matches will be returned.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid or if the full-text is not initialized.
func (c *Cache) SearchOneWord(query string, limit int, strict bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the full-text is initialized
	if c.ft == nil {
		return []map[string]interface{}{}, errors.New("full-text is not initialized")
	}

	// Search the data
	return c.searchOneWord(query, limit, strict), nil
}

// searchOneWord searches for a single word in the FullText struct's data and returns a list of maps containing the search results.
// Parameters:
//   - query (string): The search query to match against data in the FullText struct
//   - limit (int): The maximum number of results to be returned. If set to 0, there is no limit to the number of results returned.
//   - strict (bool): Determines whether the search should be strict or not. If set to true, only exact matches will be returned.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
func (c *Cache) searchOneWord(query string, limit int, strict bool) []map[string]interface{} {
	// Set the query to lowercase
	query = strings.ToLower(query)

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := c.ft.storage[query]; !ok {
			return result
		}

		// Loop through the indices
		for i := 0; i < len(c.ft.storage[query]); i++ {
			if len(result) >= limit {
				return result
			}
			var (
				index int    = c.ft.storage[query][i]
				key   string = c.ft.indices[index]
			)
			result = append(result, c.data[key])
		}

		// Return the result
		return result
	}

	// Define a map to store the indices that have already been added
	var alreadyAdded map[int]int = map[int]int{}

	// Loop through the cache keys
	for k, v := range c.ft.storage {
		switch {
		case len(result) >= limit:
			return result
		case !Utils.Contains(k, query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(v); j++ {
			if _, ok := alreadyAdded[v[j]]; ok {
				continue
			}

			// Else, append the index to the result
			result = append(result, c.data[c.ft.indices[v[j]]])
			alreadyAdded[v[j]] = 0
		}
	}

	// Return the result
	return result
}
