package utils

import (
	"encoding/base64"
	"encoding/json"
)

func Decode(s string) (map[string]interface{}, error) {
	var (
		data []byte
		err  error
	)
	if data, err = base64.StdEncoding.DecodeString(s); err != nil {
		return nil, err
	}
	var value map[string]interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return value, nil
}
