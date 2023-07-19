package main

import (
	"fmt"

	hermes "github.com/realTristan/hermes"
)

func insert_key_merge() {
	var cache *hermes.Cache = hermes.InitCache()

	// Initialize the FT cache
	cache.FTInit(-1, -1, 3)
	cache.Set("user_id1", map[string]any{
		"name": cache.WithFT("tristan"),
		"age":  17,
	})
	cache.Set("user_id2", map[string]any{
		"name": cache.WithFT("tris"),
		"age":  17,
	})

	// Search for tris
	var result, _ = cache.SearchOneWord(hermes.SearchParams{
		Query:  "tris",
		Limit:  100,
		Strict: false,
	})
	fmt.Println(result)

	// print cache info
	fmt.Println(cache.InfoForTesting())
}
