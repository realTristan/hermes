package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Check if key exists
// This is a handler function that returns a fiber context handler function
func Exists(p *Utils.Params, c *Hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = Utils.GetKeyParam(p); err != nil {
		return Utils.Error("key not provided")
	}

	// Return whether the key exists
	return Utils.Success(c.Exists(key))
}
