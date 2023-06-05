package handlers

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Set is a handler function that returns a fiber context handler function for setting a value in the cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets a value in the cache using the key and value parameters provided in the query string and returns a success message or an error message if the set fails or if the parameters are not provided.
func Set(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			key   string
			value map[string]interface{}
		)
		// Get the key from the query
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(utils.Error("invalid key"))
		}

		// Get the value from the query
		if err := utils.GetValueParam(ctx, &value); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Set the value in the cache
		if err := c.Set(key, value); err != nil {
			return ctx.Send(utils.Error(err))
		}
		return ctx.Send(utils.Success("null"))
	}
}
