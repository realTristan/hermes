import json

data: dict = json.load(open("data_hash.json"))
for k, v in data.items():
    data[k]["description"] = {
        "$hermes.full_text": True,
        "value": v["description"]
    }
    if "pre_requisites" in v:
        data[k]["pre_requisites"] = {
            "$hermes.full_text": True,
            "value": v["pre_requisites"]
        }
    data[k]["title"] = {
        "$hermes.full_text": True,
        "value": v["title"]
    }
    data[k]["name"] = {
        "$hermes.full_text": True,
        "value": v["name"]
    }

json.dump(data, open("data_hash.json", "w"), indent=4)