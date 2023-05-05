package cache

// Delete a key from the cache. If the key exists in the full-text cache,
// delete it from there as well.
// This function is thread-safe.
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.delete(key)
}

// Delete a key from the cache. If the key exists in the full-text cache,
// delete it from there as well.
// This function is not thread-safe, and should only be called from
// an exported function.
func (c *Cache) delete(key string) {
	// Delete the key from the FT cache
	if c.ft != nil {
		c.ft.delete(key)
	}

	// Delete the key from the cache
	delete(c.data, key)
}

// Delete a key from the full-text cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) delete(key string) {
	// Remove the key from the ft.wordCache
	for word, keys := range ft.wordCache {
		for i := 0; i < len(keys); i++ {
			if key != ft.indicesCache[keys[i]] {
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
