// ////////////////////////////////////////////////////////////////////////////
//
// Run Command: go run .
//
// Host URL: http://localhost:8000/courses?q=computer&limit=100&strict=false
//
// ////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"time"

	hermes "github.com/realTristan/hermes"
)

// Initialize the cache
var cache *hermes.Cache = hermes.InitCache()

func main() {
	// Initialize the full-text cache
	cache.FTInitWithJson("../../testing/data/data_hash.json", -1, -1, 3)

	// Search for a word in the cache
	var startTime time.Time = time.Now()
	var res, _ = cache.Search(hermes.SearchParams{
		Query:  "computer",
		Limit:  100,
		Strict: false,
	})
	var duration time.Duration = time.Since(startTime)
	fmt.Println("Search took", duration)
	fmt.Println("Search results:", len(res))
}
