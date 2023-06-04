package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/api/utils"
)

// FTInit is a handler function that returns a fiber context handler function for initializing the full-text search cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that initializes the full-text search cache using the max length, max bytes, and min word length parameters provided in the query string and returns a success message or an error message if the parameters are not provided or if the initialization fails.
func FTInit(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxLength     int
			maxBytes      int
			minWordLength int
		)

		// Get the max length parameter
		if err := Utils.GetMaxLengthParam(ctx, &maxLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the max bytes parameter
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the min word length parameter
		if err := Utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInit(maxLength, maxBytes, minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}

// FTInitJson is a handler function that returns a fiber context handler function for initializing the full-text search cache with a JSON object.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that initializes the full-text search cache using a JSON object, max length, max bytes, and min word length parameters provided in the query string and returns a success message or an error message if the parameters are not provided or if the initialization fails.
func FTInitJson(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxLength     int
			maxBytes      int
			minWordLength int
			json          map[string]map[string]interface{}
		)

		// Get the max length from the query
		if err := Utils.GetMaxLengthParam(ctx, &maxLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the max bytes from the query
		if err := Utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the min word length from the query
		if err := Utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the JSON from the query
		if err := Utils.GetJSONParam(ctx, &json); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInitWithMap(json, maxLength, maxBytes, minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Return success message
		return ctx.Send(Utils.Success("null"))
	}
}
