package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get the cache length
// This is a handler function that returns a fiber context handler function
func Length(_ *Utils.Params, c *Hermes.Cache) []byte {
	return Utils.Success(c.Length())
}
