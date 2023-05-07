import requests, json, utils

# Wrapper for the Hermes cache
class Hermes:
    def __init__(self, addr: str):
        self.addr = addr

    # With full text
    def with_ft(self, value: str) -> dict[str, any]:
        return {"value": value, "$hermes.full_text": True}

    # Set a value in the cache
    def set(self, key: str, value: dict[str, any]) -> any:
        # convert the value to a json string
        json_value: str = json.dumps(value)
        # base64 encode the value
        b64_value: str = utils.tob64(json_value)
        # send the request
        r: requests.request = requests.post(
            self.addr + "/set", params={"key": key, "value": b64_value})
        return r.json()
    
    # Get a value from the cache
    def get(self, key: str) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/get", params={"key": key})
        return r.json()
    
    # Delete a value from the cache
    def delete(self, key: str) -> any:
        # send the request
        r: requests.request = requests.delete(self.addr + "/delete", params={"key": key})
        return r.json()
    
    # Get all keys in the cache
    def keys(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/keys")
        return r.json()
    
    # Get all values in the cache
    def values(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/values")
        return r.json()
    
    # Get the cache length
    def length(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/length")
        return r.json()
    
    # Clear the cache
    def clean(self) -> any:
        # send the request
        r: requests.request = requests.delete(self.addr + "/clean")
        return r.json()
    
    # Get the cache info
    def info(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/info")
        return r.json()
    
    # Check if value exists in the cache
    def exists(self, key: str) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/exists", params={"key": key})
        return r.json()
    
    # Intialize the full text cache
    def ft_init(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/ft/init")
        return r.json()
    
    # Clean the full text cache
    def ft_clean(self) -> any:
        # send the request
        r: requests.request = requests.delete(self.addr + "/ft/clean")
        return r.json()
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> any:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        r: requests.request = requests.get(
            self.addr + "/ft/search", params={"q": query, "strict": strict, "limit": limit, "schema": b64_schema})
        return r.json()
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> any:
        # send the request
        r: requests.request = requests.get(
            self.addr + "/ft/searchoneword", params={"q": query, "strict": strict, "limit": limit})
        return r.json()
    
    # Search value in the full text cache
    def ft_search_value(self, query: str, limit: int, schema: dict[str, bool]) -> any:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        r: requests.request = requests.get(
            self.addr + "/ft/searchvalue", params={"q": query, "limit": limit, "schema": b64_schema})
        return r.json()
    
    # Search values with a key in the full text cache
    def ft_search_key(self, query: str, key: str, limit: int) -> any:
        # send the request
        r: requests.request = requests.get(
            self.addr + "/ft/searchkey", params={"q": query, "key": key, "limit": limit})
        return r.json()
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> any:
        # send the request
        r: requests.request = requests.post(
            self.addr + "/ft/maxbytes", params={"maxbytes": max_bytes})
        return r.json()
    
    # Set the max words of the full text cache
    def ft_set_max_words(self, max_words: int) -> any:
        # send the request
        r: requests.request = requests.post(
            self.addr + "/ft/maxwords", params={"maxwords": max_words})
        return r.json()
    
    # Get the full text cache
    def ft_cache(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/ft/cache")
        return r.json()
    
    # Get the full text cache size
    def ft_size(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/ft/cachesize")
        return r.json()
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> any:
        # send the request
        r: requests.request = requests.get(self.addr + "/ft/isinitialized")
        return r.json()
    