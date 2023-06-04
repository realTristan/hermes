package utils

import (
	"encoding/base64"
	"encoding/json"
)

// Decode is a generic function that decodes a base64-encoded string and unmarshals the resulting JSON into a provided value.
// Parameters:
//   - s (string): The base64-encoded string to decode and unmarshal.
//   - v (*T): A pointer to a value of type T to unmarshal the JSON into.
//
// Returns:
//   - error: An error if the decoding or unmarshaling fails, or nil if successful.
func Decode[T any](s string, v *T) error {
	if data, err := base64.StdEncoding.DecodeString(s); err != nil {
		return err
	} else if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
