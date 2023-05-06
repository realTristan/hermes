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
			maxWords     int
			maxSizeBytes int
			schema       map[string]bool
		)

		// Convert the max words to an integer
		if maxWordsStr := r.URL.Query().Get("maxWords"); len(maxWordsStr) == 0 {
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
		if maxSizeBytesStr := r.URL.Query().Get("maxSizeBytes"); len(maxSizeBytesStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid maxSizeBytes")))
			return
		} else {
			if maxSizeBytesInt, err := strconv.Atoi(maxSizeBytesStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxSizeBytes = maxSizeBytesInt
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
		if err := c.FTInit(maxWords, maxSizeBytes, schema); err != nil {
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
			maxWords     int
			maxSizeBytes int
			schema       map[string]bool
			json         map[string]map[string]interface{}
		)

		// Convert the max words to an integer
		if maxWordsStr := r.URL.Query().Get("maxWords"); len(maxWordsStr) == 0 {
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
		if maxSizeBytesStr := r.URL.Query().Get("maxSizeBytes"); len(maxSizeBytesStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid max size bytes")))
			return
		} else {
			if maxSizeBytesInt, err := strconv.Atoi(maxSizeBytesStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				maxSizeBytes = maxSizeBytesInt
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
		if err := c.FTInitWithMap(json, maxWords, maxSizeBytes, schema); err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Return success message
		w.Write(Utils.Success())
	}
}
