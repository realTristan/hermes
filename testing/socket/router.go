package main

import (
	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	Socket "github.com/realTristan/hermes/cloud/socket"
)

func main() {
	// Cache and fiber app
	cache := hermes.InitCache()
	app := fiber.New()

	// Set the router
	Socket.SetRouter(app, cache)

	// Listen on port 3000
	app.Listen(":3000")
}
