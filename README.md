# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

## Install
```
go get github.com/realTristan/Hermes
```

## Benchmarks
```
Dataset Array Entries: 4,115

Dataset Total Words: 208,092

Dataset Map Size: ≈2.3MB

?q=computer&limit=100&strict=false: 52.5µs

?q=computer&limit=100&strict=true: 12.102µs
```

# Remarks
1. The full-text-search from /cache is significantly slower than the base FTS. Why? Because the FTS in /cache requires more memory, keys, and utilizes a map, instead of a slice to store data.
2. If you want to use a cache along with the full text-search algorithm, then import the files from /cache. To setup a cache, check out /cache/example or /cache/testing. 
3. If you want to use only the full-text-search features, then just import hermes and load it using a .json file. (as shown in /example)
4. For small to medium-sized datasets (like the ones I used in /data), Hermes works great. Although, as the words in the dataset increases, the full-text-search cache will take up significantly more memory. I recommended setting a cache limit and/or a cache keys limit.

# Example of /cache
```go
package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	// Important Variables
	var (
		cache        *Hermes.Cache   = Hermes.InitCache()
		maxWords     int             = 10 // -1 for no limit
		maxSizeBytes int             = -1 // -1 for no limit

		// The keys you want to search through in the FTS
		schema       map[string]bool = map[string]bool{
			"name": true,
		}
	)

	// Initialize the FT cache
	if err := cache.InitFT(maxWords, maxSizeBytes, schema); err != nil {
		fmt.Println(err)
	}
	
	// The data for the user_id key
	var data = map[string]interface{}{
		"name":       "tristan",
		"age":        17,
		"expiration": time.Now(),
	}

	// Set the value in the cache
	if err := cache.Set("user_id", data); err != nil {
		fmt.Println(err)
	}

	// Get the user_id value
	var user = cache.Get("user_id")

	// Check if the user is expired
	if user["expiration"].(time.Time).Before(time.Now()) {
		fmt.Println("Expired")
	}

	// Search for a word in the cache
	var (
		startTime time.Time = time.Now()
		result = cache.FT.SearchOne("tristan", 100, false)
	)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
}
```

# Example of Json Full Text Search
```go
// /////////////////////////////////////////////////////////////////////////////
//
// Run Command: go run .
//
// Host URL: http://localhost:8000/courses?q=computer&limit=100&strict=false
//
// /////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

// Global full text variable
var ft *Hermes.FullText

// Main function
func main() {
	ft, _ = Hermes.InitJson("../data/data_array.json", map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	})

	// Print host
	fmt.Println(" >> Listening on: http://localhost:8000/")

	// Listen and serve on port 8000
	http.HandleFunc("/courses", Handler)
	http.ListenAndServe(":8000", nil)
}

// Handle the incoming http request
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the query parameter
	var query string = "CS"
	if _query := r.URL.Query().Get("q"); _query != "" {
		query = _query
	}

	// Get the limit parameter
	var limit int = 100
	if _limit := r.URL.Query().Get("limit"); _limit != "" {
		limit, _ = strconv.Atoi(_limit)
	}

	// Get the strict parameter
	var strict bool = false
	if _strict := r.URL.Query().Get("strict"); _strict != "" {
		strict, _ = strconv.ParseBool(_strict)
	}

	// Track the start time
	var start time.Time = time.Now()

	// Search for a word in the cache
	// Make sure the show which keys you do want to search through,
	// and which ones you don't
	var res []map[string]string = ft.SearchWithSpaces(query, limit, strict, map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	})

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res), time.Since(start))

	// Write the courses to the json response
	var response, _ = json.Marshal(res)
	w.Write(response)
}
```

# License
MIT License

Copyright (c) 2023 Tristan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
