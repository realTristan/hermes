package api

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/api/handlers"
)

// SetRoutes is a function that sets the routes for the Hermes Cache API.
// Parameters:
//   - app (*fiber.App): A pointer to a fiber.App struct.
//   - cache (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - void: This function does not return anything.
func SetRoutes(app *fiber.App, cache *Hermes.Cache) {
	// Dev Testing Handler
	app.Get("/dev/hermes", func(c *fiber.Ctx) error {
		return c.SendString("Hermes Cache API Successfully Running!")
	})

	// Cache Handlers
	app.Get("/cache/values", handlers.Values(cache))
	app.Get("/cache/length", handlers.Length(cache))
	app.Post("/cache/clean", handlers.Clean(cache))
	app.Post("/cache/set", handlers.Set(cache))
	app.Delete("/cache/delete", handlers.Delete(cache))
	app.Get("/cache/get", handlers.Get(cache))
	app.Get("/cache/get/all", handlers.GetAll(cache))
	app.Get("/cache/keys", handlers.Keys(cache))
	app.Get("/cache/info", handlers.Info(cache))
	app.Get("/cache/info/testing", handlers.InfoForTesting(cache))
	app.Get("/cache/exists", handlers.Exists(cache))

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
	app.Post("/ft/minwordlength", handlers.FTSetMinWordLength(cache))
	app.Get("/ft/storage", handlers.FTStorage(cache))
	app.Get("/ft/storage/size", handlers.FTStorageSize(cache))
	app.Get("/ft/storage/length", handlers.FTStorageLength(cache))
	app.Get("/ft/isinitialized", handlers.FTIsInitialized(cache))
	app.Post("/ft/indices/sequence", handlers.FTSequenceIndices(cache))
}
