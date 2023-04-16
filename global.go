package hermes

import (
	"regexp"
	"strings"
)

// Check if a string is all alphabetic
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Remove double spaces from a string
func removeDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if an int is in an array
func containsInt(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string contains a substring
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

// Check if a string is in another string ignore case
func containsIgnoreCase(s1 string, s2 string) bool {
	return strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
}

// Remove non alphanumeric characters from a string
func removeNonAlphaNum(s string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
}
