package utils

import (
	"encoding/json"
	"os"
)

// Read a json file and return the data as a map.
func ReadSliceJson(file string) ([]map[string]interface{}, error) {
	var v []map[string]interface{} = []map[string]interface{}{}

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
	}
	return v, nil
}

// Read a json file and return the data as a map.
func ReadMapJson(file string) (map[string]map[string]interface{}, error) {
	var v map[string]map[string]interface{} = map[string]map[string]interface{}{}

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
	}
	return v, nil
}
