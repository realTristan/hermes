package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get Values from cache
// This is a handler function that returns a fiber context handler function
func Values(_ *Utils.Params, c *Hermes.Cache) []byte {
	if values, err := json.Marshal(c.Values()); err != nil {
		return Utils.Error(err)
	} else {
		return values
	}
}
