package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Delete is a handler function that returns a fiber context handler function for deleting a key from the cache.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the key is not provided or the deletion fails.
func Delete(p *utils.Params, c *hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = utils.GetKeyParam(p); err != nil {
		return utils.Error("key not provided")
	}

	// Delete the key from the cache
	c.Delete(key)
	return utils.Success("null")
}
