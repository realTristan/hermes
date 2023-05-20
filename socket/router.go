package ws

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Set the router for the socket
func SetRouter(app *fiber.App, cache *Hermes.Cache) {
	// Init a new socket
	var socket *Socket = &Socket{
		active: false,
		app:    app,
		mutex:  &sync.Mutex{},
	}

	// Middleware
	socket.app.Use("/ws", func(c *fiber.Ctx) error {
		// Check if the socket is active
		if socket.IsActive() {
			return fiber.ErrLocked
		}

		// Check if the request is via socket
		if websocket.IsWebSocketUpgrade(c) {
			// Allow Locals
			c.Locals("allowed", true)

			// Set the socket to active
			socket.SetActive()

			// Return the next handler
			return c.Next()
		}

		// Return an error
		return fiber.ErrUpgradeRequired
	})

	// Main websocket handler
	socket.app.Get("/ws/hermes", websocket.New(func(c *websocket.Conn) {
		for {
			var (
				msg []byte
				err error
			)

			// Read the message
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				if IsCloseError(err) {
					socket.SetInactive()
				}
				break
			}

			// Get the data
			var p *Utils.Params
			if p, err = Utils.ParseParams(msg); err != nil {
				log.Println("parse:", err)
				break
			}

			// Get the function
			var function string
			if function, err = p.GetFunction(); err != nil {
				log.Println("function:", err)
				break
			}

			// Check if the function exists
			if fn, ok := Functions[function]; !ok {
				if err = c.WriteMessage(websocket.TextMessage, []byte("Function not found")); err != nil {
					log.Println("write:", err)
					break
				}
			} else if err = fn(p, cache, c); err != nil {
				log.Println("function:", err)
				break
			}
		}
	}))
}
