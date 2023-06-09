package nocache

import (
	"sync"
)

/*
FullText is a struct that represents a full text search cache. It has the following fields:
- mutex (*sync.RWMutex): a pointer to a read-write mutex used to synchronize access to the cache
- storage (map[string]any): a map where the keys are words and the values are arrays of integers representing the indices of the data items that contain the word
- words ([]string): a slice of strings representing all the unique words in the cache
- data ([]map[string]any): a slice of maps representing the data items in the cache, where the keys are the names of the fields and the values are the field values
*/
type FullText struct {
	mutex   *sync.RWMutex
	storage map[string]any
	words   []string
	data    []map[string]any
}
