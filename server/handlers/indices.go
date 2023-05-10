package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Sequence the ft storage indices
func FTSequenceIndices(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		c.FTSequenceIndices()
		return ctx.Send(Utils.Success("null"))
	}
}
