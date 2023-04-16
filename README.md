# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

## Install
```
go get github.com/realTristan/Hermes
```

# About
## Storing Data
Hermes works by iterating over the items in the data.json file, and then iterates over the keys and values of the items and splits the value into different words. It then stores the indices for all of the items that contain those words in a dictionary.

## Accessing Data
When searching for a word, Hermes will return a list of indices for all of the items that contain that word. It checks whether the key in the cache dictionary contains the provided word, instead of just accessing it so that short forms for words can be used.

## How to improve the speed
1. Create a history cache

## Benchmarks
`Dataset Keys: 4,115`

`Dataset Total Words: 208,092`

`Dataset Map Size: 33,048 bytes`

`?q=computer&limit=100&strict=false: 32.5µs`

`?q=computer&limit=100&strict=true: 9.102µs`


# Example
```go
package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes"
)

func main() {
	// Initialize the cache
	var cache *Hermes.Cache = Hermes.InitCache()
	cache.InitFTS()

	// Track start time
	var startTime time.Time = time.Now()

	// Set a value in the cache
	cache.Set("user_id_1", map[string]string{"name": "tristan"})
	cache.Set("user_id_2", map[string]string{"name": "michael"})

	// Print result
	fmt.Printf("Set 2 values in %s\n", time.Since(startTime))

	// Track start time
	startTime = time.Now()

	// Search for a word in the cache
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
