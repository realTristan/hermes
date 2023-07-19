package hermes

import (
	"errors"
	"strings"
)

// Search is a method of the Cache struct that searches for a query by splitting the query into separate words and returning the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps containing the search results.
//   - error: An error if the query is invalid or if the smallest words array is not found in the cache.
func (c *Cache) Search(sp SearchParams) ([]map[string]any, error) {
	// If the query is empty, return an error
	if len(sp.Query) == 0 {
		return []map[string]any{}, errors.New("invalid query")
	}

	// If no limit is provided, set it to 10
	if sp.Limit == 0 {
		sp.Limit = 10
	}

	// Lock the mutex
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Check if the FT index is initialized
	if c.ft == nil {
		return []map[string]any{}, errors.New("full-text not initialized")
	}

	// Set the query to lowercase
	sp.Query = strings.ToLower(sp.Query)

	// Search for the query
	return c.search(sp), nil
}

// search is a method of the Cache struct that searches for a query by splitting the query into separate words and returning the search results.
// Parameters:
//   - c (c *Cache): A pointer to the Cache struct
//   - sp (SearchParams): A SearchParams struct containing the search parameters.
//
// Returns:
//   - []map[string]any: A slice of maps containing the search results.
func (c *Cache) search(sp SearchParams) []map[string]any {
	// Split the query into separate words
	var words []string = strings.Split(strings.TrimSpace(sp.Query), " ")
	switch {
	// If the words array is empty
	case len(words) == 0:
		return []map[string]any{}
	// Get the search result of the first word
	case len(words) == 1:
		sp.Query = words[0]
		return c.searchOneWord(sp)
	}

	// Define variables
	var result []map[string]any = []map[string]any{}

	// Variables for storing the smallest words array
	// var smallestData []int = []int{}
	var (
		smallestIndex int = 0
		smallest      int = 0
	)

	// Check if the query is in the cache
	if indices, ok := c.ft.storage[words[0]]; !ok {
		return []map[string]any{}
	} else {
		/*for {
			if v, ok := indices.(string); ok {
				indices = c.ft.storage[v]
			} else {
				break
			}
		}*/
		if temp, ok := indices.(int); ok {
			return []map[string]any{
				c.data[c.ft.indices[temp]],
			}
		}
		// smallestData = indices.([]int)
		smallest = len(indices.([]int))
	}

	// Find the smallest words array
	// Don't include the first or last words from the query
	for i := 1; i < len(words)-1; i++ {
		if indices, ok := c.ft.storage[words[i]]; ok {
			/*for {
				if v, ok := indices.(string); ok {
					indices = c.ft.storage[v]
				} else {
					break
				}
			}*/
			if index, ok := indices.(int); ok {
				return []map[string]any{
					c.data[c.ft.indices[index]],
				}
			}
			/*if l := len(indices.([]int)); l < len(smallestData) {
				smallestData = indices.([]int)
			}*/
			if l := len(indices.([]int)); l < smallest {
				smallest = l
				smallestIndex = i
			}
		}
	}

	// Loop through the indices
	/*for i := 0; i < len(smallestData); i++ {
		for _, value := range c.data[c.ft.indices[smallestData[i]]] {
			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if strings.Contains(strings.ToLower(v), sp.Query) {
					result = append(result, c.data[c.ft.indices[smallestData[i]]])
				}
			}
		}
	}*/
	var keys []int = c.ft.storage[words[smallestIndex]].([]int)
	for i := 0; i < len(keys); i++ {
		for _, value := range c.data[c.ft.indices[keys[i]]] {
			// Check if the value contains the query
			if v, ok := value.(string); ok {
				if strings.Contains(strings.ToLower(v), sp.Query) {
					result = append(result, c.data[c.ft.indices[keys[i]]])
				}
			}
		}
	}

	// Return the result
	return result
}
