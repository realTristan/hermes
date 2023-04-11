# Hermes ![Stars](https://img.shields.io/github/stars/realTristan/Hermes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/Hermes?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/230987049-665418b1-3576-49b7-861e-29036859ad8a.png)

# About
## Storing Data
Hermes works by iterating over the items in the data.json file, and then iterates over the keys and values of the items and splits the value into different words. It then stores the indices for all of the items that contain those words in a dictionary.

## Accessing Data
When searching for a word, Hermes will return a list of indices for all of the items that contain that word. It checks whether the key in the cache dictionary contains the provided word, instead of just accessing it so that short forms for words can be used.

## How to improve the speed
Instead of iterating over all of the keys in the cache and checking whether they contain the word you're looking for, just immediately access the indices by map index. ex: return cache[word] instead of for(keys in cache) if key contains word...

## Benchmarks
### Dataset
**Keys**: 4,115

**Total Words**: 208,092

**Map Size**: 33,048 bytes

<br>

### Speeds
**Python + Flask**: 626.4µs -> 1.03ms


**Golang + net/http**: 310.8µs -> 856.6µs


**Rust + actixweb**: 833.3µs

# Example
```py
from fastapi.middleware.cors import CORSMiddleware
from fastapi import FastAPI, Request
from cache import Cache
import time

# // The FastAPI app
app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["GET"],
    allow_headers=["*"],
)

# // The cache
cache: Cache = Cache()
cache.load()

# // Courses endpoint
@app.get("/courses")
async def root(request: Request):
    # // Get the course to search for from the query params
    course: str = request.query_params.get("q", "CS")

    # // Search for a word in the cache
    start_time: float = time.time()

    # // Search for the course
    courses: list[dict] = cache.search(course)

    # // Print the result
    print(f"Found {len(courses)} results in {time.time() - start_time} seconds")

    # // Return the items
    return courses

# // Run the app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)
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
