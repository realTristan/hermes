package nocache

import (
	"sync"

	Utils "github.com/realTristan/Hermes/utils"
)

// Initialize the full-text cache with the provided data.
// This function is thread safe.
func InitWithMap(data []map[string]interface{}) (*FullText, error) {
	var ft *FullText = &FullText{
		mutex:   &sync.RWMutex{},
		storage: make(map[string]interface{}),
		words:   []string{},
		data:    []map[string]interface{}{},
	}

	// Load the cache data
	if err := ft.insert(data); err != nil {
		return nil, err
	}

	// Set the data
	ft.data = data

	// Return the full-text variable
	return ft, nil
}

// Initialize the full-text cache with the provided json file.
// This function is thread safe.
func InitWithJson(file string) (*FullText, error) {
	// Read the json data
	if data, err := Utils.ReadJson[[]map[string]interface{}](file); err != nil {
		return nil, err
	} else {
		return InitWithMap(data)
	}
}
