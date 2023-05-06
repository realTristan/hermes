package handlers

import (
	"errors"
	"fmt"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Check if key exists
// This is a handler function that returns a http.HandlerFunc
func Exists(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Return whether the key exists
		w.Write([]byte(fmt.Sprintf("%v", c.Exists(key))))
	}
}
