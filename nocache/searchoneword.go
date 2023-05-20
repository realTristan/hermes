package nocache

import (
	"errors"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// SearchOneWord searches for a single query within the data using a full-text search approach.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include exact matches.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a single word in any value associated with any key in each entry of the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If true, only exact matches will be returned. If false, partial matches will also be returned.
//
// Returns:
//   - []map[string]interface{}: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//   - error: An error object. If no error occurs, this will be nil.
//
// Note: The search is case-insensitive.
func (ft *FullText) SearchOneWord(query string, limit int, strict bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Set the query to lowercase
	query = strings.ToLower(query)

	// Lock the mutex
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()

	// Search the data
	return ft.searchOneWord(query, limit, strict), nil
}

// searchOneWord searches for a single query within the data using a full-text search approach.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include exact matches.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a single word in any value associated with any key in each entry of the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If true, only exact matches will be returned. If false, partial matches will also be returned.
//
// Returns:
//   - []map[string]interface{}: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//
// Note: The search is case-insensitive.
func (ft *FullText) searchOneWord(query string, limit int, strict bool) []map[string]interface{} {
	// Define the result variable
	var result []map[string]interface{} = []map[string]interface{}{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := ft.storage[query]; !ok {
			return result
		}

		// Loop through the indices
		var indices []int = ft.storage[query]
		for i := 0; i < len(indices); i++ {
			if len(result) >= limit {
				return result
			}
			result = append(result, ft.data[indices[i]])
		}

		// Return the result
		return result
	}

	// true for already checked
	var alreadyAdded map[int]int = map[int]int{}

	// Loop through the cache keys
	for i := 0; i < len(ft.words); i++ {
		switch {
		case len(result) >= limit:
			return result
		case !Utils.Contains(ft.words[i], query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(ft.storage[ft.words[i]]); j++ {
			var index int = ft.storage[ft.words[i]][j]
			if _, ok := alreadyAdded[index]; ok {
				continue
			}
			result = append(result, ft.data[index])
			alreadyAdded[index] = 0
		}
	}

	// Return the result
	return result
}
