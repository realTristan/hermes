package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	// Initialize the cache
	var cache *Hermes.Cache = Hermes.InitCache()
	cache.InitFTS()

	// Track start time
	var startTime time.Time = time.Now()

	// Set a value in the cache
	cache.Set("user_id_1", map[string]string{"name": "tristan"})
	cache.Set("user_id_2", map[string]string{"name": "michael"})

	// Print result
	fmt.Printf("Set 2 values in %s\n", time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Search for a word in the cache
	var result, _ = cache.Search("tristan", 100, false)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
}
