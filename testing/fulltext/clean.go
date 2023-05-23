package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func clean() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Test CleanFT()
	var data = map[string]any{
		"name": "tristan",
		"age":  17,
	}

	// Set data
	cache.Set("user_id1", data)
	cache.Set("user_id1", data)
	cache.Set("user_id2", data)

	// Initialize the FT cache
	cache.FTInit(-1, -1, 3)

	// Search for a word in the cache
	var result, _ = cache.SearchOneWord(Hermes.SearchParams{
		Query:  "tristan",
		Limit:  100,
		Strict: false,
	})
	fmt.Println(result)

	// Clean
	cache.FTClean()

	// Search for a word in the cache
	result, _ = cache.SearchOneWord(Hermes.SearchParams{
		Query:  "tristan",
		Limit:  100,
		Strict: false,
	})
	fmt.Println(result)
}
