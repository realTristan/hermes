package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Sequence the ft storage indices
func FTSequenceIndices(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	c.FTSequenceIndices()
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}
