package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Set the full text max bytes
func FTSetMaxBytes(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if maxBytes, err := strconv.Atoi(r.URL.Query().Get("value")); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if err := c.FTSetMaxBytes(maxBytes); err != nil {
				w.Write(Utils.Error(err))
			}
			w.Write(Utils.Success())
		}
	}
}

// Set the full text max words
func FTSetMaxWords(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if maxWords, err := strconv.Atoi(r.URL.Query().Get("value")); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if err := c.FTSetMaxWords(maxWords); err != nil {
				w.Write(Utils.Error(err))
			}
			w.Write(Utils.Success())
		}
	}
}

// Get the full text word cache
func FTWordCache(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cache, err := c.FTWordCache(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if data, err := json.Marshal(cache); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}

// Get the full text word cache size
func FTWordCacheSize(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if size, err := c.FTWordCacheSize(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write([]byte(fmt.Sprintf("%d", size)))
		}
	}
}

// Add key from cache to the full text cache
func FTAdd(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var key string = r.URL.Query().Get("key")
		if err := c.FTAdd(key); err != nil {
			w.Write(Utils.Error(err))
		}
		w.Write(Utils.Success())
	}
}

// Check if full text is initialized
func FTIsInitialized(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v", c.FTIsInitialized())))
	}
}

// Clean the full text cache
func FTClean(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := c.FTClean(); err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Return success message
		w.Write(Utils.Success())
	}
}

// Initialize the full text search cache
func FTInit(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// all from url params. base64 encode the schema on request
		var (
			maxWords     int             = 0
			maxSizeBytes int             = 0
			schema       map[string]bool = map[string]bool{}
		)

		// Initialize the ft cache
		if err := c.FTInit(maxWords, maxSizeBytes, schema); err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Return success message
		w.Write(Utils.Success())
	}
}

// Initialize the full text search cache
func FTInitJson(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// all from url params. base64 encode the schema on request
		var (
			maxWords     int             = 0
			maxSizeBytes int             = 0
			schema       map[string]bool = map[string]bool{}

			// from the headers
			json map[string]map[string]interface{} = nil
		)

		// Initialize the ft cache
		if err := c.FTInitWithMap(json, maxWords, maxSizeBytes, schema); err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Return success message
		w.Write(Utils.Success())
	}
}
