package handlers

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/api/utils"
)

// Info is a function that returns information about the cache.
//
// Parameters:
//   - c: A pointer to a Hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A function that takes a pointer to a fiber.Ctx struct and returns an error.
//     The function returns information about the cache.
func Info(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.Info(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(info))
		}
	}
}

// InfoForTesting is a function that returns information about the cache for testing purposes.
//
// Parameters:
//   - c: A pointer to a Hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A function that takes a pointer to a fiber.Ctx struct and returns an error.
//     The function returns information about the cache for testing purposes.
func InfoForTesting(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.InfoForTesting(); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(Utils.Success(info))
		}
	}
}
