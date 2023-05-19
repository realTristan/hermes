package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Get the cache length
// This is a handler function that returns a fiber context handler function
func Length(c *Hermes.Cache, ws *websocket.Conn) error {
	return ws.WriteMessage(websocket.TextMessage, Utils.Success(c.Length()))
}
