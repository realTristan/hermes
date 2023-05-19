package utils

import (
	"encoding/json"
	"errors"
)

// Params is a struct that represents the query parameters.
// Fields:
//   - values: a map of the query parameters
type Params struct {
	values map[string]interface{}
}

// Get the parameters from the provided data map
func ParseParams(msg []byte) (*Params, error) {
	var data = new(Params)
	if err := json.Unmarshal(msg, &data.values); err != nil {
		return nil, err
	}
	return data, nil
}

// Get the value of a query param
func (p *Params) Get(key string) interface{} {
	return p.values[key]
}

// Get the function parameter
func (p *Params) GetFunction() (string, error) {
	if f, ok := p.Get("function").(string); !ok {
		return "", errors.New("no function provided")
	} else {
		return f, nil
	}
}

// Get the key parameter
func GetKeyParam(p *Params) (string, error) {
	if k, ok := p.Get("key").(string); !ok || len(k) == 0 {
		return "", errors.New("no key provided")
	} else {
		return k, nil
	}
}

// Get the query parameter
func GetQueryParam(p *Params) (string, error) {
	if q, ok := p.Get("query").(string); !ok || len(q) == 0 {
		return "", errors.New("no query provided")
	} else {
		return q, nil
	}
}

// Get the value parameter
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

// Get the maxLength parameter
func GetMaxLengthParam(p *Params, maxLength *int) error {
	if i, ok := p.Get("maxlength").(float64); !ok {
		return errors.New("invalid maxlength")
	} else {
		*maxLength = int(i)
	}
	return nil
}

// Get the maxbytes parameter
func GetMaxBytesParam(p *Params, maxBytes *int) error {
	if i, ok := p.Get("maxbytes").(float64); !ok {
		return errors.New("invalid maxbytes")
	} else {
		*maxBytes = int(i)
	}
	return nil
}

// Get the json parameter
func GetJSONParam[T any](p *Params, json *T) error {
	if s, ok := p.Get("json").(string); !ok || len(s) == 0 {
		return errors.New("invalid json")
	} else if err := Decode(s, &json); err != nil {
		return err
	}
	return nil
}

// Get the full-text parameter
func GetFTParam(p *Params, ft *bool) error {
	if b, ok := p.Get("ft").(bool); !ok {
		return errors.New("invalid ft")
	} else {
		*ft = b
	}
	return nil
}

// Get the schema parameter
func GetSchemaParam(p *Params, schema *map[string]bool) error {
	if s, ok := p.Get("schema").(string); !ok || len(s) == 0 {
		return errors.New("invalid schema")
	} else if err := Decode(s, schema); err != nil {
		return err
	}
	return nil
}

// Get the limit parameter
func GetLimitParam(p *Params, limit *int) error {
	if i, ok := p.Get("limit").(float64); !ok {
		return errors.New("invalid limit")
	} else {
		*limit = int(i)
	}
	return nil
}

// Get the strict parameter
func GetStrictParam(p *Params, strict *bool) error {
	if b, ok := p.Get("strict").(bool); !ok {
		return errors.New("invalid strict")
	} else {
		*strict = b
	}
	return nil
}
