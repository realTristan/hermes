package cache

// Delete a value from the cache with Mutex Locking
func (ft *FullText) Delete(word string) {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.delete(word)
}

// Delete a word from the cache
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
