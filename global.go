package hermes

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

/*
isAlphaNum function checks if a string consists entirely of alphanumeric characters.
This function returns a boolean indicating whether the string is alphanumeric or not.

Parameters:

	s (string): The string to check for alphanumeric characters.

Returns:

	A boolean value indicating whether the string is alphanumeric or not. If the string is alphanumeric, this function returns true. If not, it returns false.

Example Usage:

	isAlphaNum("abc123") // true
	isAlphaNum("abc-123") // false
*/
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

/*
removeDoubleSpaces function removes double spaces from a string and returns the modified string.

Parameters:

	s (string): The string to remove double spaces from.

Returns:

	A string with all double spaces removed.

Example Usage:

	removeDoubleSpaces("Hello    world") // "Hello world"
	removeDoubleSpaces("This is a test.") // "This is a test."
*/
func removeDoubleSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

/*
contains function checks if a string contains another string as a substring.
This function returns a boolean indicating whether the string contains the substring or not.

Parameters:

	s1 (string): The string to check for a substring.
	s2 (string): The substring to look for in the main string.

Returns:

	A boolean value indicating whether the main string contains the substring or not. If the main string contains the substring,
	this function returns true. If not, it returns false.

Example Usage:

	contains("hello world", "lo w") // true
	contains("hello world", "abc") // false
*/
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

/*
containsIgnoreCase function checks if a string is in another string, ignoring the case of the strings.
This function returns a boolean indicating whether the string contains the substring or not.

Parameters:

	s1 (string): The string to check for a substring.
	s2 (string): The substring to look for in the main string.

Returns:

	A boolean value indicating whether the main string contains the substring or not. If the main string contains the substring,
	this function returns true. If not, it returns false.

Example Usage:

	containsIgnoreCase("hello world", "LO w") // true
	containsIgnoreCase("hello world", "ABC") // false
*/
func containsIgnoreCase(s1 string, s2 string) bool {
	return strings.Contains(strings.ToLower(s1), strings.ToLower(s2))
}

/*
removeNonAlphaNum function removes all non-alphanumeric characters from a string and returns the new string.

Parameters:

	s (string): The string to remove non-alphanumeric characters from.

Returns:

	A string with all non-alphanumeric characters removed.

Example Usage:

	removeNonAlphaNum("Hello World!") // "HelloWorld"
	removeNonAlphaNum("123-456-7890") // "1234567890"
*/
func removeNonAlphaNum(s string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
}

/*
readJson function reads a json file and returns a []map[string]string.

Parameters:

	file (string): The path of the json file to read.

Returns:

	A []map[string]string representing the data in the json file.
	An error if there was a problem reading or unmarshalling the json file.

Example Usage:

	if data, err := readJson("data.json"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}
*/
func readJson(file string) ([]map[string]string, error) {
	var v []map[string]string = []map[string]string{}

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

/*
	Check if an integer is in an array

Parameters:

	array: A slice of integers
	value: The integer value to check for

Returns:

	A boolean indicating whether the integer value is in the slice or not

Example usage:

	intSlice := []int{1, 2, 3, 4, 5}
	if containsInt(intSlice, 3) {
		fmt.Println("3 is in the slice")
	} else {
		fmt.Println("3 is not in the slice")
	}

	// Output: "3 is in the slice"
*/
func containsInt(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
