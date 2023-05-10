package cache

// When you delete a number of keys from the cache, the currentIndex remains
// the same. Over time, this number will grow to be very large, and will
// cause the cache to use a lot of memory. This function resets the indices
// to be sequential, starting from 0.
// This function is thread-safe.
func (c *Cache) FTSequenceIndices() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.ft.sequenceIndices()
}

// When you delete a number of keys from the cache, the currentIndex remains
// the same. Over time, this number will grow to be very large, and will
// cause the cache to use a lot of memory. This function resets the indices
// to be sequential, starting from 0.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) sequenceIndices() {
	// Store the temp variables
	var (
		tempIndices      map[int]string = make(map[int]string)
		tempCurrentIndex int            = 0
		tempKeys         map[string]int = make(map[string]int)
	)

	// Fill the temp indices by iterating over the current
	// indices and adding them to the tempIndices map
	for _, value := range ft.indices {
		tempIndices[tempCurrentIndex] = value
		tempCurrentIndex++
	}

	// Fill the temp keys with the opposites of ft.indices
	for key, value := range tempIndices {
		tempKeys[value] = key
	}

	// Iterate over the ft storage
	for word, keys := range ft.storage {
		for i := 0; i < len(keys); i++ {
			var index int = keys[i]

			// Get the key from the old indices
			var key string = ft.indices[index]

			// Set the new index
			ft.storage[word][i] = tempKeys[key]
		}
	}

	// Set the old variables to the new variables
	ft.indices = tempIndices
	ft.currentIndex = tempCurrentIndex
}
