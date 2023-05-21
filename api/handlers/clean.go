package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Clean the regular cache
// This is a handler function that returns a fiber context handler function
func Clean(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		c.Clean()
		return ctx.Send(Utils.Success("null"))
	}
}

// Clean the full-text storage
// This is a handler function that returns a fiber context handler function
func FTClean(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if err := c.FTClean(); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}
