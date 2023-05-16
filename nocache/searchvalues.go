package nocache

import (
	"errors"
	"strings"
)

// SearchValues searches for all occurrences of the given query string in the FullText object's data.
// The search is done by splitting the query into separate words and looking for each of them in the data.
// The search result is limited to the specified number of entries, and can optionally be filtered to only
// include keys that match a given schema.
// Parameters:
//   - query (string): The search query to use. This string will be split into separate words and each word will be searched for in the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - schema (map[string]bool): A map of keys to boolean values. If a key is present in this map and its value is true,
//     only entries that have a non-empty value for that key will be returned.
//
// Returns:
//   - []map[string]interface{}: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//   - error: An error object. If no error occurs, this will be nil.
//
// Note: The search is case-insensitive.
func (ft *FullText) SearchValues(query string, limit int, schema map[string]bool) ([]map[string]interface{}, error) {
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
	return ft.searchValues(query, limit, schema), nil
}

// searchValues searches for all occurrences of the given query string in the FullText object's data.
// The search is done by looking for the query as a substring in any value associated with any key in each entry of the data.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include keys that match a given schema.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a substring in any value associated with any key in each entry of the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - schema (map[string]bool): A map of keys to boolean values. If a key is present in this map and its value is true,
//     only entries that have a non-empty value for that key will be returned.
//
// Returns:
//   - []map[string]interface{}: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//
// Note: The search is case-insensitive.
func (ft *FullText) searchValues(query string, limit int, schema map[string]bool) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		// Iterate over the keys and values for the data
		for key, value := range ft.data[i] {
			if v, ok := value.(string); !ok {
				continue
			} else {
				switch {
				case len(result) >= limit:
					return result
				case !schema[key]:
					continue
				case strings.Contains(strings.ToLower(v), query):
					result = append(result, ft.data[i])
				}
			}
		}
	}

	// Return the result
	return result
}
