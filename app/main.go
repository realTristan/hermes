package main

import (
	"log"
	"os"

	Utils "app/utils"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	server "github.com/realTristan/Hermes/server"
)

// Main function
func main() {
	// Update hermes
	Utils.UpdateHermes()

	// Verify that the user is trying to serve the cache
	if len(os.Args) < 1 || os.Args[1] != "serve" {
		panic("incorrect usage. example: ./hermes serve -p {port}")
	}

	// Get the arg data
	var args, err = Utils.GetArgData(os.Args)
	if err != nil || args.Port() == nil {
		panic("incorrect usage. example: ./hermes serve -p {port}")
	}

	// Get the port and json file
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize a new fiber app
	var app *fiber.App = fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "Hermes",
	})
	server.SetRoutes(app, cache)

	// Listen on the port
	log.Fatal(app.Listen(args.Port().(string)))
}
