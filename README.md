# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/hermes?label=Watchers)
<img width="1173" alt="Screenshot 2023-06-03 at 1 55 31 PM" src="https://github.com/realTristan/hermes/assets/75189508/f884f643-00af-405d-900d-a636d119cf74">

# How to contribute
Hermes is still in a developmental phase. Helping with test cases, benchmarks, speed, and memory usage are a great way to contribute to the project.

## Hermes Go
Direct access to the hermes cache/nocache functions.
```
go get github.com/realTristan/hermes
```

## Hermes Cloud
Access the hermes cache functions via websocket.

### Python
```
pip install hermescloud
```

# What is Hermes?
Hermes is an extremely fast full-text search and caching framework written in Go.
Hermes has a cache algorithm where you can set, get, store, etc. keys and values into the cache, and a no-cache algorithm that reads data from an array or json file and uses an array to store data. Both of these algorithms provide full-text search query speeds from 10µs to 300µs.

# Example of NoCache
If you want to use only the full-text-search features, then import hermes/nocache. Load the data using a .json file or array of maps. No Cache is much faster than the With-Cache algorithm, but you can't update the data.

## Code
```go
package main

import (
  "fmt"
  "time"

  hermes "github.com/realTristan/hermes/nocache"
)

func main() {
  // Initialize the full-text cache
  var ft, _ = hermes.InitWithJson("data.json")

  // Search for a word in the cache
  var res, _ = ft.Search(hermes.SearchParams{
    Query:  "tristan",
    Limit:  100,
    Strict: false,
  })

  fmt.Println(res)
}
```

### Benchmarks
```
Dataset Map Entries: 4,115
Dataset Total Words: 208,092
Dataset Map Size: ≈ 2.3MB

Query: Computer, Limit: 100, Strict: False => 36.5µs
Query: Computer, Limit: 100, Strict: True => 12.102µs
```

### data.json
Used with hermes.InitWithJson()
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

# With-Cache Example
The with-cache algorithm is notably slower than the no-cache algorithm. Why? Because the with-cache algorithm uses maps to store data and not an array. Import hermes directly to use the with-cache features.

## Code
```go
package main

import (
  "fmt"
  "time"
  hermes "github.com/realTristan/hermes"
)

func main() {
  // Initialize the cache
  cache := hermes.InitCache()

  // MaxSize: 10, MaxBytes: -1 (no limit), MinWordLength: 3
  cache.FTInit(10, -1, 3)

  // Set the value in the cache
  cache.Set("user_id", map[string]any{
    "name":       cache.WithFT("tristan"),
    "age":        17,
  })

  // Search for a word in the cache and print the result
  result, err := cache.Search(hermes.SearchParams{
    Query:  "tristan",
    Limit:  100,
    Strict: false,
  })

  fmt.Println(result)
}
```

### Benchmarks
```
Dataset Map Entries: 4,115
Dataset Total Words: 208,092
Dataset Map Size: ≈ 2.3MB

Query: Computer, Limit: 100, Strict: False => 263.7µs
Query: Computer, Limit: 100, Strict: True => 40.84µs
```

### data.json
Used with cache.FTInitWithJson()
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
  hermes "github.com/realTristan/hermes"
  socket "github.com/realTristan/hermes/cloud/socket"
)

func main() {
  // Cache and fiber app
  cache := hermes.InitCache()
  app := fiber.New()

  // Set the router
  socket.SetRouter(app, cache)

  // Listen on port 3000
  app.Listen(":3000")
}
```

# Websocket API
## Cache

### [cache.set](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/set.go)

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

### [cache.delete](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/delete.go)

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

### [cache.get](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/get.go)

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

### [cache.keys](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/keys.go)

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

### [cache.values](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/values.go)

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

### [cache.length](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/length.go)

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

### [cache.clean](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/clean.go)

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

### [cache.info](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/info.go)

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

### [cache.exists](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/exists.go)

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

### [ft.init](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/init.go)

#### About
```
Intialize the full text cache.
```

#### Example Request
```go
{
  "function": "ft.init",
  "maxbytes": -1,
  "maxsize": -1,
  "minwordlength": 3
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.clean](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/clean.go)

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

### [ft.search](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/search.go)

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

### [ft.search.oneword](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/search.go)

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

### [ft.search.values](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/search.go)

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

### [ft.search.withkey](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/search.go)

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

### [ft.maxbytes.set](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

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

### [ft.maxsize.set](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

#### About
```
Set the maximum full-text storage words allowed to be stored.
```

#### Example Request
```go
{
  "function": "ft.maxsize.set",
  "maxsize": -1
}
```

#### Response
```go
{
  "success": true/false, 
  "data": nil
}
```

### [ft.storage](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

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

### [ft.storage.size](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

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

### [ft.storage.length](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

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

### [ft.isinitialized](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/fulltext.go)

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

### [ft.indices.sequence](https://github.com/realTristan/hermes/blob/master/cloud/socket/handlers/indices.go)

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
