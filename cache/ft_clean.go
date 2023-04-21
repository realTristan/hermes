package cache

import "errors"

/*
The Clean function is a method of the FullText struct that cleans the cache of the FullText index.
It uses Mutex Locking to ensure that the cleaning process is thread-safe and can be accessed by only one thread at a time.
@Parameters: None
@Returns: None
Usage:

	ft := &FullText{}
	ft.Clean()
*/
func (c *Cache) FTClean() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return errors.New("full text is not initialized")
	}
	c.ft.clean()
	return nil
}

/*
The `clean` method removes all data from the FullText index.

	This method is private and not meant to be called directly by external code. It is used internally by the `Cache` struct when the cache is cleared.

	The method removes all data from the wordCache map and sets the keys slice to an empty slice.

	Example usage:
	(not meant to be called directly)
*/
func (ft *FullText) clean() {
	ft.wordCache = make(map[string][]string, ft.maxWords)
}
