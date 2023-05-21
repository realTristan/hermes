import hermescloud, json, base64, time

# base64 encode a value
def base64_encode(value):
    return base64.b64encode(value.encode("utf-8")).decode("utf-8")

# Create a new cache instance
cache = hermescloud.Cache("localhost:3000")

# Initialize the full-text search engine
#print(cache.ft_init(-1, -1))

# open the data/data_hash.json file
def set_data():
    with open("data/data_hash.json", "r") as file:
        # load the data_hash.json file
        data = json.loads(file.read())

        # loop through the data
        for key in data:
            # set the key
            cache.set(key, data[key])

# set the data
#set_data()

# Track the start time
start_time = time.time()

# Search for a value
r = cache.ft_search("computer", False, 100, {
    "id":             False,
    "components":     False,
    "units":          False,
    "description":    True,
    "name":           True,
    "pre_requisites": True,
    "title":          True
})

# Print the duration
print(time.time() - start_time)

# print the results
#print(r[0])
