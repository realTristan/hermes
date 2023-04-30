package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Search for something in the cache
func Search(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			params     url.Values = r.URL.Query()
			query      string     = params.Get("q")
			limit, err            = strconv.Atoi(params.Get("limit"))
		)

		// Check the limit
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Get whether strict mode is enabled/disabled
		strict, err := strconv.ParseBool(params.Get("strict"))
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// base64 decode the schema
		var (
			schema    map[string]bool
			b64schema []byte = []byte(params.Get("schema"))
		)
		err = json.Unmarshal(b64schema, &schema)
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Search for the query
		if res, err := c.Search(query, limit, strict, schema); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}

// Search for one word
func SearchOneWord(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			params     url.Values = r.URL.Query()
			query      string     = params.Get("q")
			limit, err            = strconv.Atoi(params.Get("limit"))
		)

		// Check the limit
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Get whether strict mode is enabled/disabled
		strict, err := strconv.ParseBool(params.Get("strict"))
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Search for the query
		if res, err := c.SearchOneWord(query, limit, strict); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}

// Search in values
func SearchValues(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			params     url.Values = r.URL.Query()
			query      string     = params.Get("q")
			limit, err            = strconv.Atoi(params.Get("limit"))
		)

		// Check the limit
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// base64 decode the schema
		var (
			schema    map[string]bool
			b64schema []byte = []byte(params.Get("schema"))
		)
		err = json.Unmarshal(b64schema, &schema)
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Search for the query
		if res, err := c.SearchValues(query, limit, schema); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}

// Search for values
func SearchValuesWithKey(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			params     url.Values = r.URL.Query()
			query      string     = params.Get("q")
			key        string     = params.Get("key")
			limit, err            = strconv.Atoi(params.Get("limit"))
		)

		// Check the limit
		if err != nil {
			w.Write(Utils.Error(err))
			return
		}

		// Search for the query
		if res, err := c.SearchValuesWithKey(query, key, limit); err != nil {
			w.Write(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				w.Write(Utils.Error(err))
			} else {
				w.Write(data)
			}
		}
	}
}
