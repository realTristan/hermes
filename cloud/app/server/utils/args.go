package utils

import (
	"errors"
	"strings"
)

// Data struct
type Data struct {
	port any
}

// Get the port
func (d *Data) Port() any {
	var copy any = d.port
	return copy
}

// Get the argument data in a map
func GetArgData(args []string) (*Data, error) {
	var data *Data = &Data{
		port: nil,
	}

	// Iterate over the args
	for i := 2; i < len(args); i++ {
		// Port arg
		if args[i] == "-port" || args[i] == "-p" {
			if i+1 >= len(args) {
				return data, errors.New("invalid port")
			}

			// Add a ':' to the port
			data.port = ":" + strings.ReplaceAll(args[i+1], ":", "")

			// If all past index 1 isnt a number
			if !isNum(data.port.(string)[1:]) {
				return data, errors.New("invalid port")
			}

			// Increment i then continue
			i = i + 1
			continue
		}
	}
	return data, nil
}

// Check if a string is a number
func isNum(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
