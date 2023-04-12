package hermes

import "strings"

// SearchMultiple function with lock
func (c *Cache) SearchMultiple(query string, limit int, strict bool) []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchMultiple(query, limit, strict)
}

// Search for multiple words
func (c *Cache) _SearchMultiple(query string, limit int, strict bool) []map[string]string {
	// Split the query into words
	var words []string = strings.Split(strings.TrimSpace(query), " ")

	// If the words array is empty
	if len(words) == 0 {
		return []map[string]string{}
	}

	// If there's only one word, return the result
	if len(words) == 1 {
		var r, _ = c.Search(query, limit, strict)
		return r
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}
	var alreadyAdded []int = []int{}

	// Loop through the words and get the indices that are common
	for i := range words {
		// Search for the query inside the cache
		var queryResult, indices = c.Search(words[i], limit, strict)

		// Iterate over the result
		for j := range queryResult {
			// If the data's already been added
			if _ContainsInt(alreadyAdded, j) {
				continue
			}

			// Iterate over the keys and values for the json data for that index
			for _, value := range queryResult[j] {
				// If the data contains the query
				if strings.Contains(value, query) {
					result = append(result, queryResult[j])
					alreadyAdded = append(alreadyAdded, indices[j])
				}
			}
		}
	}

	// Return the result
	return result
}

// SearchInJsonWithKey function with lock
func (c *Cache) SearchInJsonWithKey(query string, key string, limit int, strict bool) []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchInJsonWithKey(query, key, limit, strict)
}

// SearchInJsonWithKey function
func (c *Cache) _SearchInJsonWithKey(query string, key string, limit int, strict bool) []map[string]string {
	var queryResult, _ = c._Search(query, limit, strict)

	// Create a variable to store results
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := range queryResult {
		// If the data contains the query
		if strings.Contains(c.json[i][key], query) {
			result = append(result, queryResult[i])
		}
	}

	// Return the result
	return result
}

// SearchInJson function with lock
func (c *Cache) SearchInJson(query string, limit int, strict bool) []map[string]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchInJson(query, limit, strict)
}

// _SearchInJson function
func (c *Cache) _SearchInJson(query string, limit int, strict bool) []map[string]string {
	var queryResult, _ = c._Search(query, limit, strict)

	// Create a variable to store results
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := range queryResult {
		// Iterate over the keys and values for the json data for that index
		for key := range queryResult[i] {
			// If the data contains the query
			if strings.Contains(c.json[i][key], query) {
				result = append(result, queryResult[i])
			}
		}
	}

	// Return the result
	return result
}

// Search function with lock
func (c *Cache) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._Search(query, limit, strict)
}

// Search for a single query
func (c *Cache) _Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]string{}, []int{}
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Create an array to store the indices that have already
	// been added to the result array
	var indices []int = []int{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := c.cache[query]; !ok {
			return result, indices
		}

		// Loop through the indices
		indices = c.cache[query]
		for i := range indices {
			result = append(result, c.json[indices[i]])
		}

		// Return the result
		return result, indices
	}

	// Loop through the cache keys
	for i := range c.keys {
	Switch:
		switch {
		// Check if the limit has been reached
		case len(result) >= limit:
			return result, indices

		// Check if the key is equal to the query
		case c.keys[i] == query:
			break Switch

		// If the key doesn't start with the word
		case !strings.HasPrefix(c.keys[i], query):
			continue

		// Check if the key is shorter than the query
		case len(c.keys[i]) < len(query):
			continue

		// Check if the key contains the query
		case !strings.Contains(c.keys[i], query):
			continue

		// Check if the index is already in the result
		case _ContainsInt(indices, i):
			continue
		}

		// Loop through the indices
		for j := range c.cache[c.keys[i]] {
			var index int = c.cache[c.keys[i]][j]

			// Else, append the index to the result
			result = append(result, c.json[index])
			indices = append(indices, index)
		}
	}

	// Return the result
	return result, indices
}
