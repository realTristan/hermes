package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/api/utils"
)

// Clean is a function that returns a fiber context handler function for cleaning the regular cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that cleans the regular cache and returns a success message.
func Clean(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		c.Clean()
		return ctx.Send(Utils.Success("null"))
	}
}

// FTClean is a function that returns a fiber context handler function for cleaning the full-text cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that cleans the full-text cache and returns a success message or an error message if the cleaning fails.
func FTClean(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if err := c.FTClean(); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}
