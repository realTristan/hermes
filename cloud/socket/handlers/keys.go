package handlers

import (
	"encoding/json"

	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Keys is a handler function that returns a fiber context handler function for retrieving all keys from the cache.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all keys from the cache or an error message if the retrieval fails.
func Keys(_ *utils.Params, c *hermes.Cache) []byte {
	if keys, err := json.Marshal(c.Keys()); err != nil {
		return utils.Error(err)
	} else {
		return keys
	}
}
