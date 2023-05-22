package nocache

import (
	"errors"
	"strings"
)

// SearchWithKey searches for all records containing the given query in the specified key column with a limit of results to return.
// The search result is limited to the specified number of entries.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a substring in the data value associated with the given key.
//   - key (string): The name of the key in the data whose data value should be searched.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query, key or limit is invalid.
func (ft *FullText) SearchWithKey(query string, key string, limit int) ([]map[string]any, error) {
	switch {
	case len(key) == 0:
		return []map[string]any{}, errors.New("invalid key")
	case len(query) == 0:
		return []map[string]any{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]any{}, errors.New("invalid limit")
	}

	// Set the query to lowercase
	query = strings.ToLower(query)

	// Lock the mutex
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()

	// Search for the query
	return ft.searchWithKey(query, key, limit), nil
}

// searchWithKey searches for all occurrences of the given query string in the FullText object's data associated with the specified key.
// The search is done by looking for the query as a substring in the data value associated with the given key.
// The search result is limited to the specified number of entries.
// Parameters:
//   - query (string): The search query to use. This string will be searched for as a substring in the data value associated with the given key.
//   - key (string): The name of the key in the data whose data value should be searched.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//
// Returns:
//   - []map[string]any: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//   - error: An error object. If no error occurs, this will be nil.
//
// Note: The search is case-insensitive.
func (ft *FullText) searchWithKey(query string, key string, limit int) []map[string]any {
	// Define variables
	var result []map[string]any = []map[string]any{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		if v, ok := ft.data[i][key].(string); !ok {
			continue
		} else {
			switch {
			case len(result) >= limit:
				return result
			case strings.Contains(strings.ToLower(v), query):
				result = append(result, ft.data[i])
			}
		}
	}

	// Return the result
	return result
}
