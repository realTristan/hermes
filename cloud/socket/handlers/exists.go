package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Exists is a handler function that returns a fiber context handler function for checking if a key exists in the cache.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a boolean value indicating whether the key exists in the cache or an error message if the key is not provided.
func Exists(p *utils.Params, c *hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = utils.GetKeyParam(p); err != nil {
		return utils.Error("key not provided")
	}

	// Return whether the key exists
	return utils.Success(c.Exists(key))
}
