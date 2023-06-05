package main

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	api "github.com/realTristan/hermes/cloud/api"
)

func main() {
	app := fiber.New()
	cache := hermes.InitCache()
	api.SetRoutes(app, cache)
	app.Listen(":3000")
}
