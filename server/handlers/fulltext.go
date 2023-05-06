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

// Set the full text max bytes
func FTSetMaxBytes(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the value from the query
		var value string
		if value = r.URL.Query().Get("value"); len(value) == 0 {
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
func FTSetMaxWords(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the value from the query
		var value string
		if value = r.URL.Query().Get("value"); len(value) == 0 {
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
		// Get the key from the query
		var key string
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Add the key to the full text cache
		if err := c.FTAdd(key); err != nil {
			w.Write(Utils.Error(err))
			return
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
		w.Write(Utils.Success())
	}
}

// Initialize the full text search cache
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
