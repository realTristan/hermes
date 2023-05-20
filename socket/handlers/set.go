package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Set a value in the cache
// This is a handler function that returns a fiber context handler function
func Set(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		key   string
		err   error
		value map[string]interface{}
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
