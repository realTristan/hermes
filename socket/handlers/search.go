package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Search for something in the cache
// This is a handler function that returns a fiber context handler function
func Search(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		strict bool
		query  string
		limit  int
		err    error
		schema map[string]bool
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

	// Get the schema from the params
	if err := Utils.GetSchemaParam(p, &schema); err != nil {
		return Utils.Error(err)
	}

	// Search for the query
	if res, err := c.Search(query, limit, strict, schema); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}

// Search for one word
// This is a handler function that returns a fiber context handler function
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
	if res, err := c.SearchOneWord(query, limit, strict); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}

// Search in values
// This is a handler function that returns a fiber context handler function
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
	if res, err := c.SearchValues(query, limit, schema); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}

// Search for values
// This is a handler function that returns a fiber context handler function
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
	if res, err := c.SearchWithKey(query, key, limit); err != nil {
		return Utils.Error(err)
	} else {
		if data, err := json.Marshal(res); err != nil {
			return Utils.Error(err)
		} else {
			return data
		}
	}
}
