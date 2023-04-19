package cache

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
readJson function reads a json file and returns a map[string]map[string]interface{}.

Parameters:

	file (string): The path of the json file to read.

Returns:

	A map[string]map[string]interface{} representing the data in the json file.
	An error if there was a problem reading or unmarshalling the json file.

Example Usage:

	if data, err := readJson("data.json"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}
*/
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

/*
Checks whether a given string is already present in the specified string array.

Parameters:
  - array: the string array to be searched
  - value: the string value to search for

Returns:
  - bool: true if the array contains the value, false otherwise

Example usage:

	array := []string{"apple", "banana", "orange"}
	value := "banana"
	result := containsString(array, value)
	fmt.Println(result) // Output: true
*/
func containsString(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
