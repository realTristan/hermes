package handlers

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Get Keys from cache
// This is a handler function that returns a fiber context handler function
func Keys(c *Hermes.Cache, ws *websocket.Conn) error {
	if keys, err := json.Marshal(c.Keys()); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, keys)
	}
}
