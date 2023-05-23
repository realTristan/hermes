package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// FTIsInitialized is a handler function that returns a fiber context handler function for checking if the full-text search is initialized.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that checks if the full-text search is initialized and returns a success message with a boolean value indicating whether it is initialized.
func FTIsInitialized(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Send(Utils.Success(c.FTIsInitialized()))
	}
}

// FTSetMaxBytes is a handler function that returns a fiber context handler function for setting the maximum number of bytes for full-text search.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the maximum number of bytes for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
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

// FTSetMaxLength is a handler function that returns a fiber context handler function for setting the maximum length for full-text search.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the maximum length for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
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

// FTStorage is a handler function that returns a fiber context handler function for getting the full-text storage.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the full-text storage and returns a JSON-encoded string of the data or an error message if the retrieval or encoding fails.
func FTStorage(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if data, err := c.FTStorage(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else if data, err := json.Marshal(data); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// FTStorageLength is a handler function that returns a fiber context handler function for getting the length of the full-text storage.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the length of the full-text storage and returns a success message with the length or an error message if the retrieval fails.
func FTStorageLength(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if length, err := c.FTStorageLength(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(length))
		}
	}
}

// FTStorageSize is a handler function that returns a fiber context handler function for getting the size of the full-text storage.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that gets the size of the full-text storage and returns a success message with the size or an error message if the retrieval fails.
func FTStorageSize(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if size, err := c.FTStorageSize(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(size))
		}
	}
}

// FTSetMinWordLength is a handler function that returns a fiber context handler function for setting the minimum word length for full-text search.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that sets the minimum word length for full-text search and returns a success message or an error message if the value is not provided or if the setting fails.
func FTSetMinWordLength(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the min word length from the query
		var minWordLength int
		if err := Utils.GetMinWordLengthParam(ctx, &minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Update the min word length
		if err := c.FTSetMinWordLength(minWordLength); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Return null
		return ctx.Send(Utils.Success("null"))
	}
}
