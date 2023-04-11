package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Initialize the cache
var cache map[string][]int = make(map[string][]int)
var cacheKeys []string = []string{}

// Initialize the json data
var jsonData []map[string]string = []map[string]string{}

// Main function
func main() {
	// Load the cache
	LoadCache()

	// Print host
	fmt.Println(" >> Listening on: http://localhost:8000/")

	// Listen and serve on port 8000
	http.HandleFunc("/courses", Handler)
	http.ListenAndServe(":8000", nil)
}

// Handle the incoming http request
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the query parameter
	var query string = r.URL.Query().Get("q")

	// Track the start time
	var start time.Time = time.Now()

	// Search for a word in the cache
	var courses []map[string]string = Search(query)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(courses), time.Since(start).String())

	// Write the courses to the json response
	var response, _ = json.Marshal(courses)
	w.Write(response)
}

// Search for a word in the cache
func Search(word string) []map[string]string {
	word = strings.ToLower(word)
	var result []map[string]string = []map[string]string{}
	var alreadyAdded []int = []int{}

	// Loop through the cache keys
	for i := 0; i < len(cacheKeys); i++ {
		// Check if the key contains the word
		if !strings.Contains(cacheKeys[i], word) {
			continue
		}

		// Loop through the indices
		for j := 0; j < len(cache[cacheKeys[i]]); j++ {
			var index int = cache[cacheKeys[i]][j]

			// Check if the index is already in the result
			if ContainsInt(alreadyAdded, index) {
				continue
			}

			// Else, append the index to the result
			result = append(result, jsonData[index])
			alreadyAdded = append(alreadyAdded, index)
		}
	}
	return result
}

// Load the cache
func LoadCache() {
	// Read the json file
	var data, _ = os.ReadFile("../data.json")
	json.Unmarshal(data, &jsonData)

	// Loop through the json data
	for i, item := range jsonData {
		for _, runeValue := range item {
			// Convert the rune to a string
			var value string = string(runeValue)

			// Remove double spaces
			for strings.Contains(value, "  ") {
				value = strings.Replace(value, "  ", " ", -1)
			}

			// Remove spaces from front and back of the string and convert to lowercase
			value = strings.ToLower(strings.TrimSpace(value))

			// Split the string into an array
			var array []string = strings.Split(value, " ")

			// Loop through the array
			for _, word := range array {
				// Check if the word is not all alphabetic
				if !isAlpha(word) {
					continue
				}

				// Check if the key exists in the cache
				if _, ok := cache[word]; !ok {
					cache[word] = []int{}
				}

				// Check if the index is already in the cache
				if ContainsInt(cache[word], i) {
					continue
				}

				// Update the cache keys
				if !ContainsString(cacheKeys, word) {
					cacheKeys = append(cacheKeys, word)
				}
				cache[word] = append(cache[word], i)
			}
		}
	}
}

// Check if a string is all alphabetic
func isAlpha(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Check if an int is in an array
func ContainsInt(array []int, value int) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

// Check if a string is in an array
func ContainsString(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
