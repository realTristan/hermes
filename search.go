package hermes

import "strings"

// SearchWithSpaces function with lock
func (c *Cache) SearchWithSpaces(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.searchWithSpaces(query, limit, strict, keys)
}

// Search for multiple words
func (c *Cache) searchWithSpaces(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	// Split the query into words
	var words []string = strings.Split(strings.TrimSpace(query), " ")

	// If the words array is empty
	if len(words) == 0 {
		return []map[string]string{}, []int{}
	}

	// If there's only one word, return the result
	if len(words) == 1 {
		return c.Search(words[0], limit, strict)
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Loop through the words and get the indices that are common
	for i := 0; i < len(words); i++ {
		// Search for the query inside the cache
		var queryResult, _ = c.Search(words[i], limit, strict)

		// Iterate over the result
		for j := 0; j < len(queryResult); j++ {
			// Iterate over the keys and values for the json data for that index
			for key, value := range queryResult[j] {
				switch {
				// If the value length is less than the query length
				case len(value) < len(query):
					continue

				// If the key is excluded
				case !keys[key]:
					continue

				// If the value contains the query
				case contains(value, query, len(query)):
					result = append(result, queryResult[j])
				}
			}
		}
	}

	// Return the result
	return result, []int{}
}

// SearchInJsonWithKey function with lock
func (c *Cache) SearchInJsonWithKey(query string, key string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.searchInJsonWithKey(query, key, limit, strict)
}

// SearchInJsonWithKey function
func (c *Cache) searchInJsonWithKey(query string, key string, limit int, strict bool) ([]map[string]string, []int) {
	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = make([]int, len(c.json))
	)

	// Iterate over the query result
	for i := 0; i < len(c.json); i++ {
		switch {
		// If the json value length is less than the query length
		case len(c.json[i][key]) < len(query):
			continue
		}

		// If the data contains the query
		if containsIgnoreCase(c.json[i][key], query) {
			result = append(result, c.json[i])
			indices = append(indices, i)
		}
	}

	// Return the result
	return result, indices
}

// SearchInJson function with lock
func (c *Cache) SearchInJson(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.searchInJson(query, limit, strict, keys)
}

// _SearchInJson function
func (c *Cache) searchInJson(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = make([]int, len(c.json))
	)

	// Iterate over the query result
	for i := 0; i < len(c.json); i++ {
		// Iterate over the keys and values for the json data for that index
		for key, value := range c.json[i] {
			switch {
			// If the key is excluded
			case !keys[key]:
				continue

			// If the value length is less than the query length
			case len(value) < len(query):
				continue
			}

			// If the data contains the query
			if containsIgnoreCase(value, query) {
				result = append(result, c.json[i])
				indices = append(indices, i)
			}
		}
	}

	// Return the result
	return result, indices
}

// Search function with lock
func (c *Cache) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.search(query, limit, strict)
}

// Search for a single query
func (c *Cache) search(query string, limit int, strict bool) ([]map[string]string, []int) {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]string{}, []int{}
	}

	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = make([]int, len(c.json))
	)

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := c.cache[query]; !ok {
			return result, indices
		}

		// Loop through the indices
		indices = c.cache[query]
		for i := 0; i < len(indices); i++ {
			result = append(result, c.json[indices[i]])
		}

		// Return the result
		return result, indices
	}

	// Loop through the cache keys
	var queryLength int = len(query)
	for i := 0; i < len(c.keys); i++ {
	Switch:
		switch {
		// Check if the limit has been reached
		case len(result) >= limit:
			return result, indices

		// Check if the key is equal to the query
		case c.keys[i] == query:
			break Switch

		// Check if the key is shorter than the query
		case len(c.keys[i]) < queryLength:
			continue

		// If the key doesn't start with the word
		case !contains(c.keys[i], query, queryLength):
			continue
		}

		// Loop through the indices
		for j := 0; j < len(c.cache[c.keys[i]]); j++ {
			var index int = c.cache[c.keys[i]][j]
			if indices[index] == -1 {
				continue
			}

			// Else, append the index to the result
			result = append(result, c.json[index])
			indices[index] = -1
		}
	}

	// Return the result
	return result, indices
}
