package cache

import "sync"

/*
The FullText struct represents a full text search cache that stores words and the keys of the documents where they appear.
It uses a read-write mutex to ensure safe access to the cache data, a map to store the words and their associated keys, and a map to store the document data.
The maxWords and maxSizeBytes fields represent optional limits on the size of the cache. If either limit is reached, additional data will not be added to the cache.
When initialized, the isInitialized field is set to true to indicate that the cache is ready for use.

@Variables:

	mutex *sync.RWMutex: a reader/writer mutex used to synchronize access to the wordCache and data maps when reading and writing to them.

	wordCache map[string][]string: a map of words to a slice of keys for items that contain the word. The keys represent the unique identifiers for each item in the cache.

	data map[string]map[string]interface{}: a map of keys to a map of attributes and their corresponding values. The keys represent the unique identifiers for each item in the cache.

	maxWords int: the maximum number of unique keys that can be stored in the wordCache map. If the cache size exceeds this value, an error is returned when attempting to add a new item.

	maxSizeBytes int: the maximum size of the wordCache map in bytes. If the cache size exceeds this value, an error is returned when attempting to add a new item.

	isInitialized bool: a flag indicating whether the cache has been initialized or not. If the cache is not initialized, attempting to read or write to it will result in an error.
*/
type FullText struct {
	mutex         *sync.RWMutex
	wordCache     map[string][]string
	data          map[string]map[string]interface{}
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}
