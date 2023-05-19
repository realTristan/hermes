package handlers

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/ws/utils"
)

// Check if full-text is initialized
// This is a handler function that returns a fiber context handler function
func FTIsInitialized(c *Hermes.Cache, ws *websocket.Conn) error {
	return ws.WriteMessage(websocket.TextMessage, Utils.Success(c.FTIsInitialized()))
}

// Set the full-text max bytes
// This is a handler function that returns a fiber context handler function
func FTSetMaxBytes(c *Hermes.Cache, ws *websocket.Conn) error {
	// Get the value from the query
	var value int
	if err := Utils.GetMaxBytesParam(ws, &value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Set the max bytes
	if err := c.FTSetMaxBytes(value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}

// Set the full-text maximum length
// This is a handler function that returns a fiber context handler function
func FTSetMaxLength(c *Hermes.Cache, ws *websocket.Conn) error {
	// Get the value from the query
	var value int
	if err := Utils.GetMaxLengthParam(ws, &value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}

	// Set the max length
	if err := c.FTSetMaxLength(value); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	}
	return ws.WriteMessage(websocket.TextMessage, Utils.Success("null"))
}

// Get the full-text storage
// This is a handler function that returns a fiber context handler function
func FTStorage(c *Hermes.Cache, ws *websocket.Conn) error {
	if data, err := c.FTStorage(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		// Marshal the data
		if data, err := json.Marshal(data); err != nil {
			return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
		} else {
			return ws.WriteMessage(websocket.TextMessage, data)
		}
	}
}

// Get the full-text sotrage length
// This is a handler function that returns a fiber context handler function
func FTStorageLength(c *Hermes.Cache, ws *websocket.Conn) error {
	if length, err := c.FTStorageLength(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, Utils.Success(length))
	}
}

// Get the full-text storage size
// This is a handler function that returns a fiber context handler function
func FTStorageSize(c *Hermes.Cache, ws *websocket.Conn) error {
	if size, err := c.FTStorageSize(); err != nil {
		return ws.WriteMessage(websocket.TextMessage, Utils.Error(err))
	} else {
		return ws.WriteMessage(websocket.TextMessage, Utils.Success(size))
	}
}
