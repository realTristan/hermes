import requests, json, base64, time

def base64_encode(value):
    return base64.b64encode(value.encode("utf-8")).decode("utf-8")

def test_search():
    headers = {
        "Content-Type": "application/json"
    }

    # schema
    schema = base64_encode(json.dumps({
        "id":             False,
        "components":     False,
        "units":          False,
        "description":    True,
        "name":           True,
        "pre_requisites": True,
        "title":          True
    }))

    # search for a value
    url = "http://localhost:3000/ft/search"
    params = "?query=computer&strict=false&limit=100&schema=" + schema

    # make the request
    r = requests.get(url+params, headers=headers)
    print(r.text)

def test_set_json(data):
    headers = {
        "Content-Type": "application/json"
    }

    # set the data
    for key in data:
        # set the key
        url = "http://localhost:3000/cache/set"
        params = "?key=" + key + "&value=" + base64_encode(json.dumps(data[key]))

        # make the request
        r = requests.post(url+params, headers=headers)
        #print(r.text)

def test_set():
    headers = {
        "Content-Type": "application/json"
    }

    # set the data
    url = "http://localhost:3000/cache/set"
    params = "?key=testing&value=" + base64_encode(json.dumps({
        "name": "tristan"
    }))

    # make the request
    r = requests.post(url+params, headers=headers)
    print(r.text)


def test_init_ft_json(data):
    headers = {
        "Content-Type": "application/json"
    }

    # url and params
    _json = base64_encode(json.dumps(data))
    url = "http://localhost:3000/ft/init/json"
    params = "?maxlength=-1&maxbytes=-1&json=" + _json

    # make the request
    r = requests.post(url+params, headers=headers)

    # print the response
    print(r.text)


def test_init_ft():
    headers = {
        "Content-Type": "application/json"
    }

    # url and params
    url = "http://localhost:3000/ft/init"
    params = "?maxlength=-1&maxbytes=-1"

    # make the request
    r = requests.post(url+params, headers=headers)

    # print the response
    print(r.text)


def test_get():
    headers = {
        "Content-Type": "application/json"
    }

    # get b4a3261059ea6f1b48eb8039e720e0b48d087583
    url = "http://localhost:3000/cache/get"
    params = "?key=b4a3261059ea6f1b48eb8039e720e0b48d087583"

    # make the request
    r = requests.get(url+params, headers=headers)
    print(r.text)

def test_cache_info():
    headers = {
        "Content-Type": "application/json"
    }

    # get b4a3261059ea6f1b48eb8039e720e0b48d087583
    url = "http://localhost:3000/cache/info/testing"

    # make the request
    r = requests.get(url, headers=headers)
    print(r.text)


if __name__ == "__main__":
    # read the json file from the data folder
    with open("data/data_hash.json", "r") as file:
        # load the json file
        data = json.loads(file.read())
        # test_init_ft_json(data)
        test_init_ft()
        # test_cache_info()
        # test_get()
        # test_set_json(data)
        # test_set()
        # test_cache_info()

        st = time.time()
        test_search()
        print(time.time() - st)