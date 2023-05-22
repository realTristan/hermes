package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// Info is a function that returns information about the cache.
//
// Parameters:
//   - _ : A pointer to a Utils.Params struct representing the parameters of the request. This parameter is ignored.
//   - c: A pointer to a Hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A byte slice representing the information about the cache.
func Info(_ *Utils.Params, c *Hermes.Cache) []byte {
	if info, err := c.Info(); err != nil {
		return Utils.Error(err)
	} else {
		return Utils.Success(info)
	}
}

// InfoForTesting is a function that returns information about the cache for testing purposes.
//
// Parameters:
//   - _ : A pointer to a Utils.Params struct representing the parameters of the request. This parameter is ignored.
//   - c: A pointer to a Hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A byte slice representing the information about the cache for testing purposes.
func InfoForTesting(_ *Utils.Params, c *Hermes.Cache) []byte {
	if info, err := c.InfoForTesting(); err != nil {
		return Utils.Error(err)
	} else {
		return Utils.Success(info)
	}
}
