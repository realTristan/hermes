package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Delete a key from the cache
// This is a handler function that returns a fiber context handler function
func Delete(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Delete the key from the cache
		c.Delete(key)
		return ctx.Send(Utils.Success("null"))
	}
}
