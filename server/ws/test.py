import websockets, asyncio, base64, json

def encode(value):
    return base64.b64encode(value.encode("utf-8")).decode("utf-8")

def decode(value):
    return base64.b64decode(value.encode("utf-8")).decode("utf-8")

def test_value():
    data = {
        "test": "test"
    }
    # convert data to json
    data = json.dumps(data)
    # encode data
    return encode(data)

# connect to wss://127.0.0.1:3000/ws/hermes/cache
async def test():
    async with websockets.connect("ws://127.0.0.1:3000/ws/hermes/cache") as websocket:
        # test set
        await websocket.send("set?key=test&ft=true&value=" + test_value())
        print(await websocket.recv())

        # test get
        await websocket.send("get?key=test")
        print(await websocket.recv())

        # close socket
        await websocket.close()

asyncio.run(test())