package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Base64 decode then JSON decode
func Decode[T any](s string, v *T) error {
	var (
		data []byte
		err  error
	)

	fmt.Println(s)

	// Decode the base64 string
	if data, err = base64.StdEncoding.DecodeString(s); err != nil {
		return err
	}

	// Unmarshal the data
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	// Return no error
	return nil
}
