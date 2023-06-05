package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Set is a handler function that returns a fiber context handler function for setting a value in the cache.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the set operation fails.
func Set(p *utils.Params, c *hermes.Cache) []byte {
	var (
		key   string
		err   error
		value map[string]any
	)

	// Get the key from the query
	if key, err = utils.GetKeyParam(p); err != nil {
		return utils.Error("invalid key")
	}

	// Get the value from the query
	if err := utils.GetValueParam(p, &value); err != nil {
		return utils.Error(err)
	}

	// Set the value in the cache
	if err := c.Set(key, value); err != nil {
		return utils.Error(err)
	}
	return utils.Success("null")
}
