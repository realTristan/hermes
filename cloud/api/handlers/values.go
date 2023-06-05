package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Values is a handler function that returns a fiber context handler function for getting all values from the cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets all values from the cache and returns a JSON-encoded string of the values or an error message if the retrieval fails.
func Values(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if values, err := json.Marshal(c.Values()); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(values)
		}
	}
}
