package utils

import (
	"encoding/base64"
	"encoding/json"
)

// Decode is a function that decodes a base64-encoded string and then decodes the resulting JSON into a value of type T.
// Parameters:
//   - s (string): The base64-encoded string to decode.
//   - v (*T): A pointer to a value of type T to store the decoded JSON.
//
// Returns:
//   - error: An error message if the decoding fails, or nil if the decoding is successful.
func Decode[T any](s string, v *T) error {
	if data, err := base64.StdEncoding.DecodeString(s); err != nil {
		return err
	} else if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
