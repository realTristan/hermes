package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Keys is a handler function that returns a fiber context handler function for retrieving all keys from the cache.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all keys from the cache or an error message if the retrieval fails.
func Keys(_ *Utils.Params, c *Hermes.Cache) []byte {
	if keys, err := json.Marshal(c.Keys()); err != nil {
		return Utils.Error(err)
	} else {
		return keys
	}
}
