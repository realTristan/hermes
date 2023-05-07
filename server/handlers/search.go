package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Search for something in the cache
// This is a handler function that returns a fiber context handler function
func Search(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			strict bool
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = ctx.Params("q"); len(query) == 0 {
			return ctx.Send(Utils.Error("query not provided"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the strict from the url params
		if err := Utils.GetStrictParam(ctx, &strict); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the schema from the url params
		if err := Utils.GetSchemaParam(ctx, &schema); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.Search(query, limit, strict, schema); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				return ctx.Send(Utils.Error(err))
			} else {
				return ctx.Send(data)
			}
		}
	}
}

// Search for one word
// This is a handler function that returns a fiber context handler function
func SearchOneWord(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			strict bool
			query  string
			limit  int
		)

		// Get the query from the url params
		if query = ctx.Params("q"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the strict from the url params
		if err := Utils.GetStrictParam(ctx, &strict); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchOneWord(query, limit, strict); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				return ctx.Send(Utils.Error(err))
			} else {
				return ctx.Send(data)
			}
		}
	}
}

// Search in values
// This is a handler function that returns a fiber context handler function
func SearchValues(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = ctx.Params("q"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the schema from the url params
		if err := Utils.GetSchemaParam(ctx, &schema); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchValues(query, limit, schema); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				return ctx.Send(Utils.Error(err))
			} else {
				return ctx.Send(data)
			}
		}
	}
}

// Search for values
// This is a handler function that returns a fiber context handler function
func SearchValuesWithKey(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			key    string
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = ctx.Params("q"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the key from the url params
		if key = ctx.Params("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("invalid key"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the schema from the url params
		if err := Utils.GetSchemaParam(ctx, &schema); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchValuesWithKey(query, key, limit); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			if data, err := json.Marshal(res); err != nil {
				return ctx.Send(Utils.Error(err))
			} else {
				return ctx.Send(data)
			}
		}
	}
}
