package handlers

import (
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Initialize the full-text search cache
// This is a handler function that returns a fiber context handler function
func FTInit(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		maxLength int
		maxBytes  int
	)

	// Get the max length parameter
	if err := Utils.GetMaxLengthParam(p, &maxLength); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the max bytes parameter
	if err := Utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Initialize the full-text cache
	if err := c.FTInit(maxLength, maxBytes); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}

// Initialize the full-text search cache
// This is a handler function that returns a fiber context handler function
func FTInitJson(p *Utils.Params, c *Hermes.Cache, ws *websocket.Conn) error {
	var (
		maxLength int
		maxBytes  int
		json      map[string]map[string]interface{}
	)

	// Get the max length from the query
	if err := Utils.GetMaxLengthParam(p, &maxLength); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the max bytes from the query
	if err := Utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Get the JSON from the query
	if err := Utils.GetJSONParam(p, &json); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Initialize the full-text cache
	if err := c.FTInitWithMap(json, maxLength, maxBytes); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Return success message
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}
