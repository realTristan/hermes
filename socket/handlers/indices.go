package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// FTSequenceIndices is a handler function that returns a fiber context handler function for sequencing the full-text storage indices.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the sequencing fails.
func FTSequenceIndices(_ *Utils.Params, c *Hermes.Cache) []byte {
	c.FTSequenceIndices()
	return Utils.Success("null")
}
