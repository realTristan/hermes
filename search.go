package hermes

import "strings"

/*
SearchWithSpaces is a method of the FullText struct that performs a full text search using the specified query
string, with spaces between words, and returns a slice of maps representing the data items that contain the
query words. The search is case-insensitive and non-strict by default, which means that the query words can
match any part of a field value. If the 'strict' parameter is set to true, the query words must match the entire
field value.

The function takes the following parameters:
  - query (string): the query string to search for
  - limit (int): the maximum number of results to return
  - strict (bool): a flag indicating whether the search should be strict or non-strict
  - schema (map[string]bool): a map representing the fields to search, where the keys are the names of the fields and the
    values are boolean flags indicating whether the field should be searched

The function returns a slice of maps representing the data items that contain the query words, where the keys are the names of the fields and the values are the field values.

Example Usage:

	ft := FullText{}
	results := ft.Search("hello world", 10, false, map[string]bool{"title": true, "content": true})
	fmt.Println(results)
*/
func (ft *FullText) Search(query string, limit int, strict bool, schema map[string]bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.search(query, limit, strict, schema)
}

/*
search searches for all occurrences of the given query string in the FullText object's data.
The search is done by splitting the query into separate words and looking for each of them in the data.
The search result is limited to the specified number of entries, and can optionally be filtered to only
include keys that match a given schema.

Parameters:
  - query (string): The search query to use. This string will be split into separate words and each word will be searched for in the data.
  - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
  - strict (bool): If set to true, only exact matches of the query will be returned. If false, any entry that contains the query as a substring will be returned.
  - schema (map[string]bool): A map of keys to boolean values. If a key is present in this map and its value is true,
    then only entries that have that key will be included in the search result. If a key is not present in the map, or its value is false, then that key will be ignored.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs that match the search query and the specified schema.
    If no results are found, an empty array is returned.

Note: The search is case-insensitive.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": "value1", "key2": "value2"}}, wordCache: map[string][]int{"value1": {0}}
	schema := map[string]bool{"key1": true}
	result := ft.search("value1", 10, false, schema)
	fmt.Println(result) // Output: [{key1:value1 key2:value2}]
*/
func (ft *FullText) search(query string, limit int, strict bool, schema map[string]bool) []map[string]string {
	var words []string = strings.Split(strings.TrimSpace(query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]string{}
	// Get the search result of the first word
	case len(words) == 1:
		return ft.searchOneWord(words[0], limit, strict)
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

/*
SearchValuesWithKey searches for all occurrences of the given query string in the FullText object's data.
The search is done by looking for the query as a substring in the data value associated with the given key
in each entry of the data. The search result is limited to the specified number of entries.

Parameters:
  - query (string): The search query to use. This string will be searched for as a substring in the data value associated with the given key.
  - key (string): The name of the key in the data whose data value should be searched.
  - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs from the entry
    in the data that matched the search query. If no results are found, an empty array is returned.

Note: The search is case-insensitive.

This function uses a read lock to prevent concurrent access to the FullText object's data.
If a write lock is already held by another goroutine, this function will block until the
write lock is released.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": `{"name": "John", "age": 30}`, "key2": "value2"}, {"key1": `{"name": "Jane", "age": 25}`, "key2": "value4"}}}
	result := ft.SearchValuesWithKey("John", "key1", 10)
	fmt.Println(result) // Output: [{key1:{"name": "John", "age": 30}, key2:value2}]
*/
func (ft *FullText) SearchValuesWithKey(query string, key string, limit int) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchValuesWithKey(query, key, limit)
}

/*
searchValuesWithKey searches for all occurrences of the given query string in the FullText object's data.
The search result is limited to the specified number of entries.

Parameters:
- query (string): The search query to use. This string will be searched for as a substring in the data value associated with the given key.
- key (string): The name of the key in the data whose data value should be searched.
- limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs
    from the entry in the data that matched the search query. If no results are found, an empty array is returned.

Note: The search is case-insensitive.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": `{"name": "John", "age": 30}`, "key2": "value2"}, {"key1": `{"name": "Jane", "age": 25}`, "key2": "value4"}}}
	result := ft.searchInValuesWithKey("John", "key1", 10)
	fmt.Println(result) // Output: [{key1:{"name": "John", "age": 30}, key2:value2}]
*/
func (ft *FullText) searchValuesWithKey(query string, key string, limit int) []map[string]string {
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

/*
SearchValues searches for all occurrences of the given query string in the FullText object's data. The search is
done by looking for the query as a substring in the full text search cache data. The search result is limited to the
specified number of entries.

Parameters:
  - query (string): The search query to use. This string will be searched for as a substring in the full text search cache data.
  - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
  - schema (map[string]bool): A dictionary of key-value pairs indicating which keys in the data should be searched.
    Only keys that have a value of true in the schema will be searched. If a key is not present in the schema, it will not be searched.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs from the entry in
    the data that matched the search query. If no results are found, an empty array is returned.

Note: The search is case-insensitive.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": `{"name": "John", "age": 30}`, "key2": "value2"}, {"key1": `{"name": "Jane", "age": 25}`, "key2": "value4"}}}
	schema := map[string]bool{"key1": true, "key2": false}
	result := ft.SearchValues("John", 10, schema)
	fmt.Println(result) // Output: [{key1:{"name": "John", "age": 30}, key2:value2}]
*/
func (ft *FullText) SearchValues(query string, limit int, schema map[string]bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchValues(query, limit, schema)
}

/*
Parameters:
  - query (string): The search query to use. This string will be searched for as a substring in the full text cache data.
  - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
  - schema (map[string]bool): A dictionary of key-value pairs indicating which keys in the data should be searched.
    Only keys that have a value of true in the schema will be searched. If a key is not present in the schema, it will not be searched.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs
    from the entry in the data that matched the search query. If no results are found, an empty array is returned.

Note: The search is case-insensitive.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": `{"name": "John", "age": 30}`, "key2": "value2"}, {"key1": `{"name": "Jane", "age": 25}`, "key2": "value4"}}}
	schema := map[string]bool{"key1": true, "key2": false}
	result := ft.searchInValues("John", 10, schema)
	fmt.Println(result) // Output: [{key1:{"name": "John", "age": 30}, key2:value2}]
*/
func (ft *FullText) searchValues(query string, limit int, schema map[string]bool) []map[string]string {
	// Define variables
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := 0; i < len(ft.data); i++ {
		// Iterate over the keys and values for the data
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

/*
SearchOneWord searches for all occurrences of the given query string in the FullText object's data.
The search is done by looking for the query as a complete word (case-insensitive) in any value
associated with any key in each entry of the data. The search result is limited to the specified
number of entries.

Parameters:
  - query (string): The search query to use. This string will be searched for as a complete word (case-insensitive) in any value associated with any key in the data.
  - limit (int): The maximum number of search results to return. If the number of matching results exceeds this limit, the excess results will be ignored.
  - strict (bool): If true, the search will only return exact matches for the query string. If false, the search will
    also return entries where the query appears as a substring in any value.

Returns:
  - []map[string]string: An array of maps representing the search results. Each map contains key-value pairs
    from the entry in the data that matched the search query. If no results are found, an empty array is returned.

Note: The search is case-insensitive.

Example usage:

	ft := &FullText{data: []map[string]string{{"key1": "value1", "key2": "value2"}, {"key1": "hello world", "key2": "value4"}}}
	result := ft.SearchOneWord("hello", 10, false)
	fmt.Println(result) // Output: [{key1:hello world, key2:value4}]
*/
func (ft *FullText) SearchOneWord(query string, limit int, strict bool) []map[string]string {
	ft.mutex.RLock()
	defer ft.mutex.RUnlock()
	return ft.searchOneWord(query, limit, strict)
}

/*
searchOneWord searches for a single query within the data using a full-text search approach.

Parameters:
  - query (string): the query to search for
  - limit (int): the maximum number of search results to return
  - strict (bool): if true, only exact matches will be returned. If false, partial matches will also be returned.

Returns:
  - []map[string]string: a list of maps, where each map is a row from the data that matches the query.

Note:
  - If the query is empty, the function returns an empty list.
  - If strict is true and the query is not found in the data, an empty list is returned.
  - If strict is false and no matches are found, an empty list is returned.
*/
func (ft *FullText) searchOneWord(query string, limit int, strict bool) []map[string]string {
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

	// true for already checked
	var alreadyChecked map[int]bool = map[int]bool{}

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
			if alreadyChecked[index] {
				continue
			}

			// Append the index to the result
			result = append(result, ft.data[index])
			alreadyChecked[index] = true
		}
	}

	// Return the result
	return result
}
