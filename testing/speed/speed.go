package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

// Results: Hermes is about 40x faster than a basic search
func main() {
	BasicSearch()
	HermesSearch()
}

// Basic Search
func BasicSearch() {
	// read the json data
	if data, err := readJson("../../data/data_array.json"); err != nil {
		panic(err)
	} else {
		var average int64 = 0
		for i := 0; i < 100; i++ {
			var startTime = time.Now()
			// Iterate over the data array
			for _, val := range data {
				// Iterate over the map
				for k, v := range val {
					// Check if the value contains the search term
					if strings.Contains(strings.ToLower(v), strings.ToLower("computer")) {
						var _ = k
					}
				}
			}
			average += time.Since(startTime).Nanoseconds()
		}
		var averageNanos float64 = float64(average) / 100
		var averageMillis float64 = averageNanos / 1000000
		fmt.Println("\nBasic: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
	}
}

// Hermes Search
func HermesSearch() {
	// Initialize the cache
	var cache, err = Hermes.InitWithJson("../../data/data_array.json", map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"pre_requisites": false,
		"title":          false,
		"description":    true,
		"name":           true,
	})
	if err != nil {
		panic(err)
	}

	var average int64 = 0
	for i := 0; i < 100; i++ {
		// Track the start time
		var start time.Time = time.Now()

		// Search for a word in the cache
		cache.SearchOne("computer", 100, false)

		// Print the duration
		average += time.Since(start).Nanoseconds()
	}

	var averageNanos float32 = float32(average) / 100
	var averageMillis float32 = averageNanos / 1000000
	fmt.Println("\nHermes: Average time is: ", averageNanos, "ns or", averageMillis, "ms\n")
}

// Read a json file
func readJson(file string) ([]map[string]string, error) {
	var v []map[string]string = []map[string]string{}

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
	}
	return v, nil
}
