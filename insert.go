package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// insert is a method of the FullText struct that inserts a value in the full-text cache for the specified key.
// This function is not thread-safe and should only be called from an exported function.
//
// Parameters:
//   - data: A map of maps containing the data to be inserted.
//
// Returns:
//   - An error if the full-text storage limit or byte-size limit is reached.
func (ft *FullText) insert(data map[string]map[string]interface{}) error {
	// Create a copy of the existing full-text variables
	var (
		tempStorage      map[string][]int = ft.storage
		tempIndices      map[int]string   = ft.indices
		tempCurrentIndex int              = ft.currentIndex
		tempKeys         map[string]int   = make(map[string]int)
	)

	// Loop through the json data
	for k, v := range tempIndices {
		tempKeys[v] = k
	}

	// Loop through the json data
	for cacheKey, cacheValue := range data {
		// Check if the key is in the temp keys map. If not, add it.
		if _, ok := tempKeys[cacheKey]; !ok {
			tempIndices[tempCurrentIndex] = cacheKey
			tempKeys[cacheKey] = tempCurrentIndex
			tempCurrentIndex++
		}

		// Loop through the map
		for k, v := range cacheValue {
			// Check if the value is a WFT
			var strv string
			if wft, ok := v.(WFT); ok {
				strv = wft.value
			} else if _strv := ftFromMap(v); len(_strv) > 0 {
				strv = _strv
			} else {
				continue
			}

			// Set the key in the provided value to the string wft value
			cacheValue[k] = strv

			// Clean the string value
			strv = strings.TrimSpace(strv)
			strv = Utils.RemoveDoubleSpaces(strv)
			strv = strings.ToLower(strv)

			// Loop through the words
			for _, word := range strings.Split(strv, " ") {
				if len(word) == 0 {
					continue
				}
				// Check if the storage limit has been reached
				if ft.maxLength > 0 {
					if len(tempStorage) > ft.maxLength {
						return fmt.Errorf("full-text storage limit reached (%d/%d keys). load cancelled", len(tempStorage), ft.maxLength)
					}
				}
				if ft.maxBytes > 0 {
					if cacheSize, err := Utils.Size(tempStorage); err != nil {
						return err
					} else if cacheSize > ft.maxBytes {
						return fmt.Errorf("full-text byte-size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxBytes)
					}
				}

				// Trim the word
				word = Utils.TrimNonAlphaNum(word)
				var words []string = Utils.SplitByAlphaNum(word)

				// Loop through the words
				for i := 0; i < len(words); i++ {
					// Check if the word is valid
					if len(words[i]) <= 3 {
						continue
					}
					if _, ok := tempStorage[words[i]]; !ok {
						tempStorage[words[i]] = []int{tempCurrentIndex}
						continue
					}
					if Utils.ContainsInt(tempStorage[words[i]], tempKeys[cacheKey]) {
						continue
					}
					tempStorage[words[i]] = append(tempStorage[words[i]], tempKeys[cacheKey])
				}
			}

			// Set the value in the cache data
			data[cacheKey] = cacheValue
		}
	}

	// Set the full-text cache to the temp map
	ft.storage = tempStorage
	ft.indices = tempIndices
	ft.currentIndex = tempCurrentIndex

	// Return nil for no errors
	return nil
}
