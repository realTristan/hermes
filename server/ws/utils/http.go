package utils

import (
	"fmt"
)

// Return an error body
func Error[T any](err T) []byte {
	return []byte(fmt.Sprintf(`{"success":false,"error":"%v"}`, err))
}

// Return a success body
func Success[T any](v T) []byte {
	return []byte(fmt.Sprintf(`{"success":true,"data":%v}`, v))
}
