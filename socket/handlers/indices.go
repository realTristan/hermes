package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Sequence the ft storage indices
func FTSequenceIndices(_ *Utils.Params, c *Hermes.Cache) []byte {
	c.FTSequenceIndices()
	return Utils.Success("null")
}
