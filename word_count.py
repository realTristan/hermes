import json, sys

# // Store the amount of words in the data set
total_words: int  = 0

# // Load the json data
data = json.load(open("data.json"))

# // Iterate over the json itmes
for item in data:
    for k,v in item.items():
        v.strip().split()
        total_words += len(v)

# // Print the results
print(f"Total words: {total_words}")
print(f"Total keys: {len(data)}")
print(f"Data size: {sys.getsizeof(data)} bytes")