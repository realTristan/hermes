package utils

import (
	"strings"

	"github.com/gofiber/websocket/v2"
)

// IsCloseError is a function that checks if an error is a close error by converting
// the error to a string and checking if it contains the word "close".
// Parameters:
//   - err (error): The error to check.
//
// Returns:
//   - bool: true if the error is a close error, false otherwise.
func IsCloseError(err error) bool {
	return strings.Contains(err.Error(), "close")
}

// IsSocketCloseError is a function that checks if an error is a WebSocket close error.
// Parameters:
//   - err (error): The error to check.
//
// Returns:
//   - bool: true if the error is a WebSocket close error, false otherwise.
func IsSocketCloseError(err error) bool {
	return websocket.IsCloseError(
		err,
		websocket.CloseNormalClosure,
		websocket.CloseGoingAway,
		websocket.CloseAbnormalClosure,
		websocket.CloseNoStatusReceived,
		websocket.CloseInvalidFramePayloadData,
		websocket.ClosePolicyViolation,
		websocket.CloseMessageTooBig,
		websocket.CloseMandatoryExtension,
		websocket.CloseInternalServerErr,
		websocket.CloseServiceRestart,
		websocket.CloseTryAgainLater,
		websocket.CloseTLSHandshake,
		websocket.CloseMessage,
		websocket.CloseProtocolError,
		websocket.CloseUnsupportedData,
	)
}
