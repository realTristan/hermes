package handlers

import (
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// FTInit is a handler function that returns a fiber context handler function for initializing the full-text search cache.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the initialization fails.
func FTInit(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		maxLength     int
		maxBytes      int
		minWordLength int
	)

	// Get the max length parameter
	if err := Utils.GetMaxLengthParam(p, &maxLength); err != nil {
		return Utils.Error(err)
	}

	// Get the max bytes parameter
	if err := Utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return Utils.Error(err)
	}

	// Get the min word length parameter
	if err := Utils.GetMinWordLengthParam(p, &minWordLength); err != nil {
		return Utils.Error(err)
	}

	// Initialize the full-text cache
	if err := c.FTInit(maxLength, maxBytes, minWordLength); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}

// FTInitJson is a handler function that returns a fiber context handler function for initializing the full-text search cache with a JSON object.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the initialization fails.
func FTInitJson(p *Utils.Params, c *Hermes.Cache) []byte {
	var (
		maxLength     int
		maxBytes      int
		minWordLength int
		json          map[string]map[string]any
	)

	// Get the max length from the query
	if err := Utils.GetMaxLengthParam(p, &maxLength); err != nil {
		return Utils.Error(err)
	}

	// Get the max bytes from the query
	if err := Utils.GetMaxBytesParam(p, &maxBytes); err != nil {
		return Utils.Error(err)
	}

	// Get the min word length from the query
	if err := Utils.GetMinWordLengthParam(p, &minWordLength); err != nil {
		return Utils.Error(err)
	}

	// Get the JSON from the query
	if err := Utils.GetJSONParam(p, &json); err != nil {
		return Utils.Error(err)
	}

	// Initialize the full-text cache
	if err := c.FTInitWithMap(json, maxLength, maxBytes, minWordLength); err != nil {
		return Utils.Error(err)
	}

	// Return success message
	return Utils.Success("null")
}
