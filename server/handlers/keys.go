package handlers

import (
	"encoding/json"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Get Keys from cache
// This is a handler function that returns a http.HandlerFunc
func Keys(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if keys, err := json.Marshal(c.Keys()); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write(keys)
		}
	}
}
