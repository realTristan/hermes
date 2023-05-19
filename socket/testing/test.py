import websockets, asyncio, base64, json, time

def encode(value):
    return base64.b64encode(value.encode("utf-8")).decode("utf-8")

def decode(value):
    return base64.b64decode(value.encode("utf-8")).decode("utf-8")

def test_set():
    value = {
        "name": {
            "$hermes.value": "tristan",
            "$hermes.full_text": True
        }
    }
    return json.dumps({
        "function": "cache.set",
        "key": "user_id",
        "value": encode(json.dumps(value))
    })

def test_get():
    return json.dumps({
        "function": "cache.get",
        "key": "test"
    })

def test_search():
    return json.dumps({
        "function": "ft.search",
        "query": "test",
        "limit": 100,
        "strict": False,
        "schema": encode(json.dumps({
            "test": True
        }))
    })

def test_ft_init():
    return json.dumps({
        "function": "ft.init",
        "maxlength": -1,
        "maxbytes": -1,
    })

# connect to wss://127.0.0.1:3000/ws/hermes
async def test():
    async with websockets.connect("ws://127.0.0.1:3000/ws/hermes") as websocket:
        # test ft init
        await websocket.send(test_ft_init())
        print(await websocket.recv())

        # test set
        await websocket.send(test_set())
        print(await websocket.recv())

        # test get
        await websocket.send(test_get())
        print(await websocket.recv())

        # track start time
        start = time.time()

        # test search
        await websocket.send(test_search())
        print(f"Search results: {await websocket.recv()}")

        # print time taken
        print("Time taken: " + str(time.time() - start))

        # close socket
        await websocket.close()

asyncio.run(test())