package cache

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
