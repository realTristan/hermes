import json, typing
from .utils import Utils
from websocket import create_connection

# Wrapper for the Hermes cache
class Cache:
    def __init__(self, addr: str):
        self.ws = create_connection(f"ws://{addr}/ws/hermes")

    # With full text
    def with_ft(self, value: str) -> dict[str, typing.Any]:
        return {
            "$hermes.value": value, 
            "$hermes.full_text": True
        }

    # Set a value in the cache
    def set(self, key: str, value: dict[str, typing.Any]) -> dict:
        # convert the value to a json string
        json_value: str = json.dumps(value)
        # base64 encode the value
        b64_value: str = Utils.tob64(json_value)
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.set",
            "key": key, 
            "value": b64_value,
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get a value from the cache
    def get(self, key: str) -> dict:
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.get",
            "key": key
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Delete a value from the cache
    def delete(self, key: str) -> dict:
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.delete",
            "key": key
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get all keys in the cache
    def keys(self) -> dict:
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.keys"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get all values in the cache
    def values(self) -> dict:
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.values"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get the cache length
    def length(self) -> dict:
        # send the request
        r: int = self.ws.send(json.dumps({
            "function": "cache.length"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Clear the cache
    def clean(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "cache.clean"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get the cache info
    def info(self) -> str:
        # send the request
        r =  self.ws.send(json.dumps({
            "function": "cache.info"
        }))
        if r:
            return self.ws.recv()
        return ""
    
    # Get cache info for testing
    def info_testing(self) -> str:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "cache.info.testing"
        }))
        if r:
            return self.ws.recv()
        return ""
    
    # Check if value exists in the cache
    def exists(self, key: str) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "cache.exists",
            "key": key
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Intialize the full text cache
    def ft_init(self, maxbytes: int, maxlength: int) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.init",
            "maxbytes": maxbytes,
            "maxlength": maxlength
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Clean the full text cache
    def ft_clean(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.clean"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = Utils.tob64(json_schema)
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.search",
            "query": query,
            "strict": strict,
            "limit": limit,
            "schema": b64_schema
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.search.one",
            "query": query,
            "strict": strict,
            "limit": limit
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Search value in the full text cache
    def ft_search_values(self, query: str, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = Utils.tob64(json_schema)
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.search.values",
            "query": query,
            "limit": limit,
            "schema": b64_schema
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Search values with a key in the full text cache
    def ft_search_with_key(self, query: str, key: str, limit: int) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.search.withkey",
            "query": query,
            "key": key,
            "limit": limit
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.maxbytes.set",
            "maxbytes": max_bytes
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Set the max length of the full text storage
    def ft_set_max_length(self, max_length: int) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.maxlength.set",
            "maxlength": max_length
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get the full text storage
    def ft_storage(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.storage"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get the full text storage size
    def ft_size(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.storage.size"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get the full text storage length
    def ft_length(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.storage.length"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.initialized"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}

    # Sequence the ft indices
    def ft_sequence(self) -> dict:
        # send the request
        r = self.ws.send(json.dumps({
            "function": "ft.indices.sequence"
        }))
        if r:
            return json.loads(self.ws.recv())
        return {}
    