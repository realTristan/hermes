package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Get a key from the cache
// This is a handler function that returns a http.HandlerFunc
func Get(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Get the value from the cache
		if data, err := json.Marshal(c.Get(key)); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write(data)
		}
	}
}

// Get all the data from the cache
// This is a handler function that returns a http.HandlerFunc
func GetAll(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
