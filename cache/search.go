package cache

import (
	"errors"
	"strings"
)

/* Search function with Mutex Locking
 *
 * This function is a method of the FullText struct and allows the user to search for a query string with spaces,
 * while also enforcing mutex locking for concurrency safety. The function returns a slice of maps, where each map
 * represents a document that matches the search query.
 *
 * Parameters:
 * - query (string): the search query string with spaces
 * - limit (int): the maximum number of documents to return
 * - strict (bool): a boolean flag indicating whether the search should be performed strictly (i.e., exact match) or not
 * - schema (map[string]bool): a dictionary mapping field names to boolean values that indicate whether each field
 *   should be included in the search (true) or not (false)
 *
 * Returns:
 * - result ([]map[string]interface{}): a slice of maps, where each map represents a document that matches the search query
 *
 * Example usage:
 *
 *     ft := &FullText{...}
 *     query := "open source programming"
 *     limit := 10
 *     strict := false
 *     schema := map[string]bool{"title": true, "content": true, "date": false}
 *     result := ft.Search(query, limit, strict, schema)
 */
func (c *Cache) Search(query string, limit int, strict bool, schema map[string]bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the FT index is initialized
	if !c.ft.isInitialized() {
		return []map[string]interface{}{}, errors.New("full text is not initialized")
	}

	// Search for the query
	return c.search(query, limit, strict, schema)
}

/*
This function searches for a query by splitting the query into separate words and returning the search results.
Parameters:
  - c (c *Cache): A pointer to the Cache struct
  - query (string): The search query
  - limit (int): The limit of search results to return
  - strict (bool): A boolean to indicate whether the search should be strict or not
  - schema (map[string]bool): A map containing the schema to search for

Returns:
  - []map[string]interface{}: An array of maps containing the search results
*/
func (c *Cache) search(query string, limit int, strict bool, schema map[string]bool) ([]map[string]interface{}, error) {
	// Split the query into separate words
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	// Get the search result of the first word
	case len(words) == 1:
		return c.searchOneWord(words[0], limit, strict), nil
	}

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Check if the query is in the cache
	if _, ok := c.ft.wordCache[words[0]]; !ok {
		return []map[string]interface{}{}, errors.New("invalid query: " + query)
	}

	// Loop through the indices
	var keys []string = c.ft.wordCache[words[0]]
	for i := 0; i < len(keys); i++ {
		for key, value := range c.data[keys[i]] {
			if !schema[key] {
				continue
			}

			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if containsIgnoreCase(v, query) {
					result = append(result, c.data[keys[i]])
				}
			}
		}
	}

	// Return the result
	return result, nil
}

