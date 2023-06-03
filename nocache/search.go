package nocache

import (
	"errors"
	"strings"
)

// Search searches for all occurrences of the given query string in the FullText object's data.
// The search is done by splitting the query into separate words and looking for each of them in the data.
// The search result is limited to the specified number of entries.
// Parameters:
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid.
func (ft *FullText) Search(sp SearchParams) ([]map[string]any, error) {
	switch {
	case len(sp.Query) == 0:
		return []map[string]any{}, errors.New("invalid query")
	case sp.Limit < 1:
		return []map[string]any{}, errors.New("invalid limit")
	}

	// Convert the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Lock the mutex
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()

	// Perform the search
	return ft.search(sp), nil
}

// search searches for all occurrences of the given query string in the FullText object's data.
// The search is done by splitting the query into separate words and looking for each of them in the data.
// The search result is limited to the specified number of entries.
// Parameters:
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid.
func (ft *FullText) search(sp SearchParams) []map[string]any {
	// Split the query into separate words
	var words []string = strings.Split(strings.TrimSpace(sp.Query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]any{}
	// Get the search result of the first word
	case len(words) == 1:
		sp.Query = words[0]
		return ft.searchOneWord(sp)
	}

	// Check if the query is in the cache
	if _, ok := ft.storage[words[0]]; !ok {
		return []map[string]any{}
	}

	// Define variables
	var result []map[string]any = []map[string]any{}

	// Variables for storing the smallest words array
	var (
		smallest      int = 0
		smallestIndex int = 0
	)

	// Check if the query is in the cache
	if v, ok := ft.storage[words[0]]; !ok {
		return []map[string]any{}
	} else {
		if index, ok := v.(int); ok {
			return []map[string]any{
				ft.data[index],
			}
		}
		smallest = len(v.([]int))
	}

	// Find the smallest words array
	// Don't include the first or last words from the query
	for i := 1; i < len(words)-1; i++ {
		if v, ok := ft.storage[words[i]]; ok {
			if v, ok := v.(int); ok {
				return []map[string]any{
					ft.data[v],
				}
			}
			if l := len(v.([]int)); l < smallest {
				smallest = l
				smallestIndex = i
			}
		}
	}

	// Loop through the indices
	var indices []int = ft.storage[words[smallestIndex]].([]int)
	for i := 0; i < len(indices); i++ {
		for _, value := range ft.data[indices[i]] {
			if v, ok := value.(string); !ok {
				continue
			} else if strings.Contains(strings.ToLower(v), sp.Query) {
				result = append(result, ft.data[indices[i]])
			}
		}
	}

	// Return the result
	return result
}
