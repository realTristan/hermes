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
		maxWords     int             = 10 // -1 for no limit
		maxSizeBytes int             = -1 // -1 for no limit
		schema       map[string]bool = map[string]bool{
			"name": true,
		}
	)

	// Initialize the FT cache
	if err := cache.InitFT(maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}

	// Track start time
	var startTime time.Time = time.Now()

	// The data for the user_id key
	var data = map[string]interface{}{
		"name":       "tristan",
		"age":        17,
		"expiration": time.Now(),
	}

	// Set the value in the cache
	if err := cache.Set("user_id", data); err != nil {
		fmt.Println(err)
	}

	// Print duration
	fmt.Printf("Set user_id in %s\n", time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Get the user_id value
	var user = cache.Get("user_id")

	// Check if the user is expired
	if user["expiration"].(time.Time).Before(time.Now()) {
		fmt.Println("Expired")
	}

	// Print duration
	fmt.Printf("Got %v in %s\n", user, time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Search for a word in the cache
	var result = cache.FT.SearchOne("tristan", 100, false)

	/* Reset the FT cache
	if err := cache.ResetFT(maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}*/

	// Delete the user_id key
	cache.Delete("user_id")

	// Search for a word in the cache
	result = cache.FT.SearchOne("tristan", 100, false)

	// Print the result
	fmt.Println(result)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
}
