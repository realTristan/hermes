package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	// Important Variables
	var (
		cache        *Hermes.Cache   = Hermes.InitCache()
		maxKeys      int             = 10 // -1 for no limit
		maxSizeBytes int             = -1 // -1 for no limit
		keySettings  map[string]bool = map[string]bool{
			"name": true,
		}
	)

	// Initialize the FTS cache
	cache.InitFTS(maxKeys, maxSizeBytes, keySettings)

	// Track start time
	var startTime time.Time = time.Now()

	// Set values in the cache
	if err := cache.Set("user_id", map[string]string{"name": "tristan"}); err != nil {
		fmt.Println(err)
	}

	// Print duration
	fmt.Printf("Set user_id in %s\n", time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Get the user_id value
	var user = cache.Get("user_id")

	// Print duration
	fmt.Printf("Found %v in %s\n", user, time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Search for a word in the cache
	var result, _ = cache.Search("tristan", 100, false)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
}
