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
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid or if the full-text is not initialized.
func (c Cache) SearchOneWord(query string, limit int, strict bool) ([]map[string]any, error) {
	switch {
	case len(query) == 0:
		return []map[string]any{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]any{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the full-text is initialized
	if c.ft == nil {
		return []map[string]any{}, errors.New("full-text is not initialized")
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
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
func (c *Cache) searchOneWord(query string, limit int, strict bool) []map[string]any {
	// Set the query to lowercase
	query = strings.ToLower(query)

	// Define variables
	var result []map[string]any = []map[string]any{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		return c.searchOneWordStrict(result, query, limit)
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
// This function is thread-safe.
//
// Parameters:
//   - result: A slice of map[string]any representing the current search results.
//   - query: A string representing the word to search for.
//   - limit: An integer representing the maximum number of results to return.
//
// Returns:
//   - A slice of map[string]any representing the search results.
func (c *Cache) searchOneWordStrict(result []map[string]any, query string, limit int) []map[string]any {
	// Check if the query is in the cache
	if _, ok := c.ft.storage[query]; !ok {
		return result
	}

	// If there's only one result
	if v, ok := c.ft.storage[query].(int); ok {
		return []map[string]any{c.data[c.ft.indices[v]]}
	}

	// Loop through the indices
	for i := 0; i < len(c.ft.storage[query].([]int)); i++ {
		if len(result) >= limit {
			return result
		}
		var (
			index int    = c.ft.storage[query].([]int)[i]
			key   string = c.ft.indices[index]
		)
		result = append(result, c.data[key])
	}

	// Return the result
	return result
}
