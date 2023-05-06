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

// Clean the regular cache
func Clean(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Clean()
		w.Write(Utils.Success())
	}
}

// Delete a key from the cache
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

// Get a key from the cache
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
func GetAll(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// Set a value in the cache
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

		// Get whether or not to store the full text
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

// Get Keys from cache
func Keys(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if keys, err := json.Marshal(c.Keys()); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write(keys)
		}
	}
}

// Get Values from cache
func Values(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if values, err := json.Marshal(c.Values()); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write(values)
		}
	}
}

// Get cache info
func Info(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if info, err := c.Info(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write([]byte(info))
		}
	}
}

// Get the cache length
func Length(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%d", c.Length())))
	}
}

// Check if key exists
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
