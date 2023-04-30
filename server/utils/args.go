package utils

import (
	"errors"
	"strings"
)

// Data struct
type Data struct {
	port interface{}
}

// Get the port
func (d *Data) Port() interface{} {
	var copy interface{} = d.port
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
				return data, errors.New("invalid args")
			}
			data.port = ":" + strings.ReplaceAll(args[i+1], ":", "")
			i = i + 1
			continue
		}
	}
	return data, nil
}
