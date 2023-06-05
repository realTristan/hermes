package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	hermes "github.com/realTristan/hermes/nocache"
)

// Results: hermes is about 40x faster than a basic search
func main() {
	BasicSearch()
	hermesSearch()
}

// Basic Search
func BasicSearch() {
	// read the json data
	if data, err := readJson("../../../testing/data/data_array.json"); err != nil {
		panic(err)
	} else {
		var average int64 = 0
		for i := 0; i < 100; i++ {
			var startTime = time.Now()
			// Iterate over the data array
			for _, val := range data {
				// Iterate over the map
				for k, v := range val {
					var data map[string]interface{}
					if v, ok := v.(map[string]interface{}); !ok {
						continue
					} else {
						data = v
					}
					// Get the value
					var value string = data["$hermes.value"].(string)

					// Check if the value contains the search term
					if strings.Contains(strings.ToLower(value), strings.ToLower("computer")) {
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

// hermes Search
func hermesSearch() {
	// Initialize the cache
	var cache, err = hermes.InitWithJson("../../../data/data_array.json", 3)
	if err != nil {
		panic(err)
	}

	var average int64 = 0
	for i := 0; i < 100; i++ {
		// Track the start time
		var start time.Time = time.Now()

		// Search for a word in the cache
		cache.Search(hermes.SearchParams{
			Query:  "computer",
			Limit:  100,
			Strict: false,
		})

		// Print the duration
		average += time.Since(start).Nanoseconds()
	}

	var averageNanos float32 = float32(average) / 100
	var averageMillis float32 = averageNanos / 1000000
	fmt.Println("\nhermes: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
}

// Read a json file
func readJson(file string) ([]map[string]interface{}, error) {
	var v []map[string]interface{} = []map[string]interface{}{}

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
