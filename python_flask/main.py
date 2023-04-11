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
    query: str = request.query_params.get("q", "CS")
    limit: int = int(request.query_params.get("limit", 10))


    # // Search for a word in the cache
    start_time: float = time.time()

    # // Search for the course
    courses: list[dict] = cache.search(query, limit)

    # // Print the result
    print(f"Found {len(courses)} results in {time.time() - start_time} seconds")

    # // Return the items
    return courses

# // Run the app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)