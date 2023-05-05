import requests


def get(addr: str):
    r = requests.get(addr)
    return r.json()
