# Hermes
Extremely Fast Full-Text-Search Algorithm

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
    allow_methods=["*"],
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
    indices: list[int] = cache.search(course)
    print(f"Search time: {time.time() - start_time} seconds")

    # // Convert the indices to the actual items
    items: list[dict] = cache.indices_to_data(indices)
    print(f"Conversion time: {time.time() - start_time} seconds")

    # // Return the items
    return items

# // Run the app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)
```

# Cache Code
```py
import json, sys

# // The cache class
class Cache:
    def __init__(self) -> None:
        self.cache: dict = {}
        self.data: dict = json.load(open("data.json", "r"))

    # // Load the cache from the data.json file
    def load(self) -> None:
        for i, course in enumerate(self.data):
            for _, v in course.items():
                # // Split the string by spaces, then iterate over the words
                words: list[str] = v.lower().strip().split()
                for word in words:
                    # // If the word is not in the cache, add it
                    if word not in self.cache:
                        self.cache[word] = []
                    
                    # // Append the index of the item for this word to the cache
                    self.cache[word].append(i)

    # // Search for a word in the cache
    def search(self, word: str) -> list[int]:
        res: list[int] = []
        for k, v in self.cache.items():
            if word in k:
                res.extend(v)
        return res
    
    # // Convert the indices to the actual items
    def indices_to_data(self, indices: list[int]) -> list[dict]:
        items: list[dict] = []
        for i in indices:
            items.append(self.data[i])
        return items
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
