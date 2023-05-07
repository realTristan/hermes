# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

# Import
```
Go: go get github.com/realTristan/Hermes
Python: See /wrappers/python
```

# Example of NoCache Full-Text-Search
If you want to use only the full-text-search features, then just import hermes and load it using a .json file. (as shown in /example). Note: For small to medium-sized datasets (like the ones I used in /data), Hermes works great. Although, as the words in the dataset increases, the full-text-search cache will take up significantly more memory. I recommended setting a cache limit and/or a cache keys limit.
```
Dataset Array Entries: 4,115

Dataset Total Words: 208,092

Dataset Map Size: ≈ 2.3MB

?q=computer&limit=100&strict=false: 36.5µs

?q=computer&limit=100&strict=true: 12.102µs
```

```go
package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes/nocache"
)

// Main function
func main() {
	// Create a schema. These are the fields that will be searched.
	var schema = map[string]bool{
		"id":             false,
		"components":     false,
		"units":          false,
		"description":    true,
		"name":           true,
		"pre_requisites": true,
		"title":          true,
	}

	// Define variables
	var (
		// Initialize the full text
		ft, _ = Hermes.InitWithJson("data.json", schema)

		// Track the start time
		start time.Time = time.Now()

		// Search for a word in the cache
		// @params: query, limit, strict
		res, _ = ft.Search("computer", 100, false, schema)
	)

	// Print the duration
	fmt.Printf("\nFound %v results in %v", len(res), time.Since(start))
}
```

# Example of Cache Full-Text-Search
The full-text-search from /cache is significantly slower than the nocache FTS. Why? Because the FTS in /cache requires more memory, keys, and utilizes a map instead of a slice to store data. If you want to use a cache along with the full text-search algorithm, import the files from /cache. To setup a cache, check out /cache/example or /cache/testing. 

```
Dataset Map Entries: 4,115

Dataset Total Words: 208,092

Dataset Map Size: ≈ 2.3MB

?q=computer&limit=100&strict=false: 263.7µs

?q=computer&limit=100&strict=true: 40.84µs
```

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
		maxBytes int             = -1 // -1 for no limit

		// The keys you want to search through in the FTS
		schema       map[string]bool = map[string]bool{
			"name":       true,
			"age":        false,
			"expiration": false,
		}
	)

	// Initialize the FT cache
	if err := cache.FTInit(maxWords, maxBytes, schema); err != nil {
		fmt.Println(err)
	}
	
	// The data for the user_id key
	var data = map[string]interface{}{
		"name":       "tristan",
		"age":        17,
		"expiration": time.Now(),
	}

	// Set the value in the cache
	if err := cache.Set("user_id", data, true); err != nil {
		fmt.Println(err)
	}

	// Search for a word in the cache
	var (
		startTime time.Time = time.Now()
		result, _ = cache.Search("tristan", 100, false, schema)
	)

	// Print result
	fmt.Printf("Found %d results in %s\n", len(result), time.Since(startTime))
	fmt.Println(result)
}
```

# API and CLI App
## Docker Container
```
$ git clone https://github.com/realTristan/Hermes.git
$ docker build -t hermes .
$ docker run -p 6000:6000 hermes
```

## App Usage
```
hermes serve -p 6000
```

## Install
```
MacOS:
  $ curl "https://github.com/realTristan/Hermes/raw/master/cli/hermes" -o /usr/local/bin/hermes
  
Windows:
  $ curl "https://github.com/realTristan/Hermes/raw/master/cli/hermes.exe" -o C:\hermes.exe
  $ set PATH=%PATH%;C:\hermes.exe
```

# API
## Cache

[POST /set](https://github.com/realTristan/Hermes/blob/master/server/handlers/set.go)
```
About
➤ Set a value in the cache with the corresponding key.

URL Parameters
➤ key: string
➤ value: base64(map[string]map[string]interface{})
➤ ft: bool
```

[DELETE /delete](https://github.com/realTristan/Hermes/blob/master/server/handlers/delete.go)
```
About
➤ Delete the provided key, and the data correlated to it from the cache.

