package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Add key from cache to the full-text cache
// This is a handler function that returns a fiber context handler function
func FTAdd(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Add the key to the full-text cache
		if err := c.FTAdd(key); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}
