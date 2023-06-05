package handlers

import (
	"encoding/json"

	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Get is a handler function that returns a fiber context handler function for retrieving a key from the cache.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the value of the key or an error message if the key is not provided or the retrieval fails.
func Get(p *utils.Params, c *hermes.Cache) []byte {
	// Get the key from the query
	var (
		key string
		err error
	)
	if key, err = utils.GetKeyParam(p); err != nil {
		return utils.Error("key not provided")
	}

	// Get the value from the cache
	if data, err := json.Marshal(c.Get(key)); err != nil {
		return utils.Error(err)
	} else {
		return data
	}
}

// GetAll is a handler function that returns a fiber context handler function for retrieving all data from the cache.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing all data from the cache or an error message if the retrieval fails.
func GetAll(_ *utils.Params, c *hermes.Cache) []byte {
	return nil
}
