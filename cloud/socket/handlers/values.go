package handlers

import (
	"encoding/json"

	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Values is a handler function that returns a fiber context handler function for retrieving all values from the cache.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all values from the cache or an error message if the retrieval fails.
func Values(_ *utils.Params, c *hermes.Cache) []byte {
	if values, err := json.Marshal(c.Values()); err != nil {
		return utils.Error(err)
	} else {
		return values
	}
}
