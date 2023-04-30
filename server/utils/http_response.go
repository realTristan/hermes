package utils

import "fmt"

// Return an error body
func Error(err error) []byte {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
}

// Return a success body
func Success() []byte {
	return []byte(`{"success": true}`)
}
