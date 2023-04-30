package listener

import (
	"errors"
	"net"
)

// Initialize a new listener
func New(port interface{}) (net.Listener, error) {
	if port == nil {
		return new(":0")
	}
	if p, ok := port.(string); !ok {
		return nil, errors.New("failed to parse port")
	} else {
		return new(p)
	}
}

// Initialize a new listener
func new(port string) (net.Listener, error) {
	if listener, err := net.Listen("tcp", port); err != nil {
		return nil, err
	} else {
		return listener, nil
	}
}
