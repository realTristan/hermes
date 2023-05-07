package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Get cache info
// This is a handler function that returns a fiber context handler function
func Info(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.Info(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(info))
		}
	}
}
