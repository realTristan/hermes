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
				switch {
				case len(word) <= 1:
					continue
				case !Utils.IsAlphaNum(word):
					word = Utils.RemoveNonAlphaNum(word)
				}
				if _, ok := ft.wordCache[word]; !ok {
					ft.wordCache[word] = []int{i}
					ft.words = append(ft.words, word)
					continue
				}
				if Utils.ContainsInt(ft.wordCache[word], i) {
					continue
				}
				ft.wordCache[word] = append(ft.wordCache[word], i)
			}

			// Set the value
			data[i][key] = strvNormal
		}
	}
	return nil
}
