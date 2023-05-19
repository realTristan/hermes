package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/api/utils"
)

// Get a key from the cache
// This is a handler function that returns a fiber context handler function
func Get(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Get the key from the query
		var key string
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("key not provided"))
		}

		// Get the value from the cache
		if data, err := json.Marshal(c.Get(key)); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// Get all the data from the cache
// This is a handler function that returns a fiber context handler function
func GetAll(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
