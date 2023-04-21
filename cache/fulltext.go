package cache

import "unsafe"

/*
FullText represents a data structure for storing and searching full-text documents. It uses a map to cache word positions for each word in the text, and an array to store the keys in sorted order for faster iteration. The maximum number of words that can be cached and the maximum size of the text can be set when initializing the struct.

Fields:
- wordCache: a map of words to an array of positions where the word appears in the text
- keys: an array of keys sorted in ascending order
- maxWords: the maximum number of words to cache
- maxSizeBytes: the maximum size of the text to cache in bytes
- isInitialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	wordCache     map[string][]string
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}

func (c *Cache) FTWordCache() map[string][]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Copy the wordCache map
	var copy map[string][]string = make(map[string][]string, len(c.ft.wordCache))
	copy = c.ft.wordCache
	return copy
}

func (c *Cache) FTWordCacheSize() uintptr {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return unsafe.Sizeof(c.ft.wordCache)
}

func (c *Cache) FTIsInitialized() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.ft.isInitialized
}

func (c *Cache) FTSpaceLeft() uintptr {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return uintptr(c.ft.maxSizeBytes) - unsafe.Sizeof(c.ft.wordCache)
}
