package hermes

import "strings"

// SearchWithSpaces function with Mutex Locking
func (ft *FullText) SearchWithSpaces(query string, limit int, strict bool, schema map[string]bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchWithSpaces(query, limit, strict, schema)
}

// Search with words slice
func (ft *FullText) searchWithSpaces(query string, limit int, strict bool, schema map[string]bool) []map[string]interface{} {
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]interface{}{}
	// Get the search result of the first word
	case len(words) == 1:
		return ft.searchOne(words[0], limit, strict)
	}

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Check if the query is in the cache
	if _, ok := ft.cache[words[0]]; !ok {
		return []map[string]interface{}{}
	}

	// Loop through the indices
	var indices []int = ft.cache[words[0]]
	for i := 0; i < len(indices); i++ {
		for key, value := range ft.data[indices[i]] {
			if v, ok := value.(string); ok {
				if !schema[key] {
					continue
				}
				if containsIgnoreCase(v, query) {
					result = append(result, ft.data[indices[i]])
				}
			}
		}
	}

	// Return the result
	return result
}

// SearchInJsonWithKey function with Mutex Locking
func (ft *FullText) SearchInJsonWithKey(query string, key string, limit int) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInJsonWithKey(query, key, limit)
}

// SearchInJsonWithKey function
func (ft *FullText) searchInJsonWithKey(query string, key string, limit int) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		var value interface{} = ft.data[i][key]
		if v, ok := value.(string); ok {
			if containsIgnoreCase(v, query) {
				result = append(result, ft.data[i])
			}
		}
	}

	// Return the result
	return result
}

// SearchInJson function with Mutex Locking
func (ft *FullText) SearchInJson(query string, limit int, schema map[string]bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInJson(query, limit, schema)
}

// searchInJson function
func (ft *FullText) searchInJson(query string, limit int, schema map[string]bool) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		// Iterate over the keys and values for the json data for that index
		for key, value := range ft.data[i] {
			if v, ok := value.(string); ok {
				switch {
				case !schema[key]:
					continue
				case containsIgnoreCase(v, query):
					result = append(result, ft.data[i])
				}
			}
		}
	}

	// Return the result
	return result
}

// Search function with Mutex Locking
func (ft *FullText) SearchOne(query string, limit int, strict bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchOne(query, limit, strict)
}

// Search for a single word
func (ft *FullText) searchOne(query string, limit int, strict bool) []map[string]interface{} {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]interface{}{}
	}

	// Define variables
	var (
		result  []map[string]interface{} = []map[string]interface{}{}
		indices []int                    = make([]int, len(ft.data))
	)

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := ft.cache[query]; !ok {
			return result
		}

		// Loop through the indices
		indices = ft.cache[query]
		for i := 0; i < len(indices); i++ {
			result = append(result, ft.data[indices[i]])
		}

		// Return the result
		return result
	}

	// Loop through the cache keys
	for i := 0; i < len(ft.words); i++ {
		switch {
		case len(result) >= limit:
			return result
		case !contains(ft.words[i], query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(ft.cache[ft.words[i]]); j++ {
			var index int = ft.cache[ft.words[i]][j]
			if indices[index] == -1 {
				continue
			}

			// Else, append the index to the result
			result = append(result, ft.data[index])
			indices[index] = -1
		}
	}

	// Return the result
	return result
}
