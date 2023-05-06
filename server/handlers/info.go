package handlers

import (
	"net/http"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/server/utils"
)

// Get cache info
// This is a handler function that returns a http.HandlerFunc
func Info(c *Hermes.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if info, err := c.Info(); err != nil {
			w.Write(Utils.Error(err))
		} else {
			w.Write([]byte(info))
		}
	}
}
