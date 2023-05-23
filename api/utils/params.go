package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get the value url parameter
func GetValueParam[T any](ctx *fiber.Ctx, value *T) error {
	if v := ctx.Query("value"); len(v) == 0 {
		return errors.New("invalid value")
	} else if err := Decode(v, &value); err != nil {
		return err
	}
	return nil
}

// Get the maxLength url parameter
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

// Get the maxbytes url parameter
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

// Get the json url parameter
func GetJSONParam[T any](ctx *fiber.Ctx, json *T) error {
	if s := ctx.Query("json"); len(s) == 0 {
		return errors.New("invalid json")
	} else if err := Decode(s, &json); err != nil {
		return err
	}
	return nil
}

// Get the schema url parameter
func GetSchemaParam(ctx *fiber.Ctx, schema *map[string]bool) error {
	// Get the schema from the url params
	if s := ctx.Query("schema"); len(s) == 0 {
		return errors.New("invalid schema")
	} else if err := Decode(s, schema); err != nil {
		return err
	}
	return nil
}

// Get the limit url parameter
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

// Get the strict url parameter
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
