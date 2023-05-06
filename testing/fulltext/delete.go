package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func delete() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache
	cache.FTInit(-1, -1, map[string]bool{
		"name": true,
	})

	// Test Delete()
	var data = map[string]interface{}{
		"name": "tristan",
		"age":  17,
	}

	// print cache info
	fmt.Println(cache.Info())

	// Set data
	cache.Set("user_id1", data, true)
	cache.Set("user_id1", data, true)
	cache.Set("user_id2", data, true)

	// print cache info
	fmt.Println(cache.Info())

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
	fmt.Println(cache.Info())
}
