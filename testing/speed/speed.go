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
	if data, err := readJson("../../data/data_hash.json"); err != nil {
		panic(err)
	} else {
		var average int64 = 0
		for i := 0; i < 100; i++ {
			var startTime = time.Now()
			// print the data
			for _, val := range data {
				for k, v := range val {
					if strings.Contains(strings.ToLower(v.(string)), strings.ToLower("computer")) {
						var _ = k
					}
				}
			}
			average += time.Since(startTime).Nanoseconds()
		}
		var averageNanos float64 = float64(average) / 100
		var averageMillis float64 = averageNanos / 1000000
		fmt.Println("Basic: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
	}
}

// Hermes Search
func HermesSearch() {
	// Initialize the cache
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache with a json file
	cache.FTInitWithJson("../../data/data_hash.json", -1, -1, map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	})
	var average int64 = 0
	var total int = 0
	for i := 0; i < 100; i++ {
		// Track the start time
		var start time.Time = time.Now()

		// Search for a word in the cache
		var res, _ = cache.Search("computer science", 100, false, map[string]bool{
			"id":             false,
			"components":     false,
			"units":          false,
			"description":    true,
			"name":           true,
			"pre_requisites": true,
			"title":          true,
		})
		total += len(res)

		// Print the duration
		average += time.Since(start).Nanoseconds()
	}
	var averageNanos float64 = float64(average) / 100
	var averageMillis float64 = averageNanos / 1000000
	fmt.Println("Hermes: Average time is: ", averageNanos, "ns or", averageMillis, "ms")
	fmt.Println("Hermes: Results: ", total)
}

// Read a json file
func readJson(file string) (map[string]map[string]interface{}, error) {
	var v map[string]map[string]interface{} = map[string]map[string]interface{}{}

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
