package main

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	api "github.com/realTristan/Hermes/api"
)

func main() {
	app := fiber.New()
	cache := Hermes.InitCache()
	api.SetRoutes(app, cache)
	app.Listen(":3000")
}
