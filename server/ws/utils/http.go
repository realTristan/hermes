package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/websocket/v2"
)

// Return an error body
func Error[T any](err T) []byte {
	return []byte(fmt.Sprintf(`{"success":false,"error":"%v"}`, err))
}

// Return a success body
func Success[T any](v T) []byte {
	return []byte(fmt.Sprintf(`{"success":true,"data":%v}`, v))
}

// Get the value url parameter
func GetValueParam[T any](ws *websocket.Conn, value *T) error {
	if valueStr := ws.Query("value"); len(valueStr) == 0 {
		return errors.New("invalid value")
	} else {
		if err := Decode(valueStr, &value); err != nil {
			return err
		}
	}
	return nil
}

// Get the maxLength url parameter
func GetMaxLengthParam(ws *websocket.Conn, maxLength *int) error {
	if maxLengthStr := ws.Query("maxlength"); len(maxLengthStr) == 0 {
		return errors.New("invalid maxlength")
	} else {
		if maxLengthInt, err := strconv.Atoi(maxLengthStr); err != nil {
			return err
		} else {
			*maxLength = maxLengthInt
		}
	}
	return nil
}

// Get the maxbytes url parameter
func GetMaxBytesParam(ws *websocket.Conn, maxBytes *int) error {
	if maxBytesStr := ws.Query("maxbytes"); len(maxBytesStr) == 0 {
		return errors.New("invalid maxbytes")
	} else {
		if maxBytesInt, err := strconv.Atoi(maxBytesStr); err != nil {
			return err
		} else {
			*maxBytes = maxBytesInt
		}
	}
	return nil
}

// Get the json url parameter
func GetJSONParam[T any](ws *websocket.Conn, json *T) error {
	if jsonStr := ws.Query("json"); len(jsonStr) == 0 {
		return errors.New("invalid json")
	} else {
		if err := Decode(jsonStr, &json); err != nil {
			return err
		}
	}
	return nil
}

// Get the full-text url parameter
func GetFTParam(ws *websocket.Conn, ft *bool) error {
	if ftStr := ws.Query("ft"); len(ftStr) == 0 {
		return errors.New("invalid ft")
	} else {
		if ftBool, err := strconv.ParseBool(ftStr); err != nil {
			return err
		} else {
			*ft = ftBool
		}
	}
	return nil
}

// Get the schema url parameter
func GetSchemaParam(ws *websocket.Conn, schema *map[string]bool) error {
	// Get the schema from the url params
	if schemaStr := ws.Query("schema"); len(schemaStr) == 0 {
		return errors.New("invalid schema")
	} else {
		if err := Decode(schemaStr, schema); err != nil {
			return err
		}
	}
	return nil
}

// Get the limit url parameter
func GetLimitParam(ws *websocket.Conn, limit *int) error {
	// Get the limit from the url params
	if limitStr := ws.Query("limit"); len(limitStr) == 0 {
		return errors.New("invalid limit")
	} else {
		if limitInt, err := strconv.Atoi(limitStr); err != nil {
			return err
		} else {
			*limit = limitInt
		}
	}
	return nil
}

// Get the strict url parameter
func GetStrictParam(ws *websocket.Conn, strict *bool) error {
	// Get whether strict mode is enabled/disabled
	if strictStr := ws.Query("strict"); len(strictStr) == 0 {
		return errors.New("invalid strict")
	} else {
		if strictBool, err := strconv.ParseBool(strictStr); err != nil {
			return err
		} else {
			*strict = strictBool
		}
	}
	return nil
}
