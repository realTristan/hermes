package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Clean the regular cache
// This is a handler function that returns a fiber context handler function
func Clean(_ *Utils.Params, c *Hermes.Cache) []byte {
	c.Clean()
	return Utils.Success("null")
}

// Clean the full-text storage
// This is a handler function that returns a fiber context handler function
func FTClean(_ *Utils.Params, c *Hermes.Cache) []byte {
	if err := c.FTClean(); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}
