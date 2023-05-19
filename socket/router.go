package ws

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/socket/handlers"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Map of functions that can be called from the client
var functions = map[string]func(*Utils.Params, *Hermes.Cache, *websocket.Conn) error{
	"length":              handlers.Length,
	"clean":               handlers.Clean,
	"set":                 handlers.Set,
	"delete":              handlers.Delete,
	"get":                 handlers.Get,
	"get.all":             handlers.GetAll,
	"keys":                handlers.Keys,
	"info":                handlers.Info,
	"exists":              handlers.Exists,
	"ft.init":             handlers.FTInit,
	"ft.init.json":        handlers.FTInitJson,
	"ft.clean":            handlers.FTClean,
	"ft.search":           handlers.Search,
	"ft.search.oneword":   handlers.SearchOneWord,
	"ft.search.values":    handlers.SearchValues,
	"ft.search.withkey":   handlers.SearchWithKey,
	"ft.maxbytes":         handlers.FTSetMaxBytes,
	"ft.maxlength":        handlers.FTSetMaxLength,
	"ft.storage":          handlers.FTStorage,
	"ft.storage.size":     handlers.FTStorageSize,
	"ft.storage.length":   handlers.FTStorageLength,
	"ft.isinitialized":    handlers.FTIsInitialized,
	"ft.indices.sequence": handlers.FTSequenceIndices,
}

// Socket is the websocket router
func SetRouter(app *fiber.App, cache *Hermes.Cache) {
	// Middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Main websocket handler
	app.Get("/ws/hermes/cache", websocket.New(func(c *websocket.Conn) {
		for {
			var (
				msg []byte
				err error
			)

			// Read the message
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			// Get the data
			var p *Utils.Params
			if p, err = Utils.ParseParams(msg); err != nil {
				log.Println("parse:", err)
				break
			}

			// Get the function
			var function string
			if function, err = p.GetFunction(); err != nil {
				log.Println("function:", err)
				break
			}

			// Check if the function exists
			if fn, ok := functions[function]; !ok {
				c.WriteMessage(websocket.TextMessage, []byte("Function not found"))
			} else if err = fn(p, cache, c); err != nil {
				log.Println("function:", err)
				break
			}
		}
	}))
}
