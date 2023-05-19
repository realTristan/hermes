package main

import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Socket "github.com/realTristan/Hermes/server/ws"
)

func main() {
	// Cache and fiber app
	cache := Hermes.InitCache()
	app := fiber.New()

	// Set the router
	Socket.SetRouter(app, cache)

	// Listen on port 3000
	app.Listen(":3000")
}

// python3 test.py
