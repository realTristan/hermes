package handlers

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Get a key from the cache
// This is a handler function that returns a fiber context handler function
func Get(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	// Get the key from the query
	var key string
	if key = p.Get("key"); len(key) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("key not provided"))
	}

	// Get the value from the cache
	if data, err := json.Marshal(c.Get(key)); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, data)
	}
}

// Get all the data from the cache
// This is a handler function that returns a fiber context handler function
func GetAll(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	return nil
}
