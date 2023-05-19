package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Get cache info in the form of a string
// This is a handler function that returns a fiber context handler function
func Info(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	if info, err := c.InfoString(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, []byte(info))
	}
}

// Get cache info for testing in the form of a string
// This is a handler function that returns a fiber context handler function
func InfoForTesting(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	if info, err := c.InfoStringForTesting(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, []byte(info))
	}
}
