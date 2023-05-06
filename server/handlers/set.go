package handlers

import (
	"errors"
	"net/http"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Set a value in the cache
// This is a handler function that returns a http.HandlerFunc
func Set(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Get the value from the query
		var value map[string]interface{}
		if valueStr := r.URL.Query().Get("value"); len(valueStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid value")))
			return
		} else {
			if err := Utils.Decode(valueStr, &value); err != nil {
				w.Write(Utils.Error(err))
				return
			}
		}

		// Get whether or not to store the full-text
		var ft bool
		if ftStr := r.URL.Query().Get("ft"); len(ftStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid ft")))
			return
		} else {
			if ftBool, err := strconv.ParseBool(ftStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				ft = ftBool
			}
		}

		// Set the value in the cache
		if err := c.Set(key, value, ft); err != nil {
			w.Write(Utils.Error(err))
			return
		}
		w.Write(Utils.Success())
	}
}
