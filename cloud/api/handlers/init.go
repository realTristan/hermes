package handlers

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// FTInit is a handler function that returns a fiber context handler function for initializing the full-text search cache.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that initializes the full-text search cache using the max length, max bytes, and min word length parameters provided in the query string and returns a success message or an error message if the parameters are not provided or if the initialization fails.
func FTInit(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxSize       int
			maxBytes      int
			minWordLength int
		)

		// Get the max length parameter
		if err := utils.GetMaxSizeParam(ctx, &maxSize); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Get the max bytes parameter
		if err := utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Get the min word length parameter
		if err := utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInit(maxSize, maxBytes, minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}
		return ctx.Send(utils.Success("null"))
	}
}

// FTInitJson is a handler function that returns a fiber context handler function for initializing the full-text search cache with a JSON object.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that initializes the full-text search cache using a JSON object, max length, max bytes, and min word length parameters provided in the query string and returns a success message or an error message if the parameters are not provided or if the initialization fails.
func FTInitJson(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			maxSize       int
			maxBytes      int
			minWordLength int
			json          map[string]map[string]interface{}
		)

		// Get the max length from the query
		if err := utils.GetMaxSizeParam(ctx, &maxSize); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Get the max bytes from the query
		if err := utils.GetMaxBytesParam(ctx, &maxBytes); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Get the min word length from the query
		if err := utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Get the JSON from the query
		if err := utils.GetJSONParam(ctx, &json); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Initialize the full-text cache
		if err := c.FTInitWithMap(json, maxSize, maxBytes, minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Return success message
		return ctx.Send(utils.Success("null"))
	}
}
