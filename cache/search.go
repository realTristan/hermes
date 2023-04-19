package cache

import "strings"

/* SearchWithSpaces function with Mutex Locking
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
 *     result := ft.SearchWithSpaces(query, limit, strict, schema)
 */
func (ft *FullText) SearchWithSpaces(query string, limit int, strict bool, schema map[string]bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchWithSpaces(query, limit, strict, schema)
}

/*
This function searches for a query by splitting the query into separate words and returning the search results.
Parameters:
  - ft (*FullText): A pointer to the FullText struct
  - query (string): The search query
  - limit (int): The limit of search results to return
  - strict (bool): A boolean to indicate whether the search should be strict or not
  - schema (map[string]bool): A map containing the schema to search for

Returns:
  - []map[string]interface{}: An array of maps containing the search results
*/
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

/*
SearchInDataWithKey - function to search data in FullText struct by a given query and key with Mutex locking.

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
	results := ft.SearchInDataWithKey("apple", "description", 10)
*/
func (ft *FullText) SearchInDataWithKey(query string, key string, limit int) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInDataWithKey(query, key, limit)
}

/*
searchInDataWithKey - function to search data in FullText struct by a given query and key.
This function is not thread-safe and should only be called internally by SearchInDataWithKey which provides the necessary Mutex locking.

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
	results := ft.searchInDataWithKey("apple", "description", 10)
*/
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

/*
SearchInData - function to search data in FullText struct by a given query with Mutex locking.

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
	results := ft.SearchInData("apple", 10, schema)
*/
func (ft *FullText) SearchInData(query string, limit int, schema map[string]bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchInData(query, limit, schema)
}

/*
searchInData - function to search data in FullText struct by a given query.

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
	results := ft.searchInData("apple", 10, schema)
*/
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
func (ft *FullText) SearchOne(query string, limit int, strict bool) []map[string]interface{} {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchOne(query, limit, strict)
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
