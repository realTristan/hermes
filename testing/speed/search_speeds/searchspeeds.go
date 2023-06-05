package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	hermes "github.com/realTristan/hermes"
)

// Results: hermes is about 40x faster than a basic search
func main() {
	BasicSearch()
	hermesSearch()
}

// Basic Search
func BasicSearch() {
	// read the json data
	if data, err := readJson("../../data/data_hash.json"); err != nil {
		panic(err)
	} else {
		var average int64 = 0
		for i := 0; i < 100; i++ {
			var startTime = time.Now()
			// print the data
			for _, val := range data {
				for k, v := range val {
					var data map[string]any
					if v, ok := v.(map[string]any); !ok {
						continue
					} else {
						data = v
					}
					// Get the value
					var value string = data["$hermes.value"].(string)

					if strings.Contains(strings.ToLower(value), strings.ToLower("computer")) {
						var _ = k
					}
				}
			}
			average += time.Since(startTime).Nanoseconds()
		}
		var (
			averageNanos  float64 = float64(average) / 100
			averageMillis float64 = averageNanos / 1000000
		)
		fmt.Println("Basic: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
	}
}

// hermes Search
func hermesSearch() {
	// Initialize the cache
	var cache *hermes.Cache = hermes.InitCache()

	// Initialize the FT cache with a json file
	cache.FTInitWithJson("../../data/data_hash.json", -1, -1, 3)
	var (
		average int64 = 0
		total   int   = 0
	)
	for i := 0; i < 100; i++ {
		// Track the start time
		var (
			start time.Time = time.Now()

			// Search for a word in the cache
			res, _ = cache.Search(hermes.SearchParams{
				Query:  "computer science",
				Limit:  100,
				Strict: false,
				Schema: map[string]bool{
					"name":        true,
					"description": true,
					"title":       true,
				},
			})
		)

		// Print the duration
		average += time.Since(start).Nanoseconds()
		total += len(res)
	}
	var (
		averageNanos  float64 = float64(average) / 100
		averageMillis float64 = averageNanos / 1000000
	)
	fmt.Println("hermes: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
	fmt.Println("hermes: Results: ", total)
}

// Read a json file
func readJson(file string) (map[string]map[string]any, error) {
	var v map[string]map[string]any = map[string]map[string]any{}

	// Read the json data
	if data, err := os.ReadFile(file); err != nil {
		return nil, err
	} else if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return v, nil
}
