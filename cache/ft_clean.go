package cache

/*
The Clean function is a method of the FullText struct that cleans the cache of the FullText index.
It uses Mutex Locking to ensure that the cleaning process is thread-safe and can be accessed by only one thread at a time.
@Parameters: None
@Returns: None
Usage:

	ft := &FullText{}
	ft.Clean()
*/
func (ft *FullText) Clean() {
	ft.mutex.Lock()         // Locks the mutex to ensure thread-safety
	defer ft.mutex.Unlock() // Unlocks the mutex before returning
	ft.clean()              // Calls the clean function to clean the cache
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
	ft.wordCache = map[string][]string{}          // Clears the word cache by creating a new empty map
	ft.data = map[string]map[string]interface{}{} // Clears the data by creating a new empty map of maps
}
