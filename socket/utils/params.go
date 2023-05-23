package utils

import (
	"encoding/json"
	"errors"
)

// Params is a struct that represents the query parameters.
// Fields:
//   - values: a map of the query parameters
type Params struct {
	values map[string]any
}

// ParseParams is a function that parses a JSON-encoded byte slice into a Params struct.
// Parameters:
//   - msg ([]byte): A JSON-encoded byte slice containing the parameters to parse.
//
// Returns:
//   - (*Params, error): A pointer to a Params struct and an error if the parsing fails, or nil if successful.
func ParseParams(msg []byte) (*Params, error) {
	var data = new(Params)
	if err := json.Unmarshal(msg, &data.values); err != nil {
		return nil, err
	}
	return data, nil
}

// Get is a method of the Params struct that retrieves the value of a query parameter with the provided key.
// Parameters:
//   - key (string): The key of the query parameter to retrieve.
//
// Returns:
//   - any: The value of the query parameter with the provided key.
func (p *Params) Get(key string) any {
	return p.values[key]
}

// GetFunction is a method of the Params struct that retrieves the value of the "function" query parameter.
// Returns:
//   - (string, error): The value of the "function" query parameter and an error if the parameter is not provided or is not a string, or nil if successful.
func (p *Params) GetFunction() (string, error) {
	if f, ok := p.Get("function").(string); !ok {
		return "", errors.New("no function provided")
	} else {
		return f, nil
	}
}

// GetKeyParam is a function that retrieves the value of the "key" query parameter from a Params struct.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//
// Returns:
//   - (string, error): The value of the "key" query parameter and an error if the parameter is not provided or is not a string, or nil if successful.
func GetKeyParam(p *Params) (string, error) {
	if k, ok := p.Get("key").(string); !ok || len(k) == 0 {
		return "", errors.New("no key provided")
	} else {
		return k, nil
	}
}

// GetQueryParam is a function that retrieves the value of the "query" query parameter from a Params struct.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//
// Returns:
//   - (string, error): The value of the "query" query parameter and an error if the parameter is not provided or is not a string, or nil if successful.
func GetQueryParam(p *Params) (string, error) {
	if q, ok := p.Get("query").(string); !ok || len(q) == 0 {
		return "", errors.New("no query provided")
	} else {
		return q, nil
	}
}

// GetValueParam is a generic function that retrieves the value of the "value" query parameter from a Params struct and decodes it into a provided value.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - value (*T): A pointer to a value of type T to decode the "value" query parameter into.
//
// Returns:
//   - error: An error if the decoding fails or the "value" query parameter is not provided or is not a string, or nil if successful.
func GetValueParam[T any](p *Params, value *T) error {
	if s, ok := p.Get("value").(string); !ok || len(s) == 0 {
		return errors.New("invalid value")
	} else {
		if err := Decode(s, &value); err != nil {
			return err
		}
	}
	return nil
}

// GetMaxLengthParam is a function that retrieves the value of the "maxlength" query parameter from a Params struct and stores it in a provided integer pointer.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - maxLength (*int): A pointer to an integer to store the value of the "maxlength" query parameter.
//
// Returns:
//   - error: An error if the "maxlength" query parameter is not provided or is not a float64, or nil if successful.
func GetMaxLengthParam(p *Params, maxLength *int) error {
	if i, ok := p.Get("maxlength").(float64); !ok {
		return errors.New("invalid maxlength")
	} else {
		*maxLength = int(i)
	}
	return nil
}

// GetMaxBytesParam is a function that retrieves the value of the "maxbytes" query parameter from a Params struct and stores it in a provided integer pointer.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - maxBytes (*int): A pointer to an integer to store the value of the "maxbytes" query parameter.
//
// Returns:
//   - error: An error if the "maxbytes" query parameter is not provided or is not a float64, or nil if successful.
func GetMaxBytesParam(p *Params, maxBytes *int) error {
	if i, ok := p.Get("maxbytes").(float64); !ok {
		return errors.New("invalid maxbytes")
	} else {
		*maxBytes = int(i)
	}
	return nil
}

// GetMinWordLengthParam is a function that retrieves the value of the "minwordlength" query parameter from a Params struct and stores it in a provided integer pointer.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - minWordLength (*int): A pointer to an integer to store the value of the "minwordlength" query parameter.
//
// Returns:
//   - error: An error if the "minwordlength" query parameter is not provided or is not a float64, or nil if successful.
func GetMinWordLengthParam(p *Params, minWordLength *int) error {
	if i, ok := p.Get("minwordlength").(float64); !ok {
		return errors.New("invalid minwordlength")
	} else {
		*minWordLength = int(i)
	}
	return nil
}

// GetJSONParam is a generic function that retrieves the value of the "json" query parameter from a Params struct and decodes it into a provided value.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - json (*T): A pointer to a value of type T to decode the "json" query parameter into.
//
// Returns:
//   - error: An error if the decoding fails or the "json" query parameter is not provided or is not a string, or nil if successful.
func GetJSONParam[T any](p *Params, json *T) error {
	if s, ok := p.Get("json").(string); !ok || len(s) == 0 {
		return errors.New("invalid json")
	} else if err := Decode(s, &json); err != nil {
		return err
	}
	return nil
}

// GetSchemaParam is a function that retrieves the value of the "schema" query parameter from a Params struct and decodes it into a provided map of string-boolean pairs.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - schema (*map[string]bool): A pointer to a map of string-boolean pairs to decode the "schema" query parameter into.
//
// Returns:
//   - error: An error if the decoding fails or the "schema" query parameter is not provided or is not a string, or nil if successful.
func GetSchemaParam(p *Params, schema *map[string]bool) error {
	if s, ok := p.Get("schema").(string); !ok || len(s) == 0 {
		return errors.New("invalid schema")
	} else if err := Decode(s, schema); err != nil {
		return err
	}
	return nil
}

// GetLimitParam is a function that retrieves the value of the "limit" query parameter from a Params struct and stores it in a provided integer pointer.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - limit (*int): A pointer to an integer to store the value of the "limit" query parameter.
//
// Returns:
//   - error: An error if the "limit" query parameter is not provided or is not a float64, or nil if successful.
func GetLimitParam(p *Params, limit *int) error {
	if i, ok := p.Get("limit").(float64); !ok {
		return errors.New("invalid limit")
	} else {
		*limit = int(i)
	}
	return nil
}

// GetStrictParam is a function that retrieves the value of the "strict" query parameter from a Params struct and stores it in a provided boolean pointer.
// Parameters:
//   - p (*Params): A pointer to a Params struct.
//   - strict (*bool): A pointer to a boolean to store the value of the "strict" query parameter.
//
// Returns:
//   - error: An error if the "strict" query parameter is not provided or is not a boolean, or nil if successful.
func GetStrictParam(p *Params, strict *bool) error {
	if b, ok := p.Get("strict").(bool); !ok {
		return errors.New("invalid strict")
	} else {
		*strict = b
	}
	return nil
}
