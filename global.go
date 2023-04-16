package hermes

import (
	"regexp"
	"strings"
)

// Check if a string is all alphabetic
func _IsAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string
func _RemoveDoubleSpaces(s string) string {
	for _Contains(s, "  ", 2) {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if an int is in an array
func _ContainsInt(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string is in an array
func _ContainsString(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string contains a substring
func _Contains(s string, toFind string, toFindLength int) bool {
	for i := 0; i < len(s)-toFindLength-1; i++ {
		if s[i] == toFind[0] {
			if s[i:i+toFindLength] == toFind {
				return true
			}
		}
	}
	return false
}

// Check if a string contains another string (case insensitive)
func ContainsIgnoreCase(s1 string, s2 string) bool {
	return _Contains(strings.ToLower(s1), strings.ToLower(s2), len(s2))
}
