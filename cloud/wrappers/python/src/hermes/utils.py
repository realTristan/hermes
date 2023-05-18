import base64

def tob64(value: str):
     return base64.b64encode(value.encode("utf-8")).decode("utf-8")
