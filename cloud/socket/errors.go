package ws

import (
	"strings"
)

// Check whether the error is a close error by converting
// the error to a string and checking if it contains the
// word "close".
func IsCloseError(err error) bool {
	return strings.Contains(err.Error(), "close")
}
