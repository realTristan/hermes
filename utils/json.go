package utils

import (
	"encoding/json"
	"os"
)

// ReadJson is a generic function that reads a JSON file and unmarshals its contents into a provided value of type T.
// Parameters:
//   - file (string): The path to the JSON file to read.
//
// Returns:
//   - T: The unmarshalled value of type T.
//   - error: An error if the file cannot be read or the unmarshalling fails, or nil if successful.
func ReadJson[T any](file string) (T, error) {
	var v T

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return *new(T), err
	} else if err := json.Unmarshal(data, &v); err != nil {
		return *new(T), err
	}
	return v, nil
}
