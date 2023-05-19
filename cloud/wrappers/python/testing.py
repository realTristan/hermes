from hermescloud import Cache
import time

# Create a new cache instance
cache = Cache("localhost:3000")

def main():
    # Initialize the full-text search engine
    print(cache.ft_init(-1, -1))

    # Set a value
    cache.set("user_id", {
        "name": {
            "$hermes.full_text": True,
            "$hermes.value": "tristan"
        }
    })

    # Get a value
    print(cache.get("user_id"))

    # Track the start time
    start_time = time.time()

    # Search for a value
    print(cache.ft_search("tristan", False, 100, {
        "name": True
    }))

    # Print the duration (average: 0.0006s)
    print(f"Duration: {time.time() - start_time}s")

    # Delete a value
    print(cache.delete("user_id"))

    # Get a value
    print(cache.get("user_id"))

# Run the main function
if __name__ == "__main__":
    main()