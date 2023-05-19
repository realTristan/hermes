# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

# Go Install
```
go get github.com/realTristan/Hermes
```

# Python Install
```
pip install hermescloud
```

# What is Hermes?
Hermes is an extremely fast full-text search algorithm and caching system written in Go. It's designed for API implementations which can be used by wrappers in other languages.
Hermes has two notable algorithms. The first being the with-cache algorithm. When using the with-cache algorithm, you can set, get, store, etc. keys and values into the cache. The second
being a no-cache algorithm that reads data from a map, or json file, and uses and array to store the data. Both of these algorithms provide full-text search query times from 10µs to 300µs.

# Example of NoCache Full-Text-Search
If you want to use only the full-text-search features, then just import hermes and load it using a .json file. (as shown in /example). Note: For small to medium-sized datasets (like the ones I used in /data), Hermes works great. Although, as the words in the dataset increases, the full-text-search cache will take up significantly more memory.

## Benchmarks
```
Dataset Map Entries: 4,115
Dataset Total Words: 208,092
Dataset Map Size: ≈ 2.3MB

Query: Computer, Limit: 100, Strict: False => 36.5µs
Query: Computer, Limit: 100, Strict: True => 12.102µs
```

## Code
```go
package main

import (
	"fmt"
	"time"

	Hermes "github.com/realTristan/Hermes/nocache"
)

func main() {
	// Initialize the full-text cache
	var ft, _ = Hermes.InitWithJson("data.json")
	
	// Create a schema. These are the fields that will be searched.
	var schema = map[string]bool{
		"name": true,
		"id": false,
	}

	// Search for a word in the cache
	// @params: query, limit, strict
	var res, _ = ft.Search("tristan", 100, false, schema)
	fmt.Println(res)
}
```

## data.json
```json
[
    {
	"id": 1,
        "name": {
            "$hermes.full_text": true,
            "$hermes.value": "Tristan Simpson"
        },
    },
]
```

# Example of Cache Full-Text-Search
The full-text-search from /cache is significantly slower than the nocache FTS. Why? Because the FTS in /cache requires more memory, keys, and utilizes a map instead of a slice to store data. If you want to use a cache along with the full text-search algorithm, import the files from /cache. To setup a cache, check out /cache/example or /cache/testing. 

## Benchmarks
```
Dataset Map Entries: 4,115
Dataset Total Words: 208,092
Dataset Map Size: ≈ 2.3MB

Query: Computer, Limit: 100, Strict: False => 263.7µs
Query: Computer, Limit: 100, Strict: True => 40.84µs
```

## Code
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
	
	// Initialize the full-text cache
	// MaxLength: 10, MaxBytes: -1 (no limit)
	cache.FTInit(10, -1)
	
	// Set the value in the cache
	cache.Set("user_id", map[string]interface{}{
		"name":       Hermes.WithFT("tristan"),
		"age":        17,
		"expiration": time.Now(),
	})
	
	// The keys you want to search through in the full-text search
	var schema map[string]bool = map[string]bool{
		"name":       true,
		"age":        false,
		"expiration": false,
	}

	// Search for a word in the cache and print the result
	var result, _ = cache.Search("tristan", 100, false, schema)
	fmt.Println(result)
}
```

## data.json
### Used with cache.FTInitWithJson()
```json
{
  "user_id": {
    "age": 17,
    "name": {
      "$hermes.full_text": true,
      "$hermes.value": "Tristan Simpson"
    },
  },
}
```

# Hermes Cloud App
## Install
```
Coming Soon
```

## Custom Implementation
```go
import (
	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Socket "github.com/realTristan/Hermes/socket"
)

