package handlers

import (
	"errors"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Add key from cache to the full-text cache
// This is a handler function that returns a http.HandlerFunc
func FTAdd(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Add the key to the full-text cache
		if err := c.FTAdd(key); err != nil {
			w.Write(Utils.Error(err))
			return
		}
		w.Write(Utils.Success())
	}
}
