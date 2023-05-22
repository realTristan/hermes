package nocache

import (
	"errors"
	"strings"
)

// Search searches for all occurrences of the given query string in the FullText object's data.
// The search is done by splitting the query into separate words and looking for each of them in the data.
// The search result is limited to the specified number of entries, and can optionally be filtered to only
// include keys that match a given schema.
// Parameters:
//   - query (string): The search query to use. This string will be split into separate words and each word will be searched for in the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If set to true, only exact matches of the query will be returned. If false, any entry that contains the query as a substring will be returned.
//   - schema (map[string]bool): A map of keys to boolean values. If a key is present in this map and its value is true,
//     only entries that have a non-empty value for that key will be returned.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid.
func (ft *FullText) Search(query string, limit int, strict bool, schema map[string]bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Convert the query to lowercase
	query = strings.ToLower(query)

	// Lock the mutex
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()

	// Perform the search
	return ft.search(query, limit, strict, schema), nil
}

// search searches for all occurrences of the given query string in the FullText object's data.
// The search is done by splitting the query into separate words and looking for each of them in the data.
// The search result is limited to the specified number of entries, and can optionally be filtered to only
// include keys that match a given schema.
// Parameters:
//   - query (string): The search query to use. This string will be split into separate words and each word will be searched for in the data.
//   - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
//   - strict (bool): If set to true, only exact matches of the query will be returned. If false, any entry that contains the query as a substring will be returned.
//   - schema (map[string]bool): A map of keys to boolean values. If a key is present in this map and its value is true,
//     only entries that have a non-empty value for that key will be returned.
//
// Returns:
//   - []map[string]interface{}: A slice of maps where each map represents a data record that matches the given query.
//     The keys of the map correspond to the column names of the data that were searched and returned in the result.
//   - error: An error if the query or limit is invalid.
func (ft *FullText) search(query string, limit int, strict bool, schema map[string]bool) []map[string]interface{} {
	// Split the query into separate words
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]interface{}{}
	// Get the search result of the first word
	case len(words) == 1:
		return ft.searchOneWord(words[0], limit, strict)
	}

	// Check if the query is in the cache
	if _, ok := ft.storage[words[0]]; !ok {
		return []map[string]interface{}{}
	}

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Variables for storing the smallest words array
	var (
		smallest      int = 0
		smallestIndex int = 0
	)

	// Check if the query is in the cache
	if v, ok := ft.storage[words[0]]; !ok {
		return []map[string]interface{}{}
	} else {
		if t, ok := v.(int); ok {
			return []map[string]interface{}{
				ft.data[t],
			}
		}
		smallest = len(v.([]int))
	}

	// Find the smallest words array
	// Don't include the first or last words from the query
	for i := 1; i < len(words)-1; i++ {
		if v, ok := ft.storage[words[i]]; ok {
			if t, ok := v.(int); ok {
				return []map[string]interface{}{
					ft.data[t],
				}
			}
			if len(v.([]int)) < smallest {
				smallest = len(v.([]int))
				smallestIndex = i
			}
		}
	}

	// Loop through the indices
	var indices []int = ft.storage[words[smallestIndex]].([]int)
	for i := 0; i < len(indices); i++ {
		for key, value := range ft.data[indices[i]] {
			if v, ok := value.(string); !ok {
				continue
			} else {
				switch {
				// Check if the key is in the schema
				case !schema[key]:
					continue
				// Check if the value contains the query
				case strings.Contains(strings.ToLower(v), query):
					result = append(result, ft.data[indices[i]])
				}
			}
		}
	}

	// Return the result
	return result
}
