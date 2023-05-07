package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Initialize the full-text search cache
// This is a handler function that returns a fiber context handler function
func FTInit(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxWords int
			maxBytes int
		)

		// Convert the max words to an integer
		if err := Utils.GetMaxWordsParam(ctx, &maxWords); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Convert the max size bytes to an integer
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInit(maxWords, maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}

// Initialize the full-text search cache
// This is a handler function that returns a fiber context handler function
func FTInitJson(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxWords int
			maxBytes int
			json     map[string]map[string]interface{}
		)

		// Convert the max words to an integer
		if err := Utils.GetMaxWordsParam(ctx, &maxWords); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Convert the max size bytes to an integer
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the json from the request url
		if err := Utils.GetJSONParam(ctx, &json); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInitWithMap(json, maxWords, maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Return success message
		return ctx.Send(Utils.Success("null"))
	}
}
