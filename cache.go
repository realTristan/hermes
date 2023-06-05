package hermes

import (
	"sync"
)

// Cache is a struct that represents an in-memory cache of key-value pairs.
// The cache can be used to store arbitrary data, and supports concurrent access through a mutex.
// Additionally, the cache can be configured to support full-text search using a FullText index.
// Fields:
//   - data (map[string]map[string]any): A map that stores the data in the cache. The keys of the map are strings that represent the cache keys, and the values are sub-maps that store the actual data under string keys.
//   - mutex (*sync.RWMutex): A RWMutex that guards access to the cache data.
//   - ft (*FullText): A FullText index that can be used for full-text search. If nil, full-text search is disabled.
type Cache struct {
	data  map[string]map[string]any
	mutex *sync.RWMutex
	ft    *FullText
}
