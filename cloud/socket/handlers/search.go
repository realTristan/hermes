package handlers

import (
	"encoding/json"

	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Search is a handler function that returns a fiber context handler function for searching the cache for a query.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func Search(p *utils.Params, c *hermes.Cache) []byte {
	var (
		strict bool
		query  string
		limit  int
		err    error
	)

	// Get the query from the params
	if query, err = utils.GetQueryParam(p); err != nil {
		return utils.Error("query not provided")
	}

	// Get the limit from the params
	if err := utils.GetLimitParam(p, &limit); err != nil {
		return utils.Error(err)
	}

	// Get the strict from the params
	if err := utils.GetStrictParam(p, &strict); err != nil {
		return utils.Error(err)
	}

	// Search for the query
	if res, err := c.Search(hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Strict: strict,
	}); err != nil {
		return utils.Error(err)
	} else if data, err := json.Marshal(res); err != nil {
		return utils.Error(err)
	} else {
		return data
	}
}

// SearchOneWord is a handler function that returns a fiber context handler function for searching the cache for a single word query.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchOneWord(p *utils.Params, c *hermes.Cache) []byte {
	var (
		strict bool
		query  string
		err    error
		limit  int
	)

	// Get the query from the params
	if query, err = utils.GetQueryParam(p); err != nil {
		return utils.Error("invalid query")
	}

	// Get the limit from the params
	if err := utils.GetLimitParam(p, &limit); err != nil {
		return utils.Error(err)
	}

	// Get the strict from the params
	if err := utils.GetStrictParam(p, &strict); err != nil {
		return utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchOneWord(hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Strict: strict,
	}); err != nil {
		return utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return utils.Error(err)
		} else {
			return data
		}
	}
}

// SearchValues is a handler function that returns a fiber context handler function for searching the cache for a query in values.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchValues(p *utils.Params, c *hermes.Cache) []byte {
	var (
		query  string
		limit  int
		err    error
		schema map[string]bool
	)

	// Get the query from the params
	if query, err = utils.GetQueryParam(p); err != nil {
		return utils.Error("invalid query")
	}

	// Get the limit from the params
	if err := utils.GetLimitParam(p, &limit); err != nil {
		return utils.Error(err)
	}

	// Get the schema from the params
	if err := utils.GetSchemaParam(p, &schema); err != nil {
		return utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchValues(hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Schema: schema,
	}); err != nil {
		return utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return utils.Error(err)
		} else {
			return data
		}
	}
}

// SearchWithKey is a handler function that returns a fiber context handler function for searching the cache for a query with a specific key.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the search results or an error message if the search fails.
func SearchWithKey(p *utils.Params, c *hermes.Cache) []byte {
	var (
		key    string
		query  string
		err    error
		limit  int
		schema map[string]bool
	)

	// Get the query from the params
	if query, err = utils.GetQueryParam(p); err != nil {
		return utils.Error("invalid query")
	}

	// Get the key from the params
	if key, err = utils.GetKeyParam(p); err != nil {
		return utils.Error("invalid key")
	}

	// Get the limit from the params
	if err := utils.GetLimitParam(p, &limit); err != nil {
		return utils.Error(err)
	}

	// Get the schema from the params
	if err := utils.GetSchemaParam(p, &schema); err != nil {
		return utils.Error(err)
	}

	// Search for the query
	if res, err := c.SearchWithKey(hermes.SearchParams{
		Query: query,
		Key:   key,
		Limit: limit,
	}); err != nil {
		return utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return utils.Error(err)
		} else {
			return data
		}
	}
}
