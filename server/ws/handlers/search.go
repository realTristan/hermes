package handlers

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Search for something in the cache
// This is a handler function that returns a fiber context handler function
func Search(c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		strict bool
		query  string
		limit  int
		schema map[string]bool
	)

	// Get the query from the url params
	if query = ws.Query("q"); len(query) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("query not provided"))
	}

	// Get the limit from the url params
	if err := Utils.GetLimitParam(ws, &limit); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the strict from the url params
	if err := Utils.GetStrictParam(ws, &strict); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the schema from the url params
	if err := Utils.GetSchemaParam(ws, &schema); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Search for the query
	if res, err := c.Search(query, limit, strict, schema); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		if data, err := json.Marshal(res); err != nil {
			return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
		} else {
			return ws.WriteMessage(websocket.TextMessage, data)
		}
	}
}

// Search for one word
// This is a handler function that returns a fiber context handler function
func SearchOneWord(c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		strict bool
		query  string
		limit  int
	)

	// Get the query from the url params
	if query = ws.Query("q"); len(query) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("invalid query"))
	}

	// Get the limit from the url params
	if err := Utils.GetLimitParam(ws, &limit); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the strict from the url params
	if err := Utils.GetStrictParam(ws, &strict); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Search for the query
	if res, err := c.SearchOneWord(query, limit, strict); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		if data, err := json.Marshal(res); err != nil {
			return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
		} else {
			return ws.WriteMessage(websocket.TextMessage, data)
		}
	}
}

// Search in values
// This is a handler function that returns a fiber context handler function
func SearchValues(c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		query  string
		limit  int
		schema map[string]bool
	)

	// Get the query from the url params
	if query = ws.Query("q"); len(query) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("invalid query"))
	}

	// Get the limit from the url params
	if err := Utils.GetLimitParam(ws, &limit); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the schema from the url params
	if err := Utils.GetSchemaParam(ws, &schema); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Search for the query
	if res, err := c.SearchValues(query, limit, schema); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		if data, err := json.Marshal(res); err != nil {
			return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
		} else {
			return ws.WriteMessage(websocket.TextMessage, data)
		}
	}
}

// Search for values
// This is a handler function that returns a fiber context handler function
func SearchWithKey(c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		key    string
		query  string
		limit  int
		schema map[string]bool
	)

	// Get the query from the url params
	if query = ws.Query("q"); len(query) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("invalid query"))
	}

	// Get the key from the url params
	if key = ws.Query("key"); len(key) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("invalid key"))
	}

	// Get the limit from the url params
	if err := Utils.GetLimitParam(ws, &limit); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the schema from the url params
	if err := Utils.GetSchemaParam(ws, &schema); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Search for the query
	if res, err := c.SearchWithKey(query, key, limit); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		if data, err := json.Marshal(res); err != nil {
			return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
		} else {
			return ws.WriteMessage(websocket.TextMessage, data)
		}
	}
}
