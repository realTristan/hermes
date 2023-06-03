package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes/nocache"
)

// Main function
func main() {
	// Define variables
	var (
		// Initialize the full text
		ft, _ = Hermes.InitWithJson("../../../testing/data/data_array.json", 3)

		// Track the start time
		start time.Time = time.Now()

		// Search for a word in the cache
		// @params: query, limit, strict
		res, _ = ft.Search(Hermes.SearchParams{
			Query:  "computer",
			Limit:  100,
			Strict: false,
		})
	)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res), time.Since(start))

	// Search in values with key
	var (
		// Track the start time
		start2 time.Time = time.Now()

		// Search for a word in the cache
		res2, _ = ft.SearchWithKey("CS", "title", 100)
	)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res2), time.Since(start2))
}
