package utils

import (
	"encoding/base64"
	"encoding/json"
)

// Base64 decode then JSON decode
func Decode[T any](s string, v *T) error {
	if data, err := base64.StdEncoding.DecodeString(s); err != nil {
		return err
	} else if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
