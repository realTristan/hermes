package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Get is a handler function that returns a fiber context handler function for retrieving a key from the cache.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the value of the key or an error message if the key is not provided or the retrieval fails.
func Get(p *Utils.Params, c *Hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = Utils.GetKeyParam(p); err != nil {
		return Utils.Error("key not provided")
	}

	// Get the value from the cache
	if data, err := json.Marshal(c.Get(key)); err != nil {
		return Utils.Error(err)
	} else {
		return data
	}
}

// GetAll is a handler function that returns a fiber context handler function for retrieving all data from the cache.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all data from the cache or an error message if the retrieval fails.
func GetAll(_ *Utils.Params, c *Hermes.Cache) []byte {
	return nil
}
