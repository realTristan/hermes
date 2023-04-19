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
func (c *Cache) CleanFT() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.ft == nil || !c.ft.isInitialized {
		return errors.New("full text is not initialized")
	}
	c.ft.clean()
	return nil
}

/*
*
The clean function is a private method of the FullText struct that is used to clear the cache of the FullText index.
@Parameters: None
@Returns: None
Usage:

	This function is called internally by the Clean method of the FullText struct to clear the cache.
	It should not be called directly from outside the package.
*/
func (ft *FullText) clean() {
	ft.wordCache = map[string][]string{}
}
