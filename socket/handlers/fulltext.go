package handlers

import (
	"encoding/json"

	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/socket/utils"
)

// FTIsInitialized is a handler function that returns a fiber context handler function for checking if the full-text storage is initialized.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a boolean value indicating whether the full-text storage is initialized.
func FTIsInitialized(_ *Utils.Params, c *Hermes.Cache) []byte {
	return Utils.Success(c.FTIsInitialized())
}

// FTSetMaxBytes is a handler function that returns a fiber context handler function for setting the maximum number of bytes for the full-text storage.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the value is invalid or the setting fails.
func FTSetMaxBytes(p *Utils.Params, c *Hermes.Cache) []byte {
	// Get the value from the query
	var value int
	if err := Utils.GetMaxBytesParam(p, &value); err != nil {
		return Utils.Error(err)
	}

	// Set the max bytes
	if err := c.FTSetMaxBytes(value); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}

// FTSetMaxLength is a handler function that returns a fiber context handler function for setting the maximum length for the full-text storage.
// Parameters:
//   - p (*Utils.Params): A pointer to a Utils.Params struct.
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing a success message or an error message if the value is invalid or the setting fails.
func FTSetMaxLength(p *Utils.Params, c *Hermes.Cache) []byte {
	// Get the value from the query
	var value int
	if err := Utils.GetMaxLengthParam(p, &value); err != nil {
		return Utils.Error(err)
	}

	// Set the max length
	if err := c.FTSetMaxLength(value); err != nil {
		return Utils.Error(err)
	}
	return Utils.Success("null")
}

// FTStorage is a handler function that returns a fiber context handler function for retrieving the full-text storage.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the full-text storage or an error message if the retrieval fails.
func FTStorage(_ *Utils.Params, c *Hermes.Cache) []byte {
	if data, err := c.FTStorage(); err != nil {
		return Utils.Error(err)
	} else if data, err := json.Marshal(data); err != nil {
		return Utils.Error(err)
	} else {
		return data
	}
}

// FTStorageLength is a handler function that returns a fiber context handler function for retrieving the length of the full-text storage.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the length of the full-text storage or an error message if the retrieval fails.
func FTStorageLength(_ *Utils.Params, c *Hermes.Cache) []byte {
	if length, err := c.FTStorageLength(); err != nil {
		return Utils.Error(err)
	} else {
		return Utils.Success(length)
	}
}

// FTStorageSize is a handler function that returns a fiber context handler function for retrieving the size of the full-text storage.
// Parameters:
//   - _ (*Utils.Params): A pointer to a Utils.Params struct (unused).
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - []byte: A JSON-encoded byte slice containing the size of the full-text storage or an error message if the retrieval fails.
func FTStorageSize(_ *Utils.Params, c *Hermes.Cache) []byte {
	if size, err := c.FTStorageSize(); err != nil {
		return Utils.Error(err)
	} else {
		return Utils.Success(size)
	}
}
