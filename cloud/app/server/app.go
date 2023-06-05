package main

import (
	"log"
	"os"

	utils "hermes/utils"

	"github.com/gofiber/fiber/v2"
	hermes "github.com/realTristan/hermes"
	Socket "github.com/realTristan/hermes/cloud/socket"
)

// Main function
func main() {
	// Verify that the user is trying to serve the cache
	if len(os.Args) < 1 || os.Args[1] != "serve" {
		panic("incorrect usage. example: ./hermes serve -p {port}")
	}

	// Get the arg data
	var args, err = utils.GetArgData(os.Args)
	if err != nil || args.Port() == nil {
		panic("incorrect usage. example: ./hermes serve -p {port}")
	}

	// Get the port and json file
	var cache *hermes.Cache = hermes.InitCache()

	// Initialize a new fiber app
	var app *fiber.App = fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "hermes",
	})
	Socket.SetRouter(app, cache)

	// Listen on the port
	log.Fatal(app.Listen(args.Port().(string)))
}
