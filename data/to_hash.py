import json, hashlib, time

# // Read the _data.json file
data = {}

# // Iterate over the data
for i, v in enumerate(json.load(open("../data/data.json", "r"))):
  while hash in data:
    hash: str = hashlib.sha1(str(time.time_ns()).encode("utf-8")).hexdigest()
    data[hash] = v

# // Write the data to the data.json file with spacing
json.dump(data, open("../data/data.json", "w"), indent=4)