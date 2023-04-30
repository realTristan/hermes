package main

import (
	"fmt"
	"net"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/server/listener"
	"github.com/realTristan/Hermes/server/routes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Main function
func main() {
	// Verify that the user is trying to serve the cache
	if os.Args[1] != "serve" {
		return
	}

	// Get the arg data
	var args, err = Utils.GetArgData(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the port and json file
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize a new listener
	if l, err := listener.New(args.Port()); err != nil {
		panic(err)
	} else {
		// Establish a new gorilla mux router
		var router *mux.Router = mux.NewRouter()

		// Set router handlers
		routes.Set(cache, router)

		// Handle Router
		http.Handle("/", router)

		// Print the serving port
		var port int = l.Addr().(*net.TCPAddr).Port
		if err := Utils.PrintLogoWithPort(port); err != nil {
			panic(err)
		}

		// Serve the listener
		if err := http.Serve(l, nil); err != nil {
			fmt.Println(err)
		}
	}
}
