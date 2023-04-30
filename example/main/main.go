package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

// Main function
func main() {
	// Create a schema. These are the fields that will be searched.
	var schema = map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	}

	// Define variables
	var (
		// Initialize the full text
		ft, _ = Hermes.InitWithJson("../../data/data_array.json", schema)

		// Track the start time
		start time.Time = time.Now()

		// Search for a word in the cache
		// @params: query, limit, strict
		res, _ = ft.Search("computer", 100, false, schema)
	)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res), time.Since(start))
}
