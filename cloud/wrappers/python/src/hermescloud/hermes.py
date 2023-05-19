import json, typing, websockets
from . import utils

# Wrapper for the Hermes cache
class Hermes:
    def __init__(self, addr: str):
        self.addr = addr
        self.ws = websockets.connect(f"ws://{self.addr}/ws/hermes")

    # With full text
    def with_ft(self, value: str) -> dict[str, typing.Any]:
        return {
            "$hermes.value": value, 
            "$hermes.full_text": True
        }

    # Set a value in the cache
    async def set(self, key: str, value: dict[str, typing.Any]) -> dict:
        # convert the value to a json string
        json_value: str = json.dumps(value)
        # base64 encode the value
        b64_value: str = utils.tob64(json_value)
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.set",
                "key": key, 
                "value": b64_value,
            }))
            return json.loads(await conn.recv())
    
    # Get a value from the cache
    async def get(self, key: str) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.get",
                "key": key
            }))
            return json.loads(await conn.recv())
    
    # Delete a value from the cache
    async def delete(self, key: str) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.delete",
                "key": key
            }))
            return json.loads(await conn.recv())
    
    # Get all keys in the cache
    async def keys(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.keys"
            }))
            return json.loads(await conn.recv())
    
    # Get all values in the cache
    async def values(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.values"
            }))
            return json.loads(await conn.recv())
    
    # Get the cache length
    async def length(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.length"
            }))
            return json.loads(await conn.recv())
    
    # Clear the cache
    async def clean(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.clean"
            }))
            return json.loads(await conn.recv())
    
    # Get the cache info
    async def info(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.info"
            }))
            return json.loads(await conn.recv())
    
    # Check if value exists in the cache
    async def exists(self, key: str) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "cache.exists",
                "key": key
            }))
            return json.loads(await conn.recv())
    
    # Intialize the full text cache
    async def ft_init(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.init"
            }))
            return json.loads(await conn.recv())
    
    # Clean the full text cache
    async def ft_clean(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.clean"
            }))
            return json.loads(await conn.recv())
    
    # Search the full text cache
    async def ft_search(self, query: str, strict: bool, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.search",
                "query": query,
                "strict": strict,
                "limit": limit,
                "schema": b64_schema
            }))
            return json.loads(await conn.recv())
    
    # Search one word in the full text cache
    async def ft_search_one(self, query: str, strict: bool, limit: int) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.search.one",
                "query": query,
                "strict": strict,
                "limit": limit
            }))
            return json.loads(await conn.recv())
    
    # Search value in the full text cache
    async def ft_search_values(self, query: str, limit: int, schema: dict[str, bool]) -> dict:
        # convert the schema to a json string
        json_schema: str = json.dumps(schema)
        # base64 encode the schema
        b64_schema: str = utils.tob64(json_schema)
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.search.values",
                "query": query,
                "limit": limit,
                "schema": b64_schema
            }))
            return json.loads(await conn.recv())
    
    # Search values with a key in the full text cache
    async def ft_search_with_key(self, query: str, key: str, limit: int) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.search.withkey",
                "query": query,
                "key": key,
                "limit": limit
            }))
            return json.loads(await conn.recv())
    
    # Set the max bytes of the full text cache
    async def ft_set_max_bytes(self, max_bytes: int) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.maxbytes.set",
                "maxbytes": max_bytes
            }))
            return json.loads(await conn.recv())
    
    # Set the max length of the full text storage
    async def ft_set_max_length(self, max_length: int) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.maxlength.set",
                "maxlength": max_length
            }))
            return json.loads(await conn.recv())
    
    # Get the full text storage
    async def ft_storage(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.storage"
            }))
            return json.loads(await conn.recv())
    
    # Get the full text storage size
    async def ft_size(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.storage.size"
            }))
            return json.loads(await conn.recv())
    
    # Get the full text storage length
    async def ft_length(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.storage.length"
            }))
            return json.loads(await conn.recv())
    
    # Get whether the full text cache is initialized
    async def ft_initialized(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.initialized"
            }))
            return json.loads(await conn.recv())

    # Sequence the ft indices
    async def ft_sequence(self) -> dict:
        # send the request
        async with self.ws as conn:
            await conn.send(json.dumps({
                "function": "ft.indices.sequence"
            }))
            return json.loads(await conn.recv())
    