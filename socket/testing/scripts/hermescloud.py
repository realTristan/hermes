from hermescloud import Hermes

# Create a Hermes instance
hermes = Hermes("localhost:8080")

# Initialize the full-text search engine
hermes.ft_init()

# Set a value
hermes.set("user_id", {
    "name": {
        "$hermes.value": "tristan",
        "$hermes.full_text": True
    }
})

# Get a value
print(hermes.get("user_id"))

# Search for a value
print(hermes.search("tristan"))