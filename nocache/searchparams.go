package nocache

// SearchParams is a struct that contains the search parameters for the Cache search methods.
type SearchParams struct {
	// The search query
	Query string
	// The limit of search results to return
	Limit int
	// A boolean to indicate whether the search should be strict or not
	Strict bool
	// A map containing the schema to search for
	Schema map[string]bool
	// Key to search in
	Key string
}
