package cache

import "errors"

/*
This function is used to delete a value from the cache with Mutex Locking. It takes a string word as
an argument and locks the mutex to prevent concurrent access while calling the delete method.
Once the delete method returns, it unlocks the mutex.

Note that this function is simply a wrapper around the delete method that adds mutex locking to prevent concurrent access.

@Parameters:

	word (string): The word that needs to be removed from the cache.

@Returns:

	This function does not return any values.

Example usage:

	ft.Delete("example") // Deletes the key "example" and its corresponding value from the cache with Mutex Locking.
*/
func (c *Cache) DeleteFT(word string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.ft.isInitialized() {
		return errors.New("full text is not initialized")
	}
	c.ft.delete(word)
	return nil
}

/*
The `deleteFT` method removes a key from the FullText index.

	This method is private and not meant to be called directly by external code. It is used internally by the `Cache` struct when a key is deleted from the cache.

	The method removes the key from the ft.keys slice and removes all references to the key from the ft.wordCache map.

	Example usage:
	(not meant to be called directly)
*/
func (ft *FullText) delete(key string) {
	// Remove the key from the ft.wordCache
	for word, keys := range ft.wordCache {
		for i := 0; i < len(keys); i++ {
			if key != keys[i] {
				continue
			}

			// Remove the key from the ft.wordCache slice
			ft.wordCache[word] = append(ft.wordCache[word][:i], ft.wordCache[word][i+1:]...)
			break
		}

		// If the ft.wordCache[word] is empty, remove it from the cache
		if len(ft.wordCache[word]) == 0 {
			delete(ft.wordCache, word)
		}
	}
}
