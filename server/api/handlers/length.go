package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/api/utils"
)

// Get the cache length
// This is a handler function that returns a fiber context handler function
func Length(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Send(Utils.Success(c.Length()))
	}
}
