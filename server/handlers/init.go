package handlers

import (
	"errors"
	"net/http"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Initialize the full text search cache
// This is a handler function that returns a http.HandlerFunc
func FTInit(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			maxWords int
			maxBytes int
			schema   map[string]bool
		)

		// Convert the max words to an integer
		if maxWordsStr := r.URL.Query().Get("maxwords"); len(maxWordsStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid maxWords")))
			return
		} else {
			if maxWordsInt, err := strconv.Atoi(maxWordsStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxWords = maxWordsInt
			}
		}

		// Convert the max size bytes to an integer
		if maxBytesStr := r.URL.Query().Get("maxbytes"); len(maxBytesStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid maxBytes")))
			return
		} else {
			if maxBytesInt, err := strconv.Atoi(maxBytesStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxBytes = maxBytesInt
			}
		}

		// Decode the schema
		if schemaStr := r.URL.Query().Get("schema"); len(schemaStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid schema")))
			return
		} else {
			if err := Utils.Decode(schemaStr, &schema); err != nil {
				w.Write(Utils.Error(err))
				return
			}
		}

		// Initialize the full text cache
		if err := c.FTInit(maxWords, maxBytes, schema); err != nil {
			w.Write(Utils.Error(err))
			return
		}
		w.Write(Utils.Success())
	}
}

// Initialize the full text search cache
// This is a handler function that returns a http.HandlerFunc
func FTInitJson(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			maxWords int
			maxBytes int
			schema   map[string]bool
			json     map[string]map[string]interface{}
		)

		// Convert the max words to an integer
		if maxWordsStr := r.URL.Query().Get("maxwords"); len(maxWordsStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid max words")))
			return
		} else {
			if maxWordsInt, err := strconv.Atoi(maxWordsStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxWords = maxWordsInt
			}
		}

		// Convert the max size bytes to an integer
		if maxBytesStr := r.URL.Query().Get("maxbytes"); len(maxBytesStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid max size bytes")))
			return
		} else {
			if maxBytesInt, err := strconv.Atoi(maxBytesStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxBytes = maxBytesInt
			}
		}

		// Decode the schema
		if schemaStr := r.URL.Query().Get("schema"); len(schemaStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid schema")))
			return
		} else {
			if err := Utils.Decode(schemaStr, &schema); err != nil {
				w.Write(Utils.Error(err))
				return
			}
		}

		// Get the json from the request url
		if jsonStr := r.URL.Query().Get("json"); len(jsonStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid json")))
			return
		} else {
			if err := Utils.Decode(jsonStr, &json); err != nil {
				w.Write(Utils.Error(err))
				return
			}
		}

		// Initialize the full text cache
		if err := c.FTInitWithMap(json, maxWords, maxBytes, schema); err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Return success message
		w.Write(Utils.Success())
	}
}
