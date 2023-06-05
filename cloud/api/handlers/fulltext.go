package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// FTIsInitialized is a handler function that returns a fiber context handler function for checking if the full-text search is initialized.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that checks if the full-text search is initialized and returns a success message with a boolean value indicating whether it is initialized.
func FTIsInitialized(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Send(utils.Success(c.FTIsInitialized()))
	}
}

// FTSetMaxBytes is a handler function that returns a fiber context handler function for setting the maximum number of bytes for full-text search.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the maximum number of bytes for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
func FTSetMaxBytes(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the value from the query
		var value int
		if err := utils.GetMaxBytesParam(ctx, &value); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Set the max bytes
		if err := c.FTSetMaxBytes(value); err != nil {
			return ctx.Send(utils.Error(err))
		}
		return ctx.Send(utils.Success("null"))
	}
}

// FTSetMaxSize is a handler function that returns a fiber context handler function for setting the maximum length for full-text search.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the maximum length for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
func FTSetMaxSize(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the value from the query
		var value int
		if err := utils.GetMaxSizeParam(ctx, &value); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Set the max length
		if err := c.FTSetMaxSize(value); err != nil {
			return ctx.Send(utils.Error(err))
		}
		return ctx.Send(utils.Success("null"))
	}
}

// FTStorage is a handler function that returns a fiber context handler function for getting the full-text storage.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the full-text storage and returns a JSON-encoded string of the data or an error message if the retrieval or encoding fails.
func FTStorage(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if data, err := c.FTStorage(); err != nil {
			return ctx.Send(utils.Error(err))
		} else if data, err := json.Marshal(data); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// FTStorageLength is a handler function that returns a fiber context handler function for getting the length of the full-text storage.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the length of the full-text storage and returns a success message with the length or an error message if the retrieval fails.
func FTStorageLength(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if length, err := c.FTStorageLength(); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(utils.Success(length))
		}
	}
}

// FTStorageSize is a handler function that returns a fiber context handler function for getting the size of the full-text storage.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the size of the full-text storage and returns a success message with the size or an error message if the retrieval fails.
func FTStorageSize(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if size, err := c.FTStorageSize(); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(utils.Success(size))
		}
	}
}

// FTSetMinWordLength is a handler function that returns a fiber context handler function for setting the minimum word length for full-text search.
// Parameters:
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the minimum word length for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
func FTSetMinWordLength(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the min word length from the query
		var minWordLength int
		if err := utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Update the min word length
		if err := c.FTSetMinWordLength(minWordLength); err != nil {
			return ctx.Send(utils.Error(err))
		}

		// Return null
		return ctx.Send(utils.Success("null"))
	}
}
