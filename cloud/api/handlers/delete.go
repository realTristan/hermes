package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/api/utils"
)

// Delete is a handler function that returns a fiber context handler function for deleting a key from the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that deletes a key from the cache and returns a success message or an error message if the key is not provided.
func Delete(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Delete the key from the cache
		c.Delete(key)
		return ctx.Send(Utils.Success("null"))
	}
}
