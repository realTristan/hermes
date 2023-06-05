package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// Info is a function that returns information about the cache.
//
// Parameters:
//   - _ : A pointer to a utils.Params struct representing the parameters of the request. This parameter is ignored.
//   - c: A pointer to a hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A byte slice representing the information about the cache.
func Info(_ *utils.Params, c *hermes.Cache) []byte {
	if info, err := c.Info(); err != nil {
		return utils.Error(err)
	} else {
		return utils.Success(info)
	}
}

// InfoForTesting is a function that returns information about the cache for testing purposes.
//
// Parameters:
//   - _ : A pointer to a utils.Params struct representing the parameters of the request. This parameter is ignored.
//   - c: A pointer to a hermes.Cache struct representing the cache to get information from.
//
// Returns:
//   - A byte slice representing the information about the cache for testing purposes.
func InfoForTesting(_ *utils.Params, c *hermes.Cache) []byte {
	if info, err := c.InfoForTesting(); err != nil {
		return utils.Error(err)
	} else {
		return utils.Success(info)
	}
}
