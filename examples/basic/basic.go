// ////////////////////////////////////////////////////////////////////////////
//
// Run Command: go run .
//
// Host URL: http://localhost:8000/courses?q=computer&limit=100&strict=false
//
// ////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

// Initialize the cache
var cache *Hermes.Cache = Hermes.InitCache()

// Main function
func main() {
	// Initialize the FT cache with a json file
	cache.FTInitWithJson("../../testing/data/data_hash.json", -1, -1, 3)

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
	var query string = "CS"
	if _query := r.URL.Query().Get("q"); _query != "" {
		query = _query
	}

	// Get the limit parameter
	var limit int = 100
	if _limit := r.URL.Query().Get("limit"); _limit != "" {
		limit, _ = strconv.Atoi(_limit)
	}

	// Get the strict parameter
	var strict bool = false
	if _strict := r.URL.Query().Get("strict"); _strict != "" {
		strict, _ = strconv.ParseBool(_strict)
	}

	// Track the start time
	var start time.Time = time.Now()

	// Search for a word in the cache
	// Make sure the show which keys you do want to search through,
	// and which ones you don't
	var res, _ = cache.Search(Hermes.SearchParams{
		Query:  query,
		Limit:  limit,
		Strict: strict,
	})

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res), time.Since(start))

	// Write the courses to the json response
	var response, _ = json.Marshal(res)
	w.Write(response)
}
