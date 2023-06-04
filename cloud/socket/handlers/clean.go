package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Clean is a handler function that returns a fiber context handler function for cleaning the cache.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the cleaning fails.
func Clean(_ *Utils.Params, c *Hermes.Cache) []byte {
	c.Clean()
	return Utils.Success("null")
}

// FTClean is a handler function that returns a fiber context handler function for cleaning the full-text storage.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the cleaning fails.
func FTClean(_ *Utils.Params, c *Hermes.Cache) []byte {
	if err := c.FTClean(); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}
