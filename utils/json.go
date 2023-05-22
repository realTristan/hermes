package utils

import (
	"encoding/json"
	"os"
)

// Read a json file and return the data as a map.
// map[string]map[string]interface{}
// []map[string]interface{}
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
