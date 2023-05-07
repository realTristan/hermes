package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Set a value in the cache
// This is a handler function that returns a fiber context handler function
func Set(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			key   string
			ft    bool
			value map[string]interface{}
		)
		// Get the key from the query
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("invalid key"))
		}

		// Get the value from the query
		if err := Utils.GetValueParam(ctx, &value); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get whether or not to store the full-text
		if err := Utils.GetFTParam(ctx, &ft); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Set the value in the cache
		if err := c.Set(key, value); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}
