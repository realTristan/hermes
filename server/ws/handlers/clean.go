package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Clean the regular cache
// This is a handler function that returns a fiber context handler function
func Clean(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	c.Clean()
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}

// Clean the full-text storage
// This is a handler function that returns a fiber context handler function
func FTClean(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	if err := c.FTClean(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}
