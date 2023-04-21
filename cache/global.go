package cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

// Check if a string consists entirely of alphanumeric characters.
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string and return the modified string.
func removeDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if a string contains another string.
func contains(s1 string, s2 string) bool {
	var (
		s1Len int = len(s1)
		s2Len int = len(s2)
	)
	switch {
	case s1Len == s2Len:
		return s1 == s2
	case s1Len < s2Len:
		return false
	}
	for i := 0; i < s1Len-s2Len; i++ {
		if s1[i] == s2[0] {
			if s1[i:i+s2Len] == s2 {
				return true
			}
		}
	}
	return false
}

// Remove all non-alphanumeric characters from a string.
func removeNonAlphaNum(s string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
}

// Read a json file
func readJson(file string) (map[string]map[string]interface{}, error) {
	var v map[string]map[string]interface{} = map[string]map[string]interface{}{}

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
	}
	return v, nil
}

// Check if an array contains a string.
func containsString(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Gets the real size in memory of a given value.
func size(v interface{}) (int, error) {
	var b *bytes.Buffer = new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}
