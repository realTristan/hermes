package handlers

import (
	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// FTInit is a handler function that returns a fiber context handler function for initializing the full-text search cache.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the initialization fails.
func FTInit(p *utils.Params, c *hermes.Cache) []byte {
	var (
		maxSize       int
		maxBytes      int
		minWordLength int
	)

	// Get the max length parameter
	if err := utils.GetMaxSizeParam(p, &maxSize); err != nil {
		return utils.Error(err)
	}

	// Get the max bytes parameter
	if err := utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return utils.Error(err)
	}

	// Get the min word length parameter
	if err := utils.GetMinWordLengthParam(p, &minWordLength); err != nil {
		return utils.Error(err)
	}

	// Initialize the full-text cache
	if err := c.FTInit(maxSize, maxBytes, minWordLength); err != nil {
		return utils.Error(err)
	}
	return utils.Success("null")
}

// FTInitJson is a handler function that returns a fiber context handler function for initializing the full-text search cache with a JSON object.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the initialization fails.
func FTInitJson(p *utils.Params, c *hermes.Cache) []byte {
	var (
		maxSize       int
		maxBytes      int
		minWordLength int
		json          map[string]map[string]any
	)

	// Get the max length from the query
	if err := utils.GetMaxSizeParam(p, &maxSize); err != nil {
		return utils.Error(err)
	}

	// Get the max bytes from the query
	if err := utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return utils.Error(err)
	}

	// Get the min word length from the query
	if err := utils.GetMinWordLengthParam(p, &minWordLength); err != nil {
		return utils.Error(err)
	}

	// Get the JSON from the query
	if err := utils.GetJSONParam(p, &json); err != nil {
		return utils.Error(err)
	}

	// Initialize the full-text cache
	if err := c.FTInitWithMap(json, maxSize, maxBytes, minWordLength); err != nil {
		return utils.Error(err)
	}

	// Return success message
	return utils.Success("null")
}
