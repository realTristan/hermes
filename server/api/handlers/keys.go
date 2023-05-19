package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/api/utils"
)

// Get Keys from cache
// This is a handler function that returns a fiber context handler function
func Keys(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if keys, err := json.Marshal(c.Keys()); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(keys)
		}
	}
}
