# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

## Install
```
go get github.com/realTristan/Hermes
```

## Benchmarks
`Dataset Array Entries: 4,115`

`Dataset Total Words: 208,092`

`Dataset Map Size: ≈2.3MB`

`?q=computer&limit=100&strict=false: 52.5µs`

`?q=computer&limit=100&strict=true: 12.102µs`

# Remarks
For small to medium-sized datasets (like the one I used in /examples/data.json), Hermes works great. Although, as the words in the dataset increases, the full-text-search cache will take up significantly more memory. I recommended setting a cache limit and/or a cache keys limit. If you prefer not having to worry about memory, and would prefer slower full-text-searches, I recommend BetterCache. ```https://github.com/realTristan/bettercache```

# Example
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
		maxKeys      int             = 10 // -1 for no limit
		maxSizeBytes int             = -1 // -1 for no limit
		schema  map[string]bool = map[string]bool{
			"name": true,
		}
	)

	// Initialize the FTS cache
	cache.InitFTS(maxKeys, maxSizeBytes, schema)

	// Set values in the cache
	var data = map[string]string{"name": "tristan"}
	if err := cache.Set("user_id", data); err != nil {
		fmt.Println(err)
	}

	// Get "user_id" from the cache
	var user = cache.Get("user_id")

	// Search for the name "tristan" in the cache
	// @param query: "tristan
	// @param limit: 100
	// @param strict: false
	var result, _ = cache.Search("tristan", 100, false)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
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