URL Parameters
➤ key: string
```

[GET /get](https://github.com/realTristan/Hermes/blob/master/server/handlers/get.go)
```
About
➤ Get data from the cache using a key.

URL Parameters
➤ key: string
```

[GET /keys](https://github.com/realTristan/Hermes/blob/master/server/handlers/keys.go)
```
About
➤ Get all of the keys in the cache.

URL Parameters
➤ None
```

[GET /values](https://github.com/realTristan/Hermes/blob/master/server/handlers/values.go)
```
About
➤ Get all of the values in the cache.

URL Parameters
➤ None
```

[GET /length](https://github.com/realTristan/Hermes/blob/master/server/handlers/length.go)
```
About
➤ Get the amount of keys stored in the cache.

URL Parameters
➤ None
```

[POST /clean](https://github.com/realTristan/Hermes/blob/master/server/handlers/clean.go)
```
About
➤ Clean all the data in the cache, and full-text cache.

URL Parameters
➤ None
```

[GET /info](https://github.com/realTristan/Hermes/blob/master/server/handlers/info.go)
```
About
➤ Get the cache and full-text cache statistics.

URL Parameters
➤ None
```

[GET /exists](https://github.com/realTristan/Hermes/blob/master/server/handlers/exists.go)
```
About
➤ Get whether a key exists in the cache.

URL Parameters
➤ key: string
```

## Full-Text

[POST /ft/init](https://github.com/realTristan/Hermes/blob/master/server/handlers/init.go)
```
About
➤ Intialize the full text cache.

URL Parameters
➤ None
```

[POST /ft/clean](https://github.com/realTristan/Hermes/blob/master/server/handlers/clean.go)
```
About
➤ Clean all of the data in the full-text cache.

URL Parameters
➤ None
```

[GET /ft/search](https://github.com/realTristan/Hermes/blob/master/server/handlers/search.go)
```
About
➤ Search for a query in the full-text cache.

URL Parameters
➤ q: string
➤ strict: bool
➤ limit: int
➤ schema: map[string]bool
```

[GET /ft/searchoneword](https://github.com/realTristan/Hermes/blob/master/server/handlers/search.go)
```
About
➤ Search for a single word in the full-text cache.

URL Parameters
➤ q: string
➤ strict: bool
➤ limit: int
```

[GET /ft/searchvalues](https://github.com/realTristan/Hermes/blob/master/server/handlers/search.go)
```
About
➤ Search in the cache data values. (Slower)

URL Parameters
➤ q: string
➤ limit: int
➤ schema: map[string]bool
```

[GET /ft/searchvalueswithkey](https://github.com/realTristan/Hermes/blob/master/server/handlers/search.go)
```
About
➤ Search in the cache data values for a specific key.

URL Parameters
➤ q: string
➤ limit: int
➤ key: string
```

[POST /ft/maxbytes](https://github.com/realTristan/Hermes/blob/master/server/handlers/fulltext.go)
```
About
➤ Set the maximum full-text cache size in bytes.

URL Parameters
➤ maxbytes: int
```

[POST /ft/maxwords](https://github.com/realTristan/Hermes/blob/master/server/handlers/fulltext.go)
```
About
➤ Set the maximum full-text cache words allowed to be stored.

URL Parameters
➤ maxwords: int
```

[GET /ft/cache](https://github.com/realTristan/Hermes/blob/master/server/handlers/fulltext.go)
```
About
➤ Get the current full-text word cache. Can be used for testing.

URL Parameters
➤ None
```

[GET /ft/cachesize](https://github.com/realTristan/Hermes/blob/master/server/handlers/fulltext.go)
```
About
➤ Get the current full-text cache size.

URL Parameters
➤ None
```

[GET /ft/isinitialized](https://github.com/realTristan/Hermes/blob/master/server/handlers/fulltext.go)
```
About
➤ Get whether the full-text cache has been initialized.

URL Parameters
➤ None
```

[POST /ft/add](https://github.com/realTristan/Hermes/blob/master/server/handlers/add.go)
```
About
➤ Add value data to the full-text cache. Key must already exist in the cache.

URL Parameters
➤ key: string
```

# To-do
- Testing Files
- More Documentation
- More Examples
- Wrappers for other languages

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
