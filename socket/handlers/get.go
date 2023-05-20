package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get a key from the cache
// This is a handler function that returns a fiber context handler function
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

// Get all the data from the cache
// This is a handler function that returns a fiber context handler function
func GetAll(_ *Utils.Params, c *Hermes.Cache) []byte {
	return nil
}
