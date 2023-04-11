package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Initialize the cache from the hermes.go file
var cache *Cache = InitCache("data.json")

// Main function
func main() {
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
	var courses []map[string]string = cache.Search(query, limit)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(courses), time.Since(start))

	// Write the courses to the json response
	var response, _ = json.Marshal(courses)
	w.Write(response)
}
