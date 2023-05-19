package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Delete a key from the cache
// This is a handler function that returns a fiber context handler function
func Delete(c *Hermes.Cache, ws *websocket.Conn) error {
	// Get the key from the query
	var key string
	if key = ws.Query("key"); len(key) == 0 {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("key not provided"))
	}

	// Delete the key from the cache
	c.Delete(key)
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}
