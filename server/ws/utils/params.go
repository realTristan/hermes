package utils

import (
	"errors"
	"strconv"
	"strings"
)

// Params is a struct that represents the query parameters of a url.
// Fields:
//   - values (map[string]string): A map that stores the query parameters. The keys of the map are the query parameter names, and the values are the query parameter values.
type Params struct {
	values map[string]string
}

// Get the query params from the url by splitting the url on the ? and &
func GetParams(url string) *Params {
	var p *Params = &Params{values: make(map[string]string)}
	if !strings.Contains(url, "?") {
		return p
	}
	for _, param := range strings.Split(url, "?")[1:] {
		for _, value := range strings.Split(param, "&") {
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					if value[i] == '=' {
						p.values[value[:i]] = value[i+1:]
						break
					}
				}
			}
		}
	}
	return p
}

// Get the value of a query param
func (p *Params) Get(key string) string {
	return p.values[key]
}

// Get the function of a query param
func GetFunction(msg string) string {
	var split = strings.Split(msg, "?")
	return split[0]
}

// Get the value url parameter
func GetValueParam[T any](p *Params, value *T) error {
	if valueStr := p.Get("value"); len(valueStr) == 0 {
		return errors.New("invalid value")
	} else {
		if err := Decode(valueStr, &value); err != nil {
			return err
		}
	}
	return nil
}

// Get the maxLength url parameter
func GetMaxLengthParam(p *Params, maxLength *int) error {
	if maxLengthStr := p.Get("maxlength"); len(maxLengthStr) == 0 {
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
func GetMaxBytesParam(p *Params, maxBytes *int) error {
	if maxBytesStr := p.Get("maxbytes"); len(maxBytesStr) == 0 {
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
func GetJSONParam[T any](p *Params, json *T) error {
	if jsonStr := p.Get("json"); len(jsonStr) == 0 {
		return errors.New("invalid json")
	} else {
		if err := Decode(jsonStr, &json); err != nil {
			return err
		}
	}
	return nil
}

// Get the full-text url parameter
func GetFTParam(p *Params, ft *bool) error {
	if ftStr := p.Get("ft"); len(ftStr) == 0 {
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
func GetSchemaParam(p *Params, schema *map[string]bool) error {
	if schemaStr := p.Get("schema"); len(schemaStr) == 0 {
		return errors.New("invalid schema")
	} else {
		if err := Decode(schemaStr, schema); err != nil {
			return err
		}
	}
	return nil
}

// Get the limit url parameter
func GetLimitParam(p *Params, limit *int) error {
	if limitStr := p.Get("limit"); len(limitStr) == 0 {
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
func GetStrictParam(p *Params, strict *bool) error {
	if strictStr := p.Get("strict"); len(strictStr) == 0 {
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
