import hermescloud

# Create a Hermes instance
hermes = hermescloud.Hermes("localhost:3000")

def main():
    # Initialize the full-text search engine
    print(hermes.ft_init(-1, -1))

    # Set a value
    hermes.set("user_id", {
        "name": {
            "$hermes.full_text": True,
            "$hermes.value": "tristan"
        }
    })

    # Get a value
    print(hermes.get("user_id"))

    # Search for a value
    print(hermes.ft_search("tristan", False, 100, {
        "name": True
    }))

    # Delete a value
    print(hermes.delete("user_id"))

    # Get a value
    print(hermes.get("user_id"))

# Run the main function
if __name__ == "__main__":
    main()