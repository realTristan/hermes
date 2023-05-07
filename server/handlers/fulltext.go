package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Check if full-text is initialized
// This is a handler function that returns a fiber context handler function
func FTIsInitialized(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Send(Utils.Success(c.FTIsInitialized()))
	}
}

// Set the full-text max bytes
// This is a handler function that returns a fiber context handler function
func FTSetMaxBytes(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the value from the query
		var value int
		if err := Utils.GetMaxBytesParam(ctx, &value); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Set the max bytes
		if err := c.FTSetMaxBytes(value); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}

// Set the full-text max words
// This is a handler function that returns a fiber context handler function
func FTSetMaxWords(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the value from the query
		var value int
		if err := Utils.GetMaxWordsParam(ctx, &value); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Set the max words
		if err := c.FTSetMaxWords(value); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}

// Get the full-text cache
// This is a handler function that returns a fiber context handler function
func FTWordCache(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if data, err := c.FTCache(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			// Marshal the data
			if data, err := json.Marshal(data); err != nil {
				return ctx.Send(Utils.Error(err))
			} else {
				return ctx.Send(data)
			}
		}
	}
}

// Get the full-text cache size
// This is a handler function that returns a fiber context handler function
func FTWordCacheSize(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if size, err := c.FTCacheSize(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(size))
		}
	}
}
