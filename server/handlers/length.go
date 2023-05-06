package handlers

import (
	"fmt"
	"net/http"

	Hermes "github.com/realTristan/Hermes"
)

// Get the cache length
// This is a handler function that returns a http.HandlerFunc
func Length(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%d", c.Length())))
	}
}