func main() {
	// Cache and fiber app
	cache := Hermes.InitCache()
	app := fiber.New()

	// Set the router
	Socket.SetRouter(app, cache)

	// Listen on port 3000
	app.Listen(":3000")
}
```

# Websocket API
## Cache

### [cache.set](https://github.com/realTristan/Hermes/blob/master/socket/handlers/set.go)

#### About
```
Set a value in the cache with the corresponding key.
```

#### Example Request
```go
{
  "function": "cache.set",
  "key": "user_id",
  "value": base64{
    "name": "tristan"
  }
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [cache.delete](https://github.com/realTristan/Hermes/blob/master/socket/handlers/delete.go)

#### About
```
Delete the provided key, and the data correlated to it from the cache.
```

#### Example Request
```go
{
  "function": "cache.delete",
  "key": "user_id"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [cache.get](https://github.com/realTristan/Hermes/blob/master/socket/handlers/get.go)

#### About
```
Get data from the cache using a key.
```

#### Example Request
```go
{
  "function": "cache.get",
  "key": "user_id"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": map[string]any
}
```

### [cache.keys](https://github.com/realTristan/Hermes/blob/master/socket/handlers/keys.go)

#### About
```
Get all of the keys in the cache.
```

#### Example Request
```go
{
  "function": "cache.keys"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []string
}
```

### [cache.values](https://github.com/realTristan/Hermes/blob/master/socket/handlers/values.go)

#### About
```
Get all of the values in the cache.
```

#### Example Request
```go
{
  "function": "cache.values"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []map[string]any
}
```

### [cache.length](https://github.com/realTristan/Hermes/blob/master/socket/handlers/length.go)

#### About
```
Get the amount of keys stored in the cache.
```

#### Example Request
```go
{
  "function": "cache.length"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": int
}
```

### [cache.clean](https://github.com/realTristan/Hermes/blob/master/socket/handlers/clean.go)

#### About
```
Clean all the data in the cache, and full-text storage.
```

#### Example Request
```go
{
  "function": "cache.clean"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [cache.info](https://github.com/realTristan/Hermes/blob/master/socket/handlers/info.go)

#### About
```
Get the cache and full-text storage statistics.
```

#### Example Request
```go
{
  "function": "cache.info"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": string
}
```

### [cache.exists](https://github.com/realTristan/Hermes/blob/master/socket/handlers/exists.go)

#### About
```
Get whether a key exists in the cache.
```

#### Example Request
```go
{
  "function": "cache.exists",
  "key": "user_id"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": bool
}
```

## Full-Text

### [ft.init](https://github.com/realTristan/Hermes/blob/master/socket/handlers/init.go)

#### About
```
Intialize the full text cache.
```

#### Example Request
```go
{
  "function": "ft.init",
  "maxbytes": -1,
  "maxlength": -1
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.clean](https://github.com/realTristan/Hermes/blob/master/socket/handlers/clean.go)

#### About
```
Clean all of the data in the full-text storage.
```

#### Example Request
```go
{
  "function": "ft.clean"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.search](https://github.com/realTristan/Hermes/blob/master/socket/handlers/search.go)

#### About
```
Search for a query in the full-text storage.
```

#### Example Request
```go
{
  "function": "ft.search",
  "query": "tristan",
  "strict": false,
  "limit": 10,
  "schema": {
    "key": true
  }
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []map[string]any
}
```

### [ft.search.oneword](https://github.com/realTristan/Hermes/blob/master/socket/handlers/search.go)

#### About
```
Search for a single word in the full-text storage.
```

#### Example Request
```go
{
  "function": "ft.search.oneword",
  "query": "tristan",
  "strict": false,
  "limit": 10
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []map[string]any
}
```

### [ft.search.values](https://github.com/realTristan/Hermes/blob/master/socket/handlers/search.go)

#### About
```
Search in the cache data values. (Slower)
```

#### Example Request
```go
{
  "function": "ft.search.values",
  "query": "tristan",
  "limit": 10,
  "schema": {
    "key": true
  }
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []map[string]any
}
```

### [ft.search.withkey](https://github.com/realTristan/Hermes/blob/master/socket/handlers/search.go)

#### About
```
Search in the cache data values for a specific key.
```

#### Example Request
```go
{
  "function": "ft.search.withkey",
  "query": "tristan",
  "key": "user_id",
  "limit": 10
}
```

#### Response
```go
{
  "success": true/false, 
  "data": []map[string]any
}
```

### [ft.maxbytes.set](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Set the maximum full-text storage size in bytes.
```

#### Example Request
```go
{
  "function": "ft.maxbytes.set",
  "maxbytes": -1
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.maxlength.set](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Set the maximum full-text storage words allowed to be stored.
```

#### Example Request
```go
{
  "function": "ft.maxlength.set",
  "maxlength": -1
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.storage](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Get the current full-text storage.
```

#### Example Request
```go
{
  "function": "ft.storage"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": map[string][]int
}
```

### [ft.storage.size](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Get the current full-text storage size in bytes.
```

#### Example Request
```go
{
  "function": "ft.storage.size"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": int
}
```

### [ft.storage.length](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Get the current full-text storage length.
```

#### Example Request
```go
{
  "function": "ft.storage.length"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": int
}
```

### [ft.isinitialized](https://github.com/realTristan/Hermes/blob/master/socket/handlers/fulltext.go)

#### About
```
Get whether the full-text storage has been initialized.
```

#### Example Request
```go
{
  "function": "ft.isinitialized"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": bool
}
```

### [ft.indices.sequence](https://github.com/realTristan/Hermes/blob/master/socket/handlers/indices.go)

#### About
```
Sequence the full-text storage indices.
```

#### Example Request
```go
{
  "function": "ft.indices.sequence"
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
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
