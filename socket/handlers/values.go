package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Values is a handler function that returns a fiber context handler function for retrieving all values from the cache.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all values from the cache or an error message if the retrieval fails.
func Values(_ *Utils.Params, c *Hermes.Cache) []byte {
	if values, err := json.Marshal(c.Values()); err != nil {
		return Utils.Error(err)
	} else {
		return values
	}
}
