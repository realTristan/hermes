package handlers

import (
	"encoding/json"

	hermes "github.com/realTristan/hermes"
	utils "github.com/realTristan/hermes/cloud/socket/utils"
)

// FTIsInitialized is a handler function that returns a fiber context handler function for checking if the full-text storage is initialized.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a boolean value indicating whether the full-text storage is initialized.
func FTIsInitialized(_ *utils.Params, c *hermes.Cache) []byte {
	return utils.Success(c.FTIsInitialized())
}

// FTSetMaxBytes is a handler function that returns a fiber context handler function for setting the maximum number of bytes for the full-text storage.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the value is invalid or the setting fails.
func FTSetMaxBytes(p *utils.Params, c *hermes.Cache) []byte {
	// Get the value from the query
	var value int
	if err := utils.GetMaxBytesParam(p, &value); err != nil {
		return utils.Error(err)
	}

	// Set the max bytes
	if err := c.FTSetMaxBytes(value); err != nil {
		return utils.Error(err)
	}
	return utils.Success("null")
}

// FTSetMaxSize is a handler function that returns a fiber context handler function for setting the maximum length for the full-text storage.
// Parameters:
//   - p (*utils.Params): A pointer to a utils.Params struct.
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the value is invalid or the setting fails.
func FTSetMaxSize(p *utils.Params, c *hermes.Cache) []byte {
	// Get the value from the query
	var value int
	if err := utils.GetMaxSizeParam(p, &value); err != nil {
		return utils.Error(err)
	}

	// Set the max length
	if err := c.FTSetMaxSize(value); err != nil {
		return utils.Error(err)
	}
	return utils.Success("null")
}

// FTStorage is a handler function that returns a fiber context handler function for retrieving the full-text storage.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the full-text storage or an error message if the retrieval fails.
func FTStorage(_ *utils.Params, c *hermes.Cache) []byte {
	if data, err := c.FTStorage(); err != nil {
		return utils.Error(err)
	} else if data, err := json.Marshal(data); err != nil {
		return utils.Error(err)
	} else {
		return data
	}
}

// FTStorageLength is a handler function that returns a fiber context handler function for retrieving the length of the full-text storage.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the length of the full-text storage or an error message if the retrieval fails.
func FTStorageLength(_ *utils.Params, c *hermes.Cache) []byte {
	if length, err := c.FTStorageLength(); err != nil {
		return utils.Error(err)
	} else {
		return utils.Success(length)
	}
}

// FTStorageSize is a handler function that returns a fiber context handler function for retrieving the size of the full-text storage.
// Parameters:
//   - _ (*utils.Params): A pointer to a utils.Params struct (unused).
//   - c (*hermes.Cache): A pointer to a hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the size of the full-text storage or an error message if the retrieval fails.
func FTStorageSize(_ *utils.Params, c *hermes.Cache) []byte {
	if size, err := c.FTStorageSize(); err != nil {
		return utils.Error(err)
	} else {
		return utils.Success(size)
	}
}
