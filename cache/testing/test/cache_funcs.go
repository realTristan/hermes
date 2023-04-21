package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes/cache"
)

func main() {
	// Important Variables
	var (
		cache        *Hermes.Cache   = Hermes.InitCache()
		maxWords     int             = -1 // -1 for no limit
		maxSizeBytes int             = -1 // -1 for no limit
		schema       map[string]bool = map[string]bool{
			"name":       true,
			"age":        false,
			"expiration": false,
		}
	)

	/* Initialize the FT cache
	if err := cache.InitFTWithJson("../../../data/data_indices.json", maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}
	*/

	// Initialize the FT cache
	if err := cache.InitFT(maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}

	// The data for the user_id key
	var data = map[string]interface{}{
		"name": "tristan",
		"age":  17,
	}

	// Set the value in the cache
	duration("Set", func() {
		if err := cache.Set("user_id", data); err != nil {
			fmt.Println(err)
		}

		if err := cache.Set("user_id", data); err != nil {
			fmt.Println(err)
		}
	})

	// Get the user_id value
	duration("Get", func() {
		var user = cache.Get("user_id")
		fmt.Println(user)
	})

	// Search for a word in the cache
	duration("Search", func() {
		var result, _ = cache.SearchOne("tristan", 100, false)
		fmt.Println(result)
	})

	// Print all the cache info
	//cache.Info()

	/* Reset the FT cache
	if err := cache.ResetFT(maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}*/

	// Delete the user_id key
	duration("Delete", func() {
		cache.Delete("user_id")
	})

	// Search for a word in the cache
	duration("Search", func() {
		var result, _ = cache.SearchOne("tristan", 100, false)
		fmt.Println(result)
	})

	// Print all the cache info
	//cache.Info()
}

// Track the duration of a function
func duration(key string, f func()) {
	var start time.Time = time.Now()
	f()
	fmt.Printf("\nExecution Duration for %s: %s\n", key, time.Since(start))
}
