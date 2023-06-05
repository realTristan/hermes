package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// FTSequenceIndices is a handler function that returns a fiber context handler function for sequencing the full-text storage indices.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the sequencing fails.
func FTSequenceIndices(_ *utils.Params, c *hermes.Cache) []byte {
	c.FTSequenceIndices()
	return utils.Success("null")
}
