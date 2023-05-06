package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func add() {
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
	cache.InfoForTesting()

	// Set data
	cache.Set("user_id1", data, false)
	cache.Set("user_id1", data, true)
	cache.Set("user_id2", data, false)

	cache.FTAdd("user_id1")
	cache.FTAdd("user_id1")
	cache.FTAdd("user_id2")

	// print cache info
	fmt.Println(cache.InfoForTesting())
}
