package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func delete() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache
	cache.FTInit(-1, -1, 3)

	// Test Delete()
	var data = map[string]any{
		"name": "tristan",
		"age":  17,
	}

	// print cache info
	fmt.Println(cache.InfoForTesting())

	// Set data
	cache.Set("user_id1", data)
	cache.Set("user_id1", data)
	cache.Set("user_id2", data)

	// print cache info
	fmt.Println(cache.InfoForTesting())

	// Delete data
	cache.Delete("user_id1")
	cache.Delete("user_id1")

	// Get data
	fmt.Println(cache.Get("user_id1"))
	fmt.Println(cache.Get("user_id2"))

	// Exists
	fmt.Println(cache.Exists("user_id1"))
	fmt.Println(cache.Exists("user_id2"))

	// Length
	fmt.Println(cache.Length())

	// Values
	fmt.Println(cache.Values())

	// Keys
	fmt.Println(cache.Keys())

	// Print the cache info
	fmt.Println(cache.InfoForTesting())
}
