package server

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/server/handlers"
)

// Set the routes
func SetRoutes(app *fiber.App, cache *Hermes.Cache) {
	// Dev Testing Handler
	app.Get("/dev/hermes", func(c *fiber.Ctx) error {
		return c.SendString("Hermes Cache API Successfully Running!")
	})

	// Cache Handlers
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
	app.Get("/ft/searchoneword", handlers.SearchOneWord(cache))
	app.Get("/ft/searchvalues", handlers.SearchValues(cache))
	app.Get("/ft/searchvalueswithkey", handlers.SearchValuesWithKey(cache))
	app.Post("/ft/maxbytes", handlers.FTSetMaxBytes(cache))
	app.Post("/ft/maxwords", handlers.FTSetMaxWords(cache))
	app.Get("/ft/cache", handlers.FTWordCache(cache))
	app.Get("/ft/cachesize", handlers.FTWordCacheSize(cache))
	app.Get("/ft/isinitialized", handlers.FTIsInitialized(cache))
}
