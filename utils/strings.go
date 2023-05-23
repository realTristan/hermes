package utils

import (
	"strings"
)

// IsAlphaNumChar is a function that checks if a given byte is an alphanumeric character.
// Parameters:
//   - c (byte): The byte to check.
//
// Returns:
//   - bool: true if the byte is an alphanumeric character, false otherwise.
func IsAlphaNumChar(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

// IsAlphaNum is a function that checks if a given string consists entirely of alphanumeric characters.
// Parameters:
//   - s (string): The string to check.
//
// Returns:
//   - bool: true if the string consists entirely of alphanumeric characters, false otherwise.
func IsAlphaNum(s string) bool {
	for _, c := range s {
		if !IsAlphaNumChar(byte(c)) {
			return false
		}
	}
	return true
}

// TrimNonAlphaNum is a function that removes non-alphanumeric characters from the beginning and end of a given string.
// Parameters:
//   - s (string): The string to trim.
//
// Returns:
//   - string: The trimmed string.
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

// SplitByAlphaNum is a function that splits a given string into a slice of substrings by alphanumeric characters.
// Parameters:
//   - s (string): The string to split.
//
// Returns:
//   - []string: A slice of substrings split by alphanumeric characters.
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

// RemoveDoubleSpaces is a function that removes double spaces from a given string and returns the modified string.
// Parameters:
//   - s (string): The string to remove double spaces from.
//
// Returns:
//   - string: The modified string with double spaces removed.
func RemoveDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

// Contains is a function that checks if a given string contains another string as a substring.
// Parameters:
//   - s1 (string): The string to search for the substring.
//   - s2 (string): The substring to search for in the string.
//
// Returns:
//   - bool: true if the substring is found in the string, false otherwise.
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
