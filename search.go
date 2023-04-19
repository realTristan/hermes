package hermes

import (
	"strings"
)

// SearchWithSpaces function with lock
func (ft *FullText) SearchWithSpaces(query string, limit int, strict bool, schema map[string]bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchWithSpaces(query, limit, strict, schema)
}

// Search for multiple words
func (ft *FullText) searchWithSpaces(query string, limit int, strict bool, schema map[string]bool) []map[string]string {
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]string{}
	// Get the search result of the first word
	case len(words) == 1:
		return ft.searchOne(words[0], limit, strict)
	}

	// Check if the query is in the cache
	if _, ok := ft.wordCache[words[0]]; !ok {
		return []map[string]string{}
	}

	// Define variables
	var result []map[string]string = []map[string]string{}

	// Loop through the indices
	var indices []int = ft.wordCache[words[0]]
	for i := 0; i < len(indices); i++ {
		for key, value := range ft.data[indices[i]] {
			switch {
			// Check if the key is in the schema
			case !schema[key]:
				continue
			// Check if the value contains the query
			case containsIgnoreCase(value, query):
				result = append(result, ft.data[i])
			}
		}
	}

	// Return the result
	return result
}

// SearchInJsonWithKey function with lock
func (ft *FullText) SearchInJsonWithKey(query string, key string, limit int) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInJsonWithKey(query, key, limit)
}

// SearchInJsonWithKey function
func (ft *FullText) searchInJsonWithKey(query string, key string, limit int) []map[string]string {
	// Define variables
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		if containsIgnoreCase(ft.data[i][key], query) {
			result = append(result, ft.data[i])
		}
	}

	// Return the result
	return result
}

// SearchInJson function with lock
func (ft *FullText) SearchInJson(query string, limit int, schema map[string]bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInJson(query, limit, schema)
}

// searchInJson function
func (ft *FullText) searchInJson(query string, limit int, schema map[string]bool) []map[string]string {
	// Define variables
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		// Iterate over the keys and values for the json data for that index
		for key, value := range ft.data[i] {
			switch {
			case !schema[key]:
				continue
			case containsIgnoreCase(value, query):
				result = append(result, ft.data[i])
			}
		}
	}

	// Return the result
	return result
}

// Search function with lock
func (ft *FullText) SearchOne(query string, limit int, strict bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchOne(query, limit, strict)
}

// Search for a single query
func (ft *FullText) searchOne(query string, limit int, strict bool) []map[string]string {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]string{}
	}

	// Define the result variable
	var result []map[string]string = []map[string]string{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := ft.wordCache[query]; !ok {
			return result
		}

		// Loop through the indices
		var indices []int = ft.wordCache[query]
		for i := 0; i < len(indices); i++ {
			result = append(result, ft.data[indices[i]])
		}

		// Return the result
		return result
	}

	// Define variables
	var indices []int = make([]int, len(ft.data))

	// Loop through the cache keys
	for i := 0; i < len(ft.words); i++ {
		switch {
		case len(result) >= limit:
			return result
		case !contains(ft.words[i], query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(ft.wordCache[ft.words[i]]); j++ {
			var index int = ft.wordCache[ft.words[i]][j]
			if indices[index] == -1 {
				continue
			}

			// Append the index to the result
			result = append(result, ft.data[index])
			indices[index] = -1
		}
	}

	// Return the result
	return result
}
