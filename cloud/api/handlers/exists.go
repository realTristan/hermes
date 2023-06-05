package handlers

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Exists is a handler function that returns a fiber context handler function for checking if a key exists in the cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that checks if a key exists in the cache and returns a success message with a boolean value indicating whether the key exists or an error message if the key is not provided.
func Exists(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(utils.Error("key not provided"))
		}

		// Return whether the key exists
		return ctx.Send(utils.Success(c.Exists(key)))
	}
}
