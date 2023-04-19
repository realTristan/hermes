package cache

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
func (ft *FullText) Delete(word string) {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.delete(word)
}

/*
This function is used to delete a word from the cache. It takes a string key as
an argument and removes the corresponding key-value pair from the ft.data map.
Additionally, it searches for the key in the ft.wordCache map and removes it from
the slice of keys associated with its corresponding word. If the slice becomes empty
after the removal, the word is removed from the ft.wordCache map as well.

@Parameters:

	key (string): The key that needs to be removed from the cache.

@Returns: This function does not return any values.

Example usage:

	ft.delete("example") // Deletes the key "example" and its corresponding value from the cache.
*/
func (ft *FullText) delete(key string) {
	// Remove the key from the ft.data
	delete(ft.data, key)

	// Remove the key from the ft.wordCache
	for word, keys := range ft.wordCache {
		for i := 0; i < len(keys); i++ {
			if key != keys[i] {
				continue
			}

			// Remove the key from the ft.wordCache slice
			ft.wordCache[word] = append(ft.wordCache[word][:i], ft.wordCache[word][i+1:]...)
		}

		// If the ft.wordCache[word] is empty, remove it from the cache
		if len(ft.wordCache[word]) == 0 {
			delete(ft.wordCache, word)
		}
	}
}
