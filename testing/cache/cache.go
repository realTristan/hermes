package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	var cache *Hermes.Cache = Hermes.InitCache()
	cache.FTInit(3, -1, -1)

	// Test Set, Get, Delete, Clean, Length, Values, Keys, and Exists
	var data = map[string]any{
		"name": Hermes.WithFT("Tristan"),
		"age":  17,
	}

	// Set data
	cache.Set("user_id1", data)
	cache.Set("user_id1", data)
	cache.Set("user_id2", data)
	cache.Set("user_id3", data)

	// Get data
	fmt.Println(cache.Get("user_id1"))
	fmt.Println(cache.Get("user_id2"))
	fmt.Println(cache.Get("user_id3"))

	// Exists
	fmt.Println(cache.Exists("user_id1"))
	fmt.Println(cache.Exists("user_id2"))

	// Length
	fmt.Println(cache.Length())

	// Values
	fmt.Println(cache.Values())

	// Keys
	fmt.Println(cache.Keys())

	// Delete data
	cache.Delete("user_id1")

	// Exists
	fmt.Println(cache.Exists("user_id1"))
	fmt.Println(cache.Exists("user_id2"))

	// Get data
	fmt.Println(cache.Get("user_id1"))
	fmt.Println(cache.Get("user_id2"))
	fmt.Println(cache.Get("user_id3"))

	// Clean data
	cache.Clean()

	// Get data
	fmt.Println(cache.Get("user_id1"))
	fmt.Println(cache.Get("user_id2"))
	fmt.Println(cache.Get("user_id3"))

	// Length
	fmt.Println(cache.Length())

	// Values
	fmt.Println(cache.Values())

	// Keys
	fmt.Println(cache.Keys())
}
