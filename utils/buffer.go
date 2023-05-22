package utils

import (
	"bytes"
	"encoding/gob"
)

// Gets the real size in memory of a given value.
func Size(v any) (int, error) {
	var b *bytes.Buffer = new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}
