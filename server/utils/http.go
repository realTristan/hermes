package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
func GetValueParam[T any](ctx *fiber.Ctx, value *T) error {
	if valueStr := ctx.Params("value"); len(valueStr) == 0 {
		return errors.New("invalid value")
	} else {
		if err := Decode(valueStr, &value); err != nil {
			return err
		}
	}
	return nil
}

// Get the maxwords url parameter
func GetMaxWordsParam(ctx *fiber.Ctx, maxWords *int) error {
	if maxWordsStr := ctx.Params("maxwords"); len(maxWordsStr) == 0 {
		return errors.New("invalid maxwords")
	} else {
		if maxWordsInt, err := strconv.Atoi(maxWordsStr); err != nil {
			return err
		} else {
			*maxWords = maxWordsInt
		}
	}
	return nil
}

// Get the maxbytes url parameter
func GetMaxBytesParam(ctx *fiber.Ctx, maxBytes *int) error {
	if maxBytesStr := ctx.Params("maxbytes"); len(maxBytesStr) == 0 {
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
func GetJSONParam[T any](ctx *fiber.Ctx, json *T) error {
	if jsonStr := ctx.Params("json"); len(jsonStr) == 0 {
		return errors.New("invalid json")
	} else {
		if err := Decode(jsonStr, &json); err != nil {
			return err
		}
	}
	return nil
}

// Get the full-text url parameter
func GetFTParam(ctx *fiber.Ctx, ft *bool) error {
	if ftStr := ctx.Params("ft"); len(ftStr) == 0 {
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
func GetSchemaParam(ctx *fiber.Ctx, schema *map[string]bool) error {
	// Get the schema from the url params
	if schemaStr := ctx.Params("schema"); len(schemaStr) == 0 {
		return errors.New("invalid schema")
	} else {
		if err := Decode(schemaStr, schema); err != nil {
			return err
		}
	}
	return nil
}

// Get the limit url parameter
func GetLimitParam(ctx *fiber.Ctx, limit *int) error {
	// Get the limit from the url params
	if limitStr := ctx.Params("limit"); len(limitStr) == 0 {
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
func GetStrictParam(ctx *fiber.Ctx, strict *bool) error {
	// Get whether strict mode is enabled/disabled
	if strictStr := ctx.Params("strict"); len(strictStr) == 0 {
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
