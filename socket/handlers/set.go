package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Set a value in the cache
// This is a handler function that returns a fiber context handler function
func Set(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		key   string
		ft    bool
		err   error
		value map[string]interface{}
	)

	// Get the key from the query
	if key, err = Utils.GetKeyParam(p); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error("invalid key"))
	}

	// Get the value from the query
	if err := Utils.GetValueParam(p, &value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get whether or not to store the full-text
	if err := Utils.GetFTParam(p, &ft); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Set the value in the cache
	if err := c.Set(key, value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}
