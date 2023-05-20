package utils

import (
	"strings"
)

// Check if a char is alphanumeric.
func IsAlphaNumChar(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

// Check  if a string consists entirely of alphanumeric characters.
func IsAlphaNum(s string) bool {
	for _, c := range s {
		if !IsAlphaNumChar(byte(c)) {
			return false
		}
	}
	return true
}

// Trim non alphanumeric characters from the beginning and end of a string.
func TrimNonAlphaNum(s string) string {
	if len(s) == 0 {
		return s
	}
	for !IsAlphaNumChar(s[0]) {
		if len(s) < 2 {
			return ""
		}
		s = s[1:]
	}
	for !IsAlphaNumChar(s[len(s)-1]) {
		if len(s) < 2 {
			return ""
		}
		s = s[:len(s)-1]
	}
	return s
}

// Split by alphanumeric characters.
func SplitByAlphaNum(s string) []string {
	var (
		word  string   = ""
		words []string = make([]string, 0)
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '.' || IsAlphaNumChar(s[i]) {
			word += string(s[i])
		} else {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	return words
}

// Remove double spaces from a string and return the modified string.
func RemoveDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Check if a string contains another string as a substring.
func Contains(s1 string, s2 string) bool {
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
