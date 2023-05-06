package handlers

import (
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Clean the regular cache
// This is a handler function that returns a http.HandlerFunc
func Clean(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Clean()
		w.Write(Utils.Success())
	}
}

// Clean the full text cache
// This is a handler function that returns a http.HandlerFunc
func FTClean(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := c.FTClean(); err != nil {
			w.Write(Utils.Error(err))
			return
		}
		w.Write(Utils.Success())
	}
}
