package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Keys is a handler function that returns a fiber context handler function for getting all the keys from the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets all the keys from the cache and returns a JSON-encoded string of the keys or an error message if the retrieval or encoding fails.
func Keys(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if keys, err := json.Marshal(c.Keys()); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(keys)
		}
	}
}
