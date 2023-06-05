package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Clean is a handler function that returns a fiber context handler function for cleaning the cache.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the cleaning fails.
func Clean(_ *utils.Params, c *hermes.Cache) []byte {
	c.Clean()
	return utils.Success("null")
}

// FTClean is a handler function that returns a fiber context handler function for cleaning the full-text storage.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the cleaning fails.
func FTClean(_ *utils.Params, c *hermes.Cache) []byte {
	if err := c.FTClean(); err != nil {
		return utils.Error(err)
	}
	return utils.Success("null")
}
