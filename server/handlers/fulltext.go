package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Check if full text is initialized
// This is a handler function that returns a http.HandlerFunc
func FTIsInitialized(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v", c.FTIsInitialized())))
	}
}

// Set the full text max bytes
// This is a handler function that returns a http.HandlerFunc
func FTSetMaxBytes(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the value from the query
		var value string
		if value = r.URL.Query().Get("maxbytes"); len(value) == 0 {
			w.Write(Utils.Error(errors.New("invalid value")))
			return
		}

		// Convert the value to an integer
		if maxBytes, err := strconv.Atoi(value); err != nil {
			w.Write(Utils.Error(err))
		} else {
			// Set the max bytes
			if err := c.FTSetMaxBytes(maxBytes); err != nil {
				w.Write(Utils.Error(err))
				return
			}
			w.Write(Utils.Success())
		}
	}
}

// Set the full text max words
// This is a handler function that returns a http.HandlerFunc
func FTSetMaxWords(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the value from the query
		var value string
		if value = r.URL.Query().Get("maxwords"); len(value) == 0 {
			w.Write(Utils.Error(errors.New("invalid value")))
			return
		}

		// Convert the value to an integer
		if maxWords, err := strconv.Atoi(value); err != nil {
			w.Write(Utils.Error(err))
		} else {
			// Set the max words
			if err := c.FTSetMaxWords(maxWords); err != nil {
				w.Write(Utils.Error(err))
				return
			}
			w.Write(Utils.Success())
		}
	}
}

// Get the full text word cache
// This is a handler function that returns a http.HandlerFunc
func FTWordCache(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if data, err := c.FTWordCache(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			// Marshal the data
			if data, err := json.Marshal(data); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}

// Get the full text word cache size
// This is a handler function that returns a http.HandlerFunc
func FTWordCacheSize(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if size, err := c.FTWordCacheSize(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write([]byte(fmt.Sprintf("%d", size)))
		}
	}
}
