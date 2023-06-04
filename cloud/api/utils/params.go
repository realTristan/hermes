package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetValueParam is a function that retrieves a value from a query parameter in a Fiber context and decodes it into a value of type T.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - value (*T): A pointer to a value of type T to store the decoded value.
//
// Returns:
//   - error: An error message if the decoding fails or the query parameter is invalid, or nil if the decoding is successful.
func GetValueParam[T any](ctx *fiber.Ctx, value *T) error {
	if v := ctx.Query("value"); len(v) == 0 {
		return errors.New("invalid value")
	} else if err := Decode(v, &value); err != nil {
		return err
	}
	return nil
}

// GetMaxLengthParam is a function that retrieves the "maxlength" query parameter from a Fiber context and stores it in an integer pointer.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - maxLength (*int): A pointer to an integer to store the "maxlength" query parameter.
//
// Returns:
//   - error: An error message if the "maxlength" query parameter is invalid or cannot be converted to an integer, or nil if the retrieval is successful.
func GetMaxLengthParam(ctx *fiber.Ctx, maxLength *int) error {
	if s := ctx.Query("maxlength"); len(s) == 0 {
		fmt.Println(s)
		return errors.New("invalid maxlength")
	} else if i, err := strconv.Atoi(s); err != nil {
		return err
	} else {
		*maxLength = i
	}
	return nil
}

// GetMaxBytesParam is a function that retrieves the "maxbytes" query parameter from a Fiber context and stores it in an integer pointer.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - maxBytes (*int): A pointer to an integer to store the "maxbytes" query parameter.
//
// Returns:
//   - error: An error message if the "maxbytes" query parameter is invalid or cannot be converted to an integer, or nil if the retrieval is successful.
func GetMaxBytesParam(ctx *fiber.Ctx, maxBytes *int) error {
	if s := ctx.Query("maxbytes"); len(s) == 0 {
		return errors.New("invalid maxbytes")
	} else if i, err := strconv.Atoi(s); err != nil {
		return err
	} else {
		*maxBytes = i
	}
	return nil
}

// Get the min word length url parameter
func GetMinWordLengthParam(ctx *fiber.Ctx, minWordLength *int) error {
	if s := ctx.Query("minwordlength"); len(s) == 0 {
		return errors.New("invalid minwordlength")
	} else if i, err := strconv.Atoi(s); err != nil {
		return err
	} else {
		*minWordLength = i
	}
	return nil
}

// GetJSONParam is a function that retrieves a JSON-encoded value from a query parameter in a Fiber context and decodes it into a value of type T.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - json (*T): A pointer to a value of type T to store the decoded JSON.
//
// Returns:
//   - error: An error message if the decoding fails or the query parameter is invalid, or nil if the decoding is successful.
func GetJSONParam[T any](ctx *fiber.Ctx, json *T) error {
	if s := ctx.Query("json"); len(s) == 0 {
		return errors.New("invalid json")
	} else if err := Decode(s, &json); err != nil {
		return err
	}
	return nil
}

// GetSchemaParam is a function that retrieves a schema from a query parameter in a Fiber context and decodes it into a map of string keys and boolean values.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - schema (*map[string]bool): A pointer to a map of string keys and boolean values to store the decoded schema.
//
// Returns:
//   - error: An error message if the decoding fails or the query parameter is invalid, or nil if the decoding is successful.
func GetSchemaParam(ctx *fiber.Ctx, schema *map[string]bool) error {
	// Get the schema from the url params
	if s := ctx.Query("schema"); len(s) == 0 {
		return errors.New("invalid schema")
	} else if err := Decode(s, schema); err != nil {
		return err
	}
	return nil
}

// GetLimitParam is a function that retrieves the "limit" query parameter from a Fiber context and stores it in an integer pointer.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - limit (*int): A pointer to an integer to store the "limit" query parameter.
//
// Returns:
//   - error: An error message if the "limit" query parameter is invalid or cannot be converted to an integer, or nil if the retrieval is successful.
func GetLimitParam(ctx *fiber.Ctx, limit *int) error {
	// Get the limit from the url params
	if s := ctx.Query("limit"); len(s) == 0 {
		return errors.New("invalid limit")
	} else if i, err := strconv.Atoi(s); err != nil {
		return err
	} else {
		*limit = i
	}
	return nil
}

// GetLimitParam is a function that retrieves the "limit" query parameter from a Fiber context and stores it in an integer pointer.
// Parameters:
//   - ctx (*fiber.Ctx): A pointer to a Fiber context.
//   - limit (*int): A pointer to an integer to store the "limit" query parameter.
//
// Returns:
//   - error: An error message if the "limit" query parameter is invalid or cannot be converted to an integer, or nil if the retrieval is successful.
func GetStrictParam(ctx *fiber.Ctx, strict *bool) error {
	// Get whether strict mode is enabled/disabled
	if s := ctx.Query("strict"); len(s) == 0 {
		return errors.New("invalid strict")
	} else if b, err := strconv.ParseBool(s); err != nil {
		return err
	} else {
		*strict = b
	}
	return nil
}
