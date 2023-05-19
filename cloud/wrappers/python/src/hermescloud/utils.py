import base64

class Utils:
    # Convert a string to base64
    @staticmethod
    def tob64(value: str):
        return base64.b64encode(value.encode("utf-8")).decode("utf-8")
