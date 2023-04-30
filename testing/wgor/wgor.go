package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	// Initialize the cache
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache with a json file
	cache.FTInitWithJson("../../../data/data_indices.json", -1, -1, map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	})

	// Test the WGOR function
	var start = time.Now()
	var data = cache.WGOR_SearchOneWord_Example()
	var end = time.Since(start)

	// Print the results
	fmt.Println("Hermes: Average time is: ", end.Nanoseconds(), "ns or", end.Milliseconds(), "ms")
	fmt.Println("Hermes: Total results: ", len(data))
}
