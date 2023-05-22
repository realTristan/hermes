package nocache

import (
	"strings"

	Utils "github.com/realTristan/Hermes/utils"
)

// Insert data into the full-text cache.
// This function is not thread-safe, and should only be called from
// an exported function.
func (ft *FullText) insert(data []map[string]interface{}) error {
	// Loop through the data
	for i, item := range data {
		// Loop through the map
		for key, value := range item {
			// Get the string value
			var (
				strvNormal string
				strv       string
			)
			if _strv := ftFromMap(value); len(_strv) > 0 {
				strv = _strv
				strvNormal = _strv
			} else {
				continue
			}

			// Clean the value
			strv = strings.TrimSpace(strv)
			strv = Utils.RemoveDoubleSpaces(strv)
			strv = strings.ToLower(strv)

			// Loop through the words
			for _, word := range strings.Split(strv, " ") {
				if len(word) == 0 {
					continue
				}

				// Trim the word
				word = Utils.TrimNonAlphaNum(word)
				var words []string = Utils.SplitByAlphaNum(word)

				// Loop through the words
				for j := 0; j < len(words); j++ {
					if len(words[j]) <= 3 {
						continue
					}
					if temp, ok := ft.storage[words[j]]; !ok {
						ft.storage[words[j]] = []int{i}
						ft.words = append(ft.words, words[j])
					} else {
						if indices, ok := temp.([]int); !ok {
							ft.storage[words[j]] = []int{temp.(int), i}
						} else {
							if Utils.SliceContains(indices, i) {
								continue
							}
							ft.storage[words[j]] = append(indices, i)
						}
					}
				}
			}

			// Iterate over the temp storage and set the values with len 1 to int
			for k, v := range ft.storage {
				if v, ok := v.([]int); ok && len(v) == 1 {
					ft.storage[k] = v[0]
				}
			}

			// Set the value
			data[i][key] = strvNormal
		}
	}
	return nil
}
