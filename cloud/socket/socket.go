package ws

import (
	"sync"
)

// Socket struct for storing the socket state and fiber app
type Socket struct {
	mutex  *sync.Mutex
	active bool
}

// Set the socket to active
func (s *Socket) SetActive() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.active = true
}

// Set the socket to inactive
func (s *Socket) SetInactive() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.active = false
}

// Get the socket state
func (s *Socket) IsActive() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.active
}
