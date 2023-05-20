package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get Keys from cache
// This is a handler function that returns a fiber context handler function
func Keys(_ *Utils.Params, c *Hermes.Cache) []byte {
	if keys, err := json.Marshal(c.Keys()); err != nil {
		return Utils.Error(err)
	} else {
		return keys
	}
}
