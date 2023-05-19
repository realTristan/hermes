package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/api/utils"
)

// Check if key exists
// This is a handler function that returns a fiber context handler function
func Exists(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Return whether the key exists
		return ctx.Send(Utils.Success(c.Exists(key)))
	}
}
