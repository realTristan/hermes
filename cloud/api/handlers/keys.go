package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Keys is a handler function that returns a fiber context handler function for getting all the keys from the cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets all the keys from the cache and returns a JSON-encoded string of the keys or an error message if the retrieval or encoding fails.
func Keys(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if keys, err := json.Marshal(c.Keys()); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(keys)
		}
	}
}
