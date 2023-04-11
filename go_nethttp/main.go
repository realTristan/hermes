package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Initialize the caches
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

	// parse the limit
	var limit int = 10
	if _limit := r.URL.Query().Get("limit"); _limit != "" {
		limit, _ = strconv.Atoi(_limit)
	}

	// Track the start time
	var start time.Time = time.Now()

	// Search for a word in the cache
	var courses []map[string]string = Search(query, limit)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(courses), time.Since(start).String())

	// Write the courses to the json response
	var response, _ = json.Marshal(courses)
	w.Write(response)
}

// Search for a word in the cache
func Search(word string, limit int) []map[string]string {
	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Convert the word to lowercase
	word = strings.ToLower(word)

	// Create an array to store the indices that have already
	// been added to the result array
	var alreadyAdded []int = []int{}

	// Loop through the cache keys
	for i := 0; i < len(cacheKeys); i++ {
		switch {
		// Check if the limit has been reached
		case len(result) >= limit:
			return result

		// The word doesn't start with the same letter
		case cacheKeys[i][0] != word[0]:
			continue

		// Check if the key is shorter than the word
		case len(cacheKeys[i]) < len(word):
			continue

		// Check if the key is equal to the word
		case cacheKeys[i] == word:
			// Loop through the indices
			for j := 0; j < len(cache[cacheKeys[i]]); j++ {
				var index int = cache[cacheKeys[i]][j]

				// Else, append the index to the result
				result = append(result, jsonData[index])
			}

			// Return the result
			return result

		// Check if the key contains the word
		case !strings.Contains(cacheKeys[i], word):
			continue

		// Check if the index is already in the result
		case ContainsInt(alreadyAdded, i):
			continue
		}

		// Loop through the indices
		for j := 0; j < len(cache[cacheKeys[i]]); j++ {
			var index int = cache[cacheKeys[i]][j]

			// Else, append the index to the result
			result = append(result, jsonData[index])
			alreadyAdded = append(alreadyAdded, i)
		}
	}

	// Return the result
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
				if !isAlphaNum(word) {
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
func isAlphaNum(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

// Check if an int is in an array
func ContainsInt(array []int, value int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// Check if a string is in an array
func ContainsString(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
