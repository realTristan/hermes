package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/server/ws/handlers"
)

/*

app.Get("/values", handlers.Values(cache))
	app.Get("/length", handlers.Length(cache))
	app.Post("/clean", handlers.Clean(cache))
	app.Post("/set", handlers.Set(cache))
	app.Delete("/delete", handlers.Delete(cache))
	app.Get("/get", handlers.Get(cache))
	app.Get("/get/all", handlers.GetAll(cache))
	app.Get("/keys", handlers.Keys(cache))
	app.Get("/info", handlers.Info(cache))
	app.Get("/exists", handlers.Exists(cache))

	// Full-text Cache Handlers
	app.Post("/ft/init", handlers.FTInit(cache))
	app.Post("/ft/init/json", handlers.FTInitJson(cache))
	app.Post("/ft/clean", handlers.FTClean(cache))
	app.Get("/ft/search", handlers.Search(cache))
	app.Get("/ft/search/oneword", handlers.SearchOneWord(cache))
	app.Get("/ft/search/values", handlers.SearchValues(cache))
	app.Get("/ft/search/withkey", handlers.SearchWithKey(cache))
	app.Post("/ft/maxbytes", handlers.FTSetMaxBytes(cache))
	app.Post("/ft/maxlength", handlers.FTSetMaxLength(cache))
	app.Get("/ft/storage", handlers.FTStorage(cache))
	app.Get("/ft/storage/size", handlers.FTStorageSize(cache))
	app.Get("/ft/storage/length", handlers.FTStorageLength(cache))
	app.Get("/ft/isinitialized", handlers.FTIsInitialized(cache))
	app.Post("/ft/indices/sequence", handlers.FTSequenceIndices(cache))
*/

// Map of functions that can be called from the client
var functions = map[string]func(*Hermes.Cache, *websocket.Conn) error{
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
			if _, msg, err := c.ReadMessage(); err != nil {
				c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			} else if fn, ok := functions[string(msg)]; !ok {
				c.WriteMessage(websocket.TextMessage, []byte("Function not found"))
			} else if err := fn(cache, c); err != nil {
				c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			} else {
				c.WriteMessage(websocket.TextMessage, []byte("Success"))
			}
		}
	}))
}
