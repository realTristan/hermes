package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Initialize the full-text search cache
// This is a handler function that returns a fiber context handler function
func FTInit(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxLength int
			maxBytes  int
		)

		// Get the max length parameter
		if err := Utils.GetMaxLengthParam(ctx, &maxLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the max bytes parameter
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInit(maxLength, maxBytes); err != nil {
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
			maxLength int
			maxBytes  int
			json      map[string]map[string]interface{}
		)

		// Get the max length from the query
		if err := Utils.GetMaxLengthParam(ctx, &maxLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the max bytes from the query
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the JSON from the query
		if err := Utils.GetJSONParam(ctx, &json); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInitWithMap(json, maxLength, maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Return success message
		return ctx.Send(Utils.Success("null"))
	}
}
