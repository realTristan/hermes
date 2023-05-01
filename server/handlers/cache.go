package handlers

import (
	"encoding/json"
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
		var key string = r.URL.Query().Get("key")
		c.Delete(key)
		w.Write(Utils.Success())
	}
}

// Get a key from the cache
func Get(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var key string = r.URL.Query().Get("key")
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
		var (
			key        string = r.URL.Query().Get("key")
			value, err        = Utils.Decode(r.URL.Query().Get("value"))
		)
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Get whether or not to store the full text
		fullText, err := strconv.ParseBool(r.URL.Query().Get("fulltext"))
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Set the value in the cache
		if err := c.Set(key, value, fullText); err != nil {
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
		var key string = r.URL.Query().Get("key")
		w.Write([]byte(fmt.Sprintf("%v", c.Exists(key))))
	}
}
