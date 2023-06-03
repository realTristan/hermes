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
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//   - error: An error object. If no error occurs, this will be nil.
//
// Note: The search is case-insensitive.
func (ft *FullText) SearchValues(sp SearchParams) ([]map[string]any, error) {
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
	return ft.searchValues(sp), nil
}

// searchValues searches for all occurrences of the given query string in the FullText object's data.
// The search is done by looking for the query as a substring in any value associated with any key in each entry of the data.
// The search result is limited to the specified number of entries, and can optionally be filtered to only include keys that match a given schema.
// Parameters:
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: An array of maps representing the search results. Each map contains key-value pairs
//     from the entry in the data that matched the search query. If no results are found, an empty array is returned.
//
// Note: The search is case-insensitive.
func (ft *FullText) searchValues(sp SearchParams) []map[string]any {
	// Define variables
	var result []map[string]any = []map[string]any{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		// Iterate over the keys and values for the data
		for key, value := range ft.data[i] {
			if v, ok := value.(string); !ok {
				continue
			} else {
				switch {
				case len(result) >= sp.Limit:
					return result
				case !sp.Schema[key]:
					continue
				case strings.Contains(strings.ToLower(v), sp.Query):
					result = append(result, ft.data[i])
				}
			}
		}
	}

	// Return the result
	return result
}