/*
SearchValuesWithKey - function to search data in FullText struct by a given query and key with Mutex locking.
Parameters:
  - query: string, the search query to match against data in the FullText struct.
  - key: string, the key to search data on.
  - limit: int, the maximum number of results to be returned.

Returns:
  - []map[string]interface{}, a slice of maps where each map represents a data record that matches the given query and key.
    The keys of the map correspond to the column names of the data.

Mutex locking is used to ensure that concurrent access to the FullText struct is safe.
Example usage:

	Assume we have a FullText struct instance named ft, containing data with columns "id", "name", and "description"
	We can search for all records containing the word "apple" in the "description" column with a limit of 10 as follows:
	results := ft.SearchValuesWithKey("apple", "description", 10)
*/
func (c *Cache) SearchValuesWithKey(query string, key string, limit int) ([]map[string]interface{}, error) {
	switch {
	case len(key) == 0:
		return []map[string]interface{}{}, errors.New("invalid key")
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Search the data
	return c.searchValuesWithKey(query, key, limit), nil
}

/*
searchValuesWithKey - function to search data in FullText struct by a given query and key.
This function is not thread-safe and should only be called internally by SearchValuesWithKey which provides the necessary Mutex locking.
Parameters:
  - query: string, the search query to match against data in the FullText struct.
  - key: string, the key to search data on.
  - limit: int, the maximum number of results to be returned.

Returns:
  - []map[string]interface{}, a slice of maps where each map represents a data record that matches the given query and key.
    The keys of the map correspond to the column names of the data.

Example usage:

	Assume we have a FullText struct instance named ft, containing data with columns "id", "name", and "description"
	We can search for all records containing the word "apple" in the "description" column with a limit of 10 as follows:
	results := ft.searchValuesWithKey("apple", "description", 10)
*/
func (c *Cache) searchValuesWithKey(query string, key string, limit int) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range c.data {
		for _, v := range item {
			if len(result) >= limit {
				return result
			}

			// Check if the value contains the query
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

/*
SearchValues - function to search data in FullText struct by a given query with Mutex locking.
Parameters:
  - query: string, the search query to match against data in the FullText struct.
  - limit: int, the maximum number of results to be returned.
  - schema: map[string]bool, a map of column names in the FullText struct's data that should be searched. The boolean value indicates whether the column should be searched or not.

Returns:
  - []map[string]interface{}, a slice of maps where each map represents a data record that matches the given query.
    The keys of the map correspond to the column names of the data that were searched and returned in the result.

Mutex locking is used to ensure that concurrent access to the FullText struct is safe.
Example usage:

	Assume we have a FullText struct instance named ft, containing data with columns "id", "name", and "description"
	We can search for all records containing the word "apple" in the "description" and "name" columns with a limit of 10 as follows:
	schema := map[string]bool{"description": true, "name": true}
	results := ft.SearchValues("apple", 10, schema)
*/
func (c *Cache) SearchValues(query string, limit int, schema map[string]bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Search the data
	return c.searchValues(query, limit, schema), nil
}

/*
searchValues - function to search data in FullText struct by a given query.
Parameters:
  - query: string, the search query to match against data in the FullText struct.
  - limit: int, the maximum number of results to be returned.
  - schema: map[string]bool, a map of column names in the FullText struct's data that should be searched. The boolean value indicates whether the column should be searched or not.

Returns:
  - []map[string]interface{}, a slice of maps where each map represents a data record that matches the given query.
    The keys of the map correspond to the column names of the data that were searched and returned in the result.

Example usage:

	Assume we have a FullText struct instance named ft, containing data with columns "id", "name", and "description"
	We can search for all records containing the word "apple" in the "description" and "name" columns with a limit of 10 as follows:
	schema := map[string]bool{"description": true, "name": true}
	results := ft.searchValues("apple", 10, schema)
*/
func (c *Cache) searchValues(query string, limit int, schema map[string]bool) []map[string]interface{} {
	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// Iterate over the query result
	for _, item := range c.data {
		// Iterate over the keys and values for the data for that index
		for key, value := range item {
			switch {
			case len(result) >= limit:
				return result
			case !schema[key]:
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

/*
SearchOne function searches for a query string within the FullText object and returns the results in a list of maps. This function uses mutex locking to ensure thread-safety.
@Parameters:

	query (string): The query string to search for.
	limit (int): The maximum number of results to return. If set to 0, there is no limit to the number of results returned.
	strict (bool): Determines whether the search should be strict or not. If set to true, only exact matches will be returned.

@Returns:

	A list of maps, where each map represents a result. Each map contains key-value pairs, where the key is a field name and the value is the corresponding field value.

Example Usage:

	results := ft.SearchOne("example", 0, false)
	for _, result := range results {
		fmt.Printf("Result: %v\n", result)
	}
*/
func (c *Cache) SearchOneWord(query string, limit int, strict bool) ([]map[string]interface{}, error) {
	switch {
	case len(query) == 0:
		return []map[string]interface{}{}, errors.New("invalid query")
	case limit < 1:
		return []map[string]interface{}{}, errors.New("invalid limit")
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the full text is initialized
	if !c.ft.isInitialized() {
		return []map[string]interface{}{}, errors.New("full text is not initialized")
	}

	// Search the data
	return c.searchOneWord(query, limit, strict), nil
}

/*
searchOne function is an internal function used by the FullText object to search for a single word. This function returns the results in a list of maps.
Parameters:

	query (string): The query string to search for.
	limit (int): The maximum number of results to return. If set to 0, there is no limit to the number of results returned.
	strict (bool): Determines whether the search should be strict or not. If set to true, only exact matches will be returned.

Returns:

	A list of maps, where each map represents a result. Each map contains key-value pairs, where the key is a field name and the value is the corresponding field value.

Example Usage:

	This is an internal function and is not intended to be called directly by the user.
*/
func (c *Cache) searchOneWord(query string, limit int, strict bool) []map[string]interface{} {
	// Set the query to lowercase
	query = strings.ToLower(query)

	// Define variables
	var result []map[string]interface{} = []map[string]interface{}{}

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := c.ft.wordCache[query]; !ok {
			return result
		}

		// Loop through the indices
		for i := 0; i < len(c.ft.wordCache[query]); i++ {
			if len(result) >= limit {
				return result
			}
			result = append(result, c.data[c.ft.wordCache[query][i]])
		}

		// Return the result
		return result
	}

	// Define a map to store the indices that have already been added
	var alreadyAdded map[string]int = map[string]int{}

	// Loop through the cache keys
	for k, v := range c.ft.wordCache {
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
			result = append(result, c.data[v[j]])
			alreadyAdded[v[j]] = -1
		}
	}

	// Return the result
	return result
}
