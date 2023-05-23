package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Length is a handler function that returns a fiber context handler function for retrieving the length of the cache.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the length of the cache.
func Length(_ *Utils.Params, c *Hermes.Cache) []byte {
	return Utils.Success(c.Length())
}
