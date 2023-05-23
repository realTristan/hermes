package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// FTSequenceIndices is a handler function that returns a fiber context handler function for sequencing the full-text storage indices.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sequences the full-text storage indices and returns a success message or an error message if the sequencing fails.
func FTSequenceIndices(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		c.FTSequenceIndices()
		return ctx.Send(Utils.Success("null"))
	}
}
