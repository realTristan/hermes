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

// Set the full-text maximum length
// This is a handler function that returns a fiber context handler function
func FTSetMaxLength(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the value from the query
		var value int
		if err := Utils.GetMaxLengthParam(ctx, &value); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Set the max length
		if err := c.FTSetMaxLength(value); err != nil {
			return ctx.Send(Utils.Error(err))
		}
		return ctx.Send(Utils.Success("null"))
	}
}

// Get the full-text storage
// This is a handler function that returns a fiber context handler function
func FTStorage(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if data, err := c.FTStorage(); err != nil {
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

// Get the full-text sotrage length
// This is a handler function that returns a fiber context handler function
func FTStorageLength(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if length, err := c.FTStorageLength(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(length))
		}
	}
}

// Get the full-text storage size
// This is a handler function that returns a fiber context handler function
func FTStorageSize(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if size, err := c.FTStorageSize(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(size))
		}
	}
}
