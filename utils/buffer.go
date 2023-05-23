package utils

import (
	"bytes"
	"encoding/gob"
)

// Size is a function that calculates the real size in memory of a given value by encoding it using the gob package and returning the length of the resulting byte buffer.
// Parameters:
//   - v (any): The value to calculate the size of.
//
// Returns:
//   - int: The size of the value in memory.
//   - error: An error if the encoding fails, or nil if successful.
func Size(v any) (int, error) {
	var b *bytes.Buffer = new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}
