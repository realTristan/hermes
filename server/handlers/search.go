package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Search for something in the cache
// This is a handler function that returns a http.HandlerFunc
func Search(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			strict bool
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = r.URL.Query().Get("q"); len(query) == 0 {
			w.Write(Utils.Error(errors.New("invalid query")))
			return
		}

		// Get the limit from the url params
		if limitStr := r.URL.Query().Get("limit"); len(limitStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid limit")))
			return
		} else {
			if limitInt, err := strconv.Atoi(limitStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				limit = limitInt
			}
		}

		// Get whether strict mode is enabled/disabled
		if strictStr := r.URL.Query().Get("strict"); len(strictStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid strict")))
			return
		} else {
			if strictBool, err := strconv.ParseBool(strictStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				strict = strictBool
			}
		}

		// Get the schema from the url params
		if schemaStr := r.URL.Query().Get("schema"); len(schemaStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid schema")))
			return
		} else {
			if err := Utils.Decode(schemaStr, &schema); err != nil {
				w.Write(Utils.Error(err))
				return
			}
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
// This is a handler function that returns a http.HandlerFunc
func SearchOneWord(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			strict bool
			query  string
			limit  int
		)

		// Get the query from the url params
		if query = r.URL.Query().Get("q"); len(query) == 0 {
			w.Write(Utils.Error(errors.New("invalid query")))
			return
		}

		// Get the limit from the url params
		if limitStr := r.URL.Query().Get("limit"); len(limitStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid limit")))
			return
		} else {
			if limitInt, err := strconv.Atoi(limitStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				limit = limitInt
			}
		}

		// Get whether strict mode is enabled/disabled
		if strictStr := r.URL.Query().Get("strict"); len(strictStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid strict")))
			return
		} else {
			if strictBool, err := strconv.ParseBool(strictStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				strict = strictBool
			}
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
// This is a handler function that returns a http.HandlerFunc
func SearchValues(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = r.URL.Query().Get("q"); len(query) == 0 {
			w.Write(Utils.Error(errors.New("invalid query")))
			return
		}

		// Get the limit from the url params
		if limitStr := r.URL.Query().Get("limit"); len(limitStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid limit")))
			return
		} else {
			if limitInt, err := strconv.Atoi(limitStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				limit = limitInt
			}
		}

		// Get the schema from the url params
		if schemaStr := r.URL.Query().Get("schema"); len(schemaStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid schema")))
			return
		} else {
			if err := Utils.Decode(schemaStr, &schema); err != nil {
				w.Write(Utils.Error(err))
				return
			}
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
// This is a handler function that returns a http.HandlerFunc
func SearchValuesWithKey(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			key    string
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = r.URL.Query().Get("q"); len(query) == 0 {
			w.Write(Utils.Error(errors.New("invalid query")))
			return
		}

		// Get the key from the url params
		if key = r.URL.Query().Get("key"); len(key) == 0 {
			w.Write(Utils.Error(errors.New("invalid key")))
			return
		}

		// Get the limit from the url params
		if limitStr := r.URL.Query().Get("limit"); len(limitStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid limit")))
			return
		} else {
			if limitInt, err := strconv.Atoi(limitStr); err != nil {
				w.Write(Utils.Error(err))
				return
			} else {
				limit = limitInt
			}
		}

		// Get the schema from the url params
		if schemaStr := r.URL.Query().Get("schema"); len(schemaStr) == 0 {
			w.Write(Utils.Error(errors.New("invalid schema")))
			return
		} else {
			if err := Utils.Decode(schemaStr, &schema); err != nil {
				w.Write(Utils.Error(err))
				return
			}
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
