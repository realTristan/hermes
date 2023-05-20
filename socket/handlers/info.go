package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get cache info in the form of a string
// This is a handler function that returns a fiber context handler function
func Info(_ *Utils.Params, c *Hermes.Cache) []byte {
	if info, err := c.InfoString(); err != nil {
		return Utils.Error(err)
	} else {
		return []byte(info)
	}
}

// Get cache info for testing in the form of a string
// This is a handler function that returns a fiber context handler function
func InfoForTesting(_ *Utils.Params, c *Hermes.Cache) []byte {
	if info, err := c.InfoStringForTesting(); err != nil {
		return Utils.Error(err)
	} else {
		return []byte(info)
	}
}
