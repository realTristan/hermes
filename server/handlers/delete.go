package handlers

import (
	"errors"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Delete a key from the cache
// This is a handler function that returns a http.HandlerFunc
func Delete(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Delete the key from the cache
		c.Delete(key)
		w.Write(Utils.Success())
	}
}
