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
        self.ws.send(json.dumps({
            "function": "cache.set",
            "key": key, 
            "value": b64_value,
        }))
        return json.loads(self.ws.recv())
    
    # Get a value from the cache
    def get(self, key: str) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.get",
            "key": key
        }))
        return json.loads(self.ws.recv())
    
    # Delete a value from the cache
    def delete(self, key: str) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.delete",
            "key": key
        }))
        return json.loads(self.ws.recv())
    
    # Get all keys in the cache
    def keys(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.keys"
        }))
        return json.loads(self.ws.recv())
    
    # Get all values in the cache
    def values(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.values"
        }))
        return json.loads(self.ws.recv())
    
    # Get the cache length
    def length(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.length"
        }))
        return json.loads(self.ws.recv())
    
    # Clear the cache
    def clean(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.clean"
        }))
        return json.loads(self.ws.recv())
    
    # Get the cache info
    def info(self) -> str:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.info"
        }))
        return self.ws.recv()
    
    # Get cache info for testing
    def info_testing(self) -> str:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.info.testing"
        }))
        return self.ws.recv()
    
    # Check if value exists in the cache
    def exists(self, key: str) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.exists",
            "key": key
        }))
        return json.loads(self.ws.recv())
    
    # Intialize the full text cache
    def ft_init(self, maxbytes: int, maxlength: int) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.init",
            "maxbytes": maxbytes,
            "maxlength": maxlength
        }))
        return json.loads(self.ws.recv())
    
    # Clean the full text cache
    def ft_clean(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.clean"
        }))
        return json.loads(self.ws.recv())
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = Utils.tob64(json_schema)
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search",
            "query": query,
            "strict": strict,
            "limit": limit,
            "schema": b64_schema
        }))
        return json.loads(self.ws.recv())
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search.one",
            "query": query,
            "strict": strict,
            "limit": limit
        }))
        return json.loads(self.ws.recv())
    
    # Search value in the full text cache
    def ft_search_values(self, query: str, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = Utils.tob64(json_schema)
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search.values",
            "query": query,
            "limit": limit,
            "schema": b64_schema
        }))
        return json.loads(self.ws.recv())
    
    # Search values with a key in the full text cache
    def ft_search_with_key(self, query: str, key: str, limit: int) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search.withkey",
            "query": query,
            "key": key,
            "limit": limit
        }))
        return json.loads(self.ws.recv())
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.maxbytes.set",
            "maxbytes": max_bytes
        }))
        return json.loads(self.ws.recv())
    
    # Set the max length of the full text storage
    def ft_set_max_length(self, max_length: int) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.maxlength.set",
            "maxlength": max_length
        }))
        return json.loads(self.ws.recv())
    
    # Get the full text storage
    def ft_storage(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage"
        }))
        return json.loads(self.ws.recv())
    
    # Get the full text storage size
    def ft_size(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage.size"
        }))
        return json.loads(self.ws.recv())
    
    # Get the full text storage length
    def ft_length(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage.length"
        }))
        return json.loads(self.ws.recv())
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.initialized"
        }))
        return json.loads(self.ws.recv())

    # Sequence the ft indices
    def ft_sequence(self) -> dict:
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.indices.sequence"
        }))
        return json.loads(self.ws.recv())
    