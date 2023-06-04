package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Delete is a handler function that returns a fiber context handler function for deleting a key from the cache.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the key is not provided or the deletion fails.
func Delete(p *Utils.Params, c *Hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = Utils.GetKeyParam(p); err != nil {
		return Utils.Error("key not provided")
	}

	// Delete the key from the cache
	c.Delete(key)
	return Utils.Success("null")
}
