package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Get cache info in the form of a string
// This is a handler function that returns a fiber context handler function
func Info(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.InfoString(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send([]byte(info))
		}
	}
}

// Get cache info in the form of a string for testing
func InfoForTesting(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.InfoStringForTesting(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send([]byte(info))
		}
	}
}