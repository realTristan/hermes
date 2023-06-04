package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Set is a handler function that returns a fiber context handler function for setting a value in the cache.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the set operation fails.
func Set(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		key   string
		err   error
		value map[string]any
	)

	// Get the key from the query
	if key, err = Utils.GetKeyParam(p); err != nil {
		return Utils.Error("invalid key")
	}

	// Get the value from the query
	if err := Utils.GetValueParam(p, &value); err != nil {
		return Utils.Error(err)
	}

	// Set the value in the cache
	if err := c.Set(key, value); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}
