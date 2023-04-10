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
    indices: list[int] = cache.search(course)

    # // Convert the indices to the actual items
    items: list[dict] = cache.indices_to_data(indices)
    print(f"Time: {time.time() - start_time}s")

    # // Return the items
    return items

# // Run the app
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)