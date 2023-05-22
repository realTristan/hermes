package utils

import (
	"strings"

	"github.com/gofiber/websocket/v2"
)

// Check whether the error is a close error by converting
// the error to a string and checking if it contains the
// word "close".
func IsCloseError(err error) bool {
	return strings.Contains(err.Error(), "close")
}

// Check whether the error is a close error using websocket
// error codes and IsCloseError function.
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
