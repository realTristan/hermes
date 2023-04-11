import json, sys

# // Store the amount of words in the data set
total_words: int  = 0

# // Load the json data
data = json.load(open("data.json"))

# // Iterate over the json itmes
for item in data:
    for k, v in item.items():
        for word in v.strip().split():
            if word.isalnum():
                total_words += 1

# // Print the results
print(f"Total words: {total_words}")
print(f"Total keys: {len(data)}")
print(f"Data size: {sys.getsizeof(data)} bytes")
