package cache

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
	if _, ok := ft.wordCache[words[0]]; !ok {
		return []map[string]interface{}{}
	}

	// Loop through the indices
	var keys []string = ft.wordCache[words[0]]
	for i := 0; i < len(keys); i++ {
		for key, value := range ft.data[keys[i]] {
			if !schema[key] {
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if containsIgnoreCase(v, query) {
					result = append(result, ft.data[keys[i]])
				}
			}
		}
	}

	// Return the result
	return result
}

// SearchInDataWithKey function with Mutex Locking
func (ft *FullText) SearchInDataWithKey(query string, key string, limit int) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInDataWithKey(query, key, limit)
}

// SearchInDataWithKey function
func (ft *FullText) searchInDataWithKey(query string, key string, limit int) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range ft.data {
		for _, v := range item {
			if v, ok := v.(string); ok {
				if containsIgnoreCase(v, query) {
					result = append(result, item)
				}
			}
		}
	}

	// Return the result
	return result
}

// SearchInData function with Mutex Locking
func (ft *FullText) SearchInData(query string, limit int, schema map[string]bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInData(query, limit, schema)
}

// searchInData function
func (ft *FullText) searchInData(query string, limit int, schema map[string]bool) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range ft.data {
		// Iterate over the keys and values for the data for that index
		for key, value := range item {
			if !schema[key] {
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if containsIgnoreCase(v, query) {
					result = append(result, item)
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

	// Set the query to lowercase
	query = strings.ToLower(query)

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := ft.wordCache[query]; !ok {
			return result
		}

		// Loop through the indices
		for i := 0; i < len(ft.wordCache[query]); i++ {
			result = append(result, ft.data[ft.wordCache[query][i]])
		}

		// Return the result
		return result
	}

	// Define a map to store the indices that have already been added
	var alreadyAdded map[string]int = map[string]int{}

	// Loop through the cache keys
	for k, v := range ft.wordCache {
		switch {
		case len(result) >= limit:
			return result
		case !contains(k, query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(v); j++ {
			if _, ok := alreadyAdded[v[j]]; ok {
				continue
			}

			// Else, append the index to the result
			result = append(result, ft.data[v[j]])
			alreadyAdded[v[j]] = -1
		}
	}

	// Return the result
	return result
}
