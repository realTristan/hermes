package main

import (
	"fmt"
	"time"

	hermes "github.com/realTristan/hermes"
)

func main() {
	// Cache
	var cache *hermes.Cache = hermes.InitCache()

	/* Initialize the FT cache
	if err := cache.InitFTWithJson("../../data/data_hash.json", maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}
	*/
	// cache.Info()

	// Initialize the FT cache
	if err := cache.FTInit(-1, -1, 3); err != nil {
		fmt.Println(err)
	}

	// The data for the user_id and user_id2 key
	var data = map[string]any{
		"name": cache.WithFT("tristan1"),
		"age":  17,
	}
	var data2 = map[string]any{
		"name": map[string]any{
			"$hermes.full_text": true,
			"value":             "tristan2",
		},
		"age": 17,
	}

	// Set the value in the cache
	duration("Set", func() {
		if err := cache.Set("user_id", data); err != nil {
			fmt.Println(err)
		}

		if err := cache.Set("user_id2", data2); err != nil {
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
		var result, _ = cache.SearchOneWord(hermes.SearchParams{
			Query:  "tristan",
			Limit:  100,
			Strict: false,
		})
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
		var result, _ = cache.SearchOneWord(hermes.SearchParams{
			Query:  "tristan",
			Limit:  100,
			Strict: false,
		})
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
