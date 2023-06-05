package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Length is a handler function that returns a fiber context handler function for retrieving the length of the cache.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the length of the cache.
func Length(_ *utils.Params, c *hermes.Cache) []byte {
	return utils.Success(c.Length())
}
