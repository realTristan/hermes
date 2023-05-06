package handlers

import (
	"encoding/json"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Get Values from cache
// This is a handler function that returns a http.HandlerFunc
func Values(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if values, err := json.Marshal(c.Values()); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write(values)
		}
	}
}
