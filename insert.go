package cache

import (
	"fmt"
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// NewTempStorage is a function that creates a new TempStorage object for a given FullText object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to create the TempStorage object for.
//
// Returns:
//   - (*TempStorage): A pointer to the newly created TempStorage object.
type TempStorage struct {
	data    map[string]any
	indices map[int]string
	index   int
	keys    map[string]int
}

// NewTempStorage is a function that creates a new TempStorage object for a given FullText object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to create the TempStorage object for.
//
// Returns:
//   - (*TempStorage): A pointer to the newly created TempStorage object.
func NewTempStorage(ft *FullText) *TempStorage {
	var ts = &TempStorage{
		data:    ft.storage,
		indices: ft.indices,
		index:   ft.index,
		keys:    make(map[string]int),
	}

	// Loop through the data
	for k, v := range ts.indices {
		ts.keys[v] = k
	}
	return ts
}

// updateFullText is a method of the TempStorage struct that updates the FullText object with the data in the TempStorage object.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to update.
//
// Returns:
//   - None.
func (ts *TempStorage) updateFullText(ft *FullText) {
	ft.storage = ts.data
	ft.indices = ts.indices
	ft.index = ts.index
}

// cleanSingleArrays is a method of the TempStorage struct that replaces single-element integer arrays with their single integer value.
// Parameters:
//   - None.
//
// Returns:
//   - None.
func (ts *TempStorage) cleanSingleArrays() {
	for k, v := range ts.data {
		if v, ok := v.([]int); ok && len(v) == 1 {
			ts.data[k] = v[0]
		}
	}
}

// error is a method of the TempStorage struct that checks if the storage limit has been reached and returns an error if it has.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to check the storage limit against.
//
// Returns:
//   - (error): An error if the storage limit has been reached, nil otherwise.
func (ts *TempStorage) error(ft *FullText) error {
	// Check if the storage limit has been reached
	if ft.maxLength > 0 {
		if len(ts.data) > ft.maxLength {
			return fmt.Errorf("full-text storage limit reached (%d/%d keys). load cancelled", len(ts.data), ft.maxLength)
		}
	}
	if ft.maxBytes > 0 {
		if cacheSize, err := Utils.Size(ts.data); err != nil {
			return err
		} else if cacheSize > ft.maxBytes {
			return fmt.Errorf("full-text byte-size limit reached (%d/%d bytes). load cancelled", cacheSize, ft.maxBytes)
		}
	}
	return nil
}

// update is a method of the TempStorage struct that updates the TempStorage object with the given words and cache key.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to update.
//   - words ([]string): A slice of strings representing the words to update.
//   - cacheKey (string): A string representing the cache key to update.
//
// Returns:
//   - None.
func (ts *TempStorage) update(ft *FullText, words []string, cacheKey string) {
	// Loop through the words
	for i := 0; i < len(words); i++ {
		// Check if the word is valid
		if len(words[i]) < ft.minWordLength {
			continue
		}
		if temp, ok := ts.data[words[i]]; !ok {
			ts.data[words[i]] = []int{ts.index}
		} else if v, ok := temp.([]int); !ok {
			ts.data[words[i]] = []int{temp.(int), ts.keys[cacheKey]}
		} else {
			if Utils.SliceContains(v, ts.keys[cacheKey]) {
				continue
			}
			ts.data[words[i]] = append(v, ts.keys[cacheKey])
		}
	}
}

// getFTValue is a method of the TempStorage struct that returns the full-text value of a given value.
// Parameters:
//   - value (any): A value to get the full-text value of.
//
// Returns:
//   - (string): The full-text value of the given value.
func (ts *TempStorage) getFTValue(value any) string {
	if wft, ok := value.(WFT); ok {
		return wft.Value
	} else if v := ftFromMap(value); len(v) > 0 {
		return v
	} else {
		return ""
	}
}

// updateKeys is a method of the TempStorage struct that sets the given cache key in the temp storage keys
// Parameters:
//   - cacheKey (string): A string representing the cache key to set.
//
// Returns:
//   - None.
func (ts *TempStorage) updateKeys(cacheKey string) {
	if _, ok := ts.keys[cacheKey]; !ok {
		ts.index++
		ts.indices[ts.index] = cacheKey
		ts.keys[cacheKey] = ts.index
	}
}

// insert is a method of the FullText struct that inserts a value in the full-text cache for the specified key.
// This function is not thread-safe and should only be called from an exported function.
//
// Parameters:
//   - data: A map of maps containing the data to be inserted.
//
// Returns:
//   - An error if the full-text storage limit or byte-size limit is reached.
func (ft *FullText) insert(data *map[string]map[string]any) error {
	// Create a new temp storage
	var ts *TempStorage = NewTempStorage(ft)

	// Loop through the json data
	for cacheKey, cacheValue := range *data {
		for k, v := range cacheValue {
			if ftv := ts.getFTValue(v); len(ftv) == 0 {
				continue
			} else {
				// Set the key in the provided value to the string wft value
				(*data)[cacheKey][k] = ftv

				// Insert the value in the temp storage
				if err := ts.insert(ft, cacheKey, ftv); err != nil {
					return err
				}
			}
		}
	}

	// Iterate over the temp storage and set the values with len 1 to int
	ts.cleanSingleArrays()

	// Set the full-text cache to the temp map
	ts.updateFullText(ft)

	// Return nil for no errors
	return nil
}

// insertWords is a method of the TempStorage struct that inserts data into the temp storage.
// Parameters:
//   - ft (*FullText): A pointer to the FullText object to check the storage limit against.
//   - cacheKey (string): A string representing the cache key to insert.
//   - ftv (string): A string representing the value to insert.
//
// Returns:
//   - (error): An error if the storage limit has been reached, nil otherwise.
func (ts *TempStorage) insert(ft *FullText, cacheKey string, ftv string) error {
	// Set the cache key in the temp storage keys
	ts.updateKeys(cacheKey)

	// Clean the string value
	ftv = strings.TrimSpace(ftv)
	ftv = Utils.RemoveDoubleSpaces(ftv)
	ftv = strings.ToLower(ftv)

	// Loop through the words
	for _, word := range strings.Split(ftv, " ") {
		if len(word) == 0 {
			continue
		} else if len(word) < ft.minWordLength {
			continue
		} else if err := ts.error(ft); err != nil {
			return err
		}

		// Trim the word
		word = Utils.TrimNonAlphaNum(word)
		var words []string = Utils.SplitByAlphaNum(word)

		// Update the temp storage
		ts.update(ft, words, cacheKey)
	}

	// Return no error
	return nil
}
