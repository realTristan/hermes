package nocache

import (
	"errors"
	"strings"

	utils "github.com/realTristan/hermes/utils"
)

// SearchOneWord searches for a single query within the data using a full-text search approach.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include exact matches.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a single word in any value associated with any key in each entry of the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If true, only exact matches will be returned. If false, partial matches will also be returned.
//
// Returns:
//   - []map[string]any: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//   - error: An error object. If no error occurs, this will be nil.
//
// Note: The search is case-insensitive.
func (ft *FullText) SearchOneWord(sp SearchParams) ([]map[string]any, error) {
	switch {
	case len(sp.Query) == 0:
		return []map[string]any{}, errors.New("invalid query")
	case sp.Limit < 1:
		return []map[string]any{}, errors.New("invalid limit")
	}

	// Set the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Lock the mutex
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()

	// Search the data
	return ft.searchOneWord(sp), nil
}

// searchOneWord searches for a single query within the data using a full-text search approach.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include exact matches.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a single word in any value associated with any key in each entry of the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If true, only exact matches will be returned. If false, partial matches will also be returned.
//
// Returns:
//   - []map[string]any: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//
// Note: The search is case-insensitive.
func (ft *FullText) searchOneWord(sp SearchParams) []map[string]any {
	// Define the result variable
	var result []map[string]any = []map[string]any{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if sp.Strict {
		return ft.searchOneWordStrict(result, sp)
	}

	// true for already checked
	var alreadyAdded map[int]int = map[int]int{}

	// Loop through the cache keys
	for i := 0; i < len(ft.words); i++ {
		switch {
		case len(result) >= sp.Limit:
			return result
		case !utils.Contains(ft.words[i], sp.Query):
			continue
		}

		// Loop through the cache indices
		if v, ok := ft.storage[ft.words[i]].(int); ok {
			if _, ok := alreadyAdded[v]; ok {
				continue
			}
			result = append(result, ft.data[v])
			alreadyAdded[v] = 0
			continue
		}

		var indices []int = ft.storage[ft.words[i]].([]int)
		for j := 0; j < len(indices); j++ {
			if _, ok := alreadyAdded[indices[j]]; ok {
				continue
			}
			result = append(result, ft.data[indices[j]])
			alreadyAdded[indices[j]] = 0
		}
	}

	// Return the result
	return result
}

// searchOneWordStrict is a method of the FullText struct that searches for a single word in the full-text cache and returns the results.
//
// Parameters:
//   - result: A slice of map[string]any representing the current search results.
//   - query: A string representing the word to search for.
//   - limit: An integer representing the maximum number of results to return.
//
// Returns:
//   - A slice of map[string]any representing the search results.
func (ft *FullText) searchOneWordStrict(result []map[string]any, sp SearchParams) []map[string]any {
	// Check if the query is in the cache
	if _, ok := ft.storage[sp.Query]; !ok {
		return result
	}

	// Check if the cache value is an integer
	if v, ok := ft.storage[sp.Query].(int); ok {
		return []map[string]any{ft.data[v]}
	}

	// Loop through the indices
	var indices []int = ft.storage[sp.Query].([]int)
	for i := 0; i < len(indices); i++ {
		if len(result) >= sp.Limit {
			return result
		}
		result = append(result, ft.data[indices[i]])
	}

	// Return the result
	return result
}
