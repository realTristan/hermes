package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Get is a handler function that returns a fiber context handler function for getting a value from the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets a value from the cache using a key provided in the query string and returns a JSON-encoded string of the value or an error message if the key is not provided or if the retrieval or encoding fails.
func Get(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Get the value from the cache
		if data, err := json.Marshal(c.Get(key)); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// GetAll is a handler function that returns a fiber context handler function for getting all the data from the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets all the data from the cache and returns a success message with a JSON-encoded string of the data or an error message if the retrieval or encoding fails.
func GetAll(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
