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
	for _Contains(s, "  ") {
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

// Check if a string contains a substring
func _Contains(s1 string, s2 string) bool {
	switch {
	case s1 == s2:
		return true
	case len(s1) < len(s2):
		return false
	}
	for i := 0; i < len(s1)-len(s2); i++ {
		if s1[i] == s2[0] {
			if s1[i:i+len(s2)] == s2 {
				return true
			}
		}
	}
	return false
}

// Check if a string contains another string (case insensitive)
func _ContainsIgnoreCase(s1 string, s2 string) bool {
	return _Contains(strings.ToLower(s1), strings.ToLower(s2))
}
