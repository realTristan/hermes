package hermes

import (
	"strings"
)

// SearchWithSpaces function with lock
func (c *Cache) SearchWithSpaces(query string, limit int, strict bool, excludeKeys []string) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchWithSpaces(query, limit, strict, excludeKeys)
}

// Search for multiple words
func (c *Cache) _SearchWithSpaces(query string, limit int, strict bool, excludeKeys []string) ([]map[string]string, []int) {
	// Split the query into words
	var words []string = strings.Split(strings.TrimSpace(query), " ")

	// If the words array is empty
	if len(words) == 0 {
		return []map[string]string{}, []int{}
	}

	// If there's only one word, return the result
	if len(words) == 1 {
		return c.Search(query, limit, strict)
	}

	// Create an array to store the result
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = []int{}
	)

	// Loop through the words and get the indices that are common
	for i := range words {
		// Search for the query inside the cache
		var queryResult, queryIndices = c.Search(words[i], limit, strict)

		// Iterate over the result
		for j := range queryResult {
			// If the data's already been added
			if _ContainsInt(indices, j) {
				continue
			}

			// Iterate over the keys and values for the json data for that index
			for key, value := range queryResult[j] {
				switch {
				// If the key is in the excludeKeys array
				case _ContainsString(excludeKeys, key):
					continue

				// If the value length is less than the query length
				case len(value) < len(query):
					continue
				}

				// If the data contains the query
				if strings.Contains(value, query) {
					result = append(result, queryResult[j])
					indices = append(indices, queryIndices[j])
				}
			}
		}
	}

	// Return the result
	return result, indices
}

// SearchInJsonWithKey function with lock
func (c *Cache) SearchInJsonWithKey(query string, key string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchInJsonWithKey(query, key, limit, strict)
}

// SearchInJsonWithKey function
func (c *Cache) _SearchInJsonWithKey(query string, key string, limit int, strict bool) ([]map[string]string, []int) {
	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = []int{}
	)

	// Iterate over the query result
	for i := range c.json {
		switch {
		// If the json value length is less than the query length
		case len(c.json[i][key]) < len(query):
			continue
		}

		// If the data contains the query
		if ContainsIgnoreCase(c.json[i][key], query) {
			result = append(result, c.json[i])
			indices = append(indices, i)
		}
	}

	// Return the result
	return result, indices
}

// SearchInJson function with lock
func (c *Cache) SearchInJson(query string, limit int, strict bool, excludeKeys []string) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchInJson(query, limit, strict, excludeKeys)
}

// _SearchInJson function
func (c *Cache) _SearchInJson(query string, limit int, strict bool, excludeKeys []string) ([]map[string]string, []int) {
	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = []int{}
	)

	// Iterate over the query result
	for i := range c.json {
		// Iterate over the keys and values for the json data for that index
		for key, value := range c.json[i] {
			switch {
			// If the key is in the excludeKeys array
			case _ContainsString(excludeKeys, key):
				continue

			// If the value length is less than the query length
			case len(value) < len(query):
				continue
			}

			// If the data contains the query
			if ContainsIgnoreCase(value, query) {
				result = append(result, c.json[i])
				indices = append(indices, i)
			}
		}
	}

	// Return the result
	return result, indices
}

// Search common queries with lock
func (c *Cache) SearchCommon(queries []string, limit int, strict bool) ([]map[string]string, []int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c._SearchCommon(queries, limit, strict)
}

// Search common queries
func (c *Cache) _SearchCommon(queries []string, limit int, strict bool) ([]map[string]string, []int) {
	// If the queries array is empty
	if len(queries) == 0 {
		return []map[string]string{}, []int{}
	}

	// If there's only one word, return the result
	if len(queries) == 1 {
		return c._Search(queries[0], limit, strict)
	}

	// Store results
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = []int{}
	)

	// Loop through the queries and get the indices that are common
	for i := 1; i < len(queries); i++ {
		// Search for the word
		var _, queryIndices = c._Search(queries[i], limit, strict)

		// If the allIndices array is empty, set it to the indices
		if len(indices) == 0 {
			indices = queryIndices
			continue
		}

		// Loop through the indices and remove the ones that are not common
		for j := 0; j < len(indices); j++ {
			// Check if the index is in the indices array
			if !_ContainsInt(queryIndices, indices[j]) {
				// Remove the index from the firstIndices array
				indices = append(indices[:j], indices[j+1:]...)
			}
		}
	}

	// Loop through the indices
	for i := range indices {
		result = append(result, c.json[indices[i]])
	}

	// Return the result
	return result, indices
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

	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = []int{}
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

		// Check if the key is shorter than the query
		case len(c.keys[i]) < len(query):
			continue

		// If the key doesn't start with the word
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
