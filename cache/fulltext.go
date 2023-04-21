package cache

import "fmt"

/*
FullText represents a data structure for storing and searching full-text documents.
It uses a map to cache word positions for each word in the text. The maximum number of words that
can be cached and the maximum size of the text can be set when initializing the struct.

Fields:
- wordCache: a map of words to an array of positions where the word appears in the text
- maxWords: the maximum number of words to cache
- maxSizeBytes: the maximum size of the text to cache in bytes
- initialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	wordCache    map[string][]string
	maxWords     int
	maxSizeBytes int
	initialized  bool
}

func (c *Cache) FTIsInitialized() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.ft.isInitialized()
}

func (ft *FullText) isInitialized() bool {
	return ft != nil && ft.initialized
}

func (c *Cache) FTSetMaxBytes(maxBytes int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return fmt.Errorf("full text cache not initialized")
	}
	c.ft.maxSizeBytes = maxBytes
	return nil
}

func (c *Cache) FTSetMaxWords(maxWords int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return fmt.Errorf("full text cache not initialized")
	}
	c.ft.maxWords = maxWords
	return nil
}

func (c *Cache) FTWordCache() map[string][]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// Copy the wordCache map
	var copy map[string][]string = make(map[string][]string, len(c.ft.wordCache))
	copy = c.ft.wordCache
	return copy
}

func (c *Cache) FTWordCacheSize() (int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return size(c.ft.wordCache)
}
