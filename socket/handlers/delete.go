package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Delete a key from the cache
// This is a handler function that returns a fiber context handler function
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
