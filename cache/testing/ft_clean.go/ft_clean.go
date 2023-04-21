package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes/cache"
)

func main() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Test CleanFT()
	var data = map[string]interface{}{
		"name": "tristan",
		"age":  17,
	}

	// Set data
	cache.Set("user_id1", data, true)
	cache.Set("user_id1", data, true)
	cache.Set("user_id2", data, true)

	// Initialize the FT cache
	cache.FTInit(-1, -1, map[string]bool{
		"name": true,
	})

	// Search for a word in the cache
	var result, _ = cache.SearchOne("tristan", 100, false)
	fmt.Println(result)

	// Clean
	cache.FTClean()

	// Search for a word in the cache
	result, _ = cache.SearchOne("tristan", 100, false)
	fmt.Println(result)
}
