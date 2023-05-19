package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Check if key exists
// This is a handler function that returns a fiber context handler function
func Exists(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	// Get the key from the query
	var key string
	if key = p.Get("key"); len(key) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("key not provided"))
	}

	// Return whether the key exists
	return ws.WriteMessage(websocket.TextMessage, Utils.Success(c.Exists(key)))
}
