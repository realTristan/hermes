package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Search is a handler function that returns a fiber context handler function for searching the cache for a query.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func Search(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		strict bool
		query  string
		limit  int
		err    error
	)

	// Get the query from the params
	if query, err = Utils.GetQueryParam(p); err != nil {
		return Utils.Error("query not provided")
	}

	// Get the limit from the params
	if err := Utils.GetLimitParam(p, &limit); err != nil {
		return Utils.Error(err)
	}

	// Get the strict from the params
	if err := Utils.GetStrictParam(p, &strict); err != nil {
		return Utils.Error(err)
	}

	// Search for the query
	if res, err := c.Search(Hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Strict: strict,
	}); err != nil {
		return Utils.Error(err)
	} else if data, err := json.Marshal(res); err != nil {
		return Utils.Error(err)
	} else {
		return data
	}
}

// SearchOneWord is a handler function that returns a fiber context handler function for searching the cache for a single word query.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchOneWord(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		strict bool
		query  string
		err    error
		limit  int
	)

	// Get the query from the params
	if query, err = Utils.GetQueryParam(p); err != nil {
		return Utils.Error("invalid query")
	}

	// Get the limit from the params
	if err := Utils.GetLimitParam(p, &limit); err != nil {
		return Utils.Error(err)
	}

	// Get the strict from the params
	if err := Utils.GetStrictParam(p, &strict); err != nil {
		return Utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchOneWord(Hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Strict: strict,
	}); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}

// SearchValues is a handler function that returns a fiber context handler function for searching the cache for a query in values.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchValues(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		query  string
		limit  int
		err    error
		schema map[string]bool
	)

	// Get the query from the params
	if query, err = Utils.GetQueryParam(p); err != nil {
		return Utils.Error("invalid query")
	}

	// Get the limit from the params
	if err := Utils.GetLimitParam(p, &limit); err != nil {
		return Utils.Error(err)
	}

	// Get the schema from the params
	if err := Utils.GetSchemaParam(p, &schema); err != nil {
		return Utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchValues(Hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Schema: schema,
	}); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}

// SearchWithKey is a handler function that returns a fiber context handler function for searching the cache for a query with a specific key.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchWithKey(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		key    string
		query  string
		err    error
		limit  int
		schema map[string]bool
	)

	// Get the query from the params
	if query, err = Utils.GetQueryParam(p); err != nil {
		return Utils.Error("invalid query")
	}

	// Get the key from the params
	if key, err = Utils.GetKeyParam(p); err != nil {
		return Utils.Error("invalid key")
	}

	// Get the limit from the params
	if err := Utils.GetLimitParam(p, &limit); err != nil {
		return Utils.Error(err)
	}

	// Get the schema from the params
	if err := Utils.GetSchemaParam(p, &schema); err != nil {
		return Utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchWithKey(Hermes.SearchParams{
		Query: query,
		Key:   key,
		Limit: limit,
	}); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}
