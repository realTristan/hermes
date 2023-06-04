package utils

import (
	"fmt"
)

// Error is a function that returns an error message with the provided error.
// Parameters:
//   - err (T): The error to include in the error message.
//
// Returns:
//   - []byte: A byte slice containing the error message with the provided error.
func Error[T any](err T) []byte {
	return []byte(fmt.Sprintf(`{"success":false,"error":"%v"}`, err))
}

// Success is a function that returns a success message with the provided data.
// Parameters:
//   - v (T): The data to include in the success message.
//
// Returns:
//   - []byte: A byte slice containing the success message with the provided data.
func Success[T any](v T) []byte {
	return []byte(fmt.Sprintf(`{"success":true,"data":%v}`, v))
}
