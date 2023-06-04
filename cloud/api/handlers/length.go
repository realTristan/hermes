package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/api/utils"
)

// Length is a handler function that returns a fiber context handler function for getting the length of the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the length of the cache and returns a success message with the length or an error message if the retrieval fails.
func Length(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Send(Utils.Success(c.Length()))
	}
}
