package handlers

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/api/utils"
)

// Info is a function that returns information about the cache.
//
// Parameters:
//   - c: A pointer to a hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A function that takes a pointer to a fiber.Ctx struct and returns an error.
//     The function returns information about the cache.
func Info(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.Info(); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(utils.Success(info))
		}
	}
}

// InfoForTesting is a function that returns information about the cache for testing purposes.
//
// Parameters:
//   - c: A pointer to a hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A function that takes a pointer to a fiber.Ctx struct and returns an error.
//     The function returns information about the cache for testing purposes.
func InfoForTesting(c *hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if info, err := c.InfoForTesting(); err != nil {
			return ctx.Send(utils.Error(err))
		} else {
			return ctx.Send(utils.Success(info))
		}
	}
}
