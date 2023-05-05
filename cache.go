package cache

import (
	"sync"
)

/*
Fields:
  - data: A map that stores the data in the cache. The keys of the map are strings that represent the cache keys, and the values are sub-maps that store the actual data under string keys.
  - mutex: A RWMutex that guards access to the cache data.
  - ft: A FullText index that can be used for full-text search. If nil, full-text search is disabled.
*/
type Cache struct {
	data  map[string]map[string]interface{}
	mutex *sync.RWMutex
	ft    *FullText
}
