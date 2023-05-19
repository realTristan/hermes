import requests, json, typing, socket
from . import utils

# Wrapper for the Hermes cache
class Hermes:
    def __init__(self, addr: str):
        self.addr = addr

        # connect to the address
        self.connection = None

    # With full text
    def with_ft(self, value: str) -> dict[str, typing.Any]:
        return {"$hermes.value": value, "$hermes.full_text": True}

    # Set a value in the cache
    def set(self, key: str, value: dict[str, typing.Any]) -> dict:
        # convert the value to a json string
        json_value: str = json.dumps(value)
        # base64 encode the value
        b64_value: str = utils.tob64(json_value)
        # send the request
        r: requests.Response = requests.post(
            self.addr + "/set", params={"key": key, "value": b64_value})
        return r.json()
    
    # Get a value from the cache
    def get(self, key: str) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/get", params={"key": key})
        return r.json()
    
    # Delete a value from the cache
    def delete(self, key: str) -> dict:
        # send the request
        r: requests.Response = requests.delete(self.addr + "/delete", params={"key": key})
        return r.json()
    
    # Get all keys in the cache
    def keys(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/keys")
        return r.json()
    
    # Get all values in the cache
    def values(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/values")
        return r.json()
    
    # Get the cache length
    def length(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/length")
        return r.json()
    
    # Clear the cache
    def clean(self) -> dict:
        # send the request
        r: requests.Response = requests.delete(self.addr + "/clean")
        return r.json()
    
    # Get the cache info
    def info(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/info")
        return r.json()
    
    # Check if value exists in the cache
    def exists(self, key: str) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/exists", params={"key": key})
        return r.json()
    
    # Intialize the full text cache
    def ft_init(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/init")
        return r.json()
    
    # Clean the full text cache
    def ft_clean(self) -> dict:
        # send the request
        r: requests.Response = requests.delete(self.addr + "/ft/clean")
        return r.json()
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        r: requests.Response = requests.get(
            self.addr + "/ft/search", params={"q": query, "strict": strict, "limit": limit, "schema": b64_schema})
        return r.json()
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> dict:
        # send the request
        r: requests.Response = requests.get(
            self.addr + "/ft/search/oneword", params={"q": query, "strict": strict, "limit": limit})
        return r.json()
    
    # Search value in the full text cache
    def ft_search_values(self, query: str, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        r: requests.Response = requests.get(
            self.addr + "/ft/search/values", params={"q": query, "limit": limit, "schema": b64_schema})
        return r.json()
    
    # Search values with a key in the full text cache
    def ft_search_with_key(self, query: str, key: str, limit: int) -> dict:
        # send the request
        r: requests.Response = requests.get(
            self.addr + "/ft/search/withkey", params={"q": query, "key": key, "limit": limit})
        return r.json()
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> dict:
        # send the request
        r: requests.Response = requests.post(
            self.addr + "/ft/maxbytes", params={"maxbytes": max_bytes})
        return r.json()
    
    # Set the max length of the full text storage
    def ft_set_max_length(self, max_length: int) -> dict:
        # send the request
        r: requests.Response = requests.post(
            self.addr + "/ft/maxlength", params={"maxlength": max_length})
        return r.json()
    
    # Get the full text storage
    def ft_storage(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/storage")
        return r.json()
    
    # Get the full text storage size
    def ft_size(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/storage/size")
        return r.json()
    
    # Get the full text storage length
    def ft_length(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/storage/length")
        return r.json()
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/isinitialized")
        return r.json()

    # Sequence the ft indices
    def ft_sequence(self) -> dict:
        # send the request
        r: requests.Response = requests.get(self.addr + "/ft/indices/sequence")
        return r.json()
    