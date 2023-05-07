import json

data: list = json.load(open("data_array.json"))
for i, item in enumerate(data):
    data[i]["description"] = {
        "$hermes.full_text": True,
        "value": item["description"]
    }
    if "pre_requisites" in item:
        data[i]["pre_requisites"] = {
            "$hermes.full_text": True,
            "value": item["pre_requisites"]
        }
    data[i]["title"] = {
        "$hermes.full_text": True,
        "value": item["title"]
    }
    data[i]["name"] = {
        "$hermes.full_text": True,
        "value": item["name"]
    }

json.dump(data, open("data_array.json", "w"), indent=4)