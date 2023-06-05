package handlers

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Delete is a handler function that returns a fiber context handler function for deleting a key from the cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that deletes a key from the cache and returns a success message or an error message if the key is not provided.
func Delete(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(utils.Error("key not provided"))
		}

		// Delete the key from the cache
		c.Delete(key)
		return ctx.Send(utils.Success("null"))
	}
}
