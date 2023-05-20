import json, typing, threading, asyncio
from .utils import Utils
from websocket import create_connection

# Wrapper for the Hermes cache
class Cache:
    def __init__(self, addr: str):
        # socket
        self.ws = create_connection(f"ws://{addr}/ws/hermes")
        # threading lock
        self.lock = threading.Lock()

    # With full text
    def with_ft(self, value: str) -> dict[str, typing.Any]:
        return {
            "$hermes.value": value, 
            "$hermes.full_text": True
        }

    # Set a value in the cache
    def set(self, key: str, value: dict[str, typing.Any]) -> dict:
        # lock
        self.lock.acquire()
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
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get a value from the cache
    def get(self, key: str) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.get",
            "key": key
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Delete a value from the cache
    def delete(self, key: str) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.delete",
            "key": key
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get all keys in the cache
    def keys(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.keys"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get all values in the cache
    def values(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.values"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get the cache length
    def length(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.length"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Clear the cache
    def clean(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.clean"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get the cache info
    def info(self) -> str:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.info"
        }))
        # get the response
        response: str = self.ws.recv()
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get cache info for testing
    def info_testing(self) -> str:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.info.testing"
        }))
        # get the response
        response: str = self.ws.recv()
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Check if value exists in the cache
    def exists(self, key: str) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "cache.exists",
            "key": key
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Intialize the full text cache
    def ft_init(self, maxbytes: int, maxlength: int) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.init",
            "maxbytes": maxbytes,
            "maxlength": maxlength
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Clean the full text cache
    def ft_clean(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.clean"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Search the full text cache
    def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> dict:
        # lock
        self.lock.acquire()
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
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Search one word in the full text cache
    def ft_search_one(self, query: str, strict: bool, limit: int) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search.one",
            "query": query,
            "strict": strict,
            "limit": limit
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Search value in the full text cache
    def ft_search_values(self, query: str, limit: int, schema: dict[str, bool]) -> dict:
        # lock
        self.lock.acquire()
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
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Search values with a key in the full text cache
    def ft_search_with_key(self, query: str, key: str, limit: int) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.search.withkey",
            "query": query,
            "key": key,
            "limit": limit
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Set the max bytes of the full text cache
    def ft_set_max_bytes(self, max_bytes: int) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.maxbytes.set",
            "maxbytes": max_bytes
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Set the max length of the full text storage
    def ft_set_max_length(self, max_length: int) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.maxlength.set",
            "maxlength": max_length
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get the full text storage
    def ft_storage(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get the full text storage size
    def ft_size(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage.size"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get the full text storage length
    def ft_length(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.storage.length"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    
    # Get whether the full text cache is initialized
    def ft_initialized(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.initialized"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response

    # Sequence the ft indices
    def ft_sequence(self) -> dict:
        # lock
        self.lock.acquire()
        # send the request
        self.ws.send(json.dumps({
            "function": "ft.indices.sequence"
        }))
        # get the response
        response: dict = json.loads(self.ws.recv())
        # unlock
        self.lock.release()
        # return the response
        return response
    