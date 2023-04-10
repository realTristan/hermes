import json

# // The cache class
class Cache:
    def __init__(self) -> None:
        self.cache: dict = {}
        self.data: dict = json.load(open("data.json", "r"))

    # // Load the cache from the data.json file
    def load(self) -> None:
        for i, course in enumerate(self.data):
            for _, v in course.items():
                # // Split the string by spaces, then iterate over the words
                words: list[str] = v.lower().strip().split()
                for word in words:
                    # // If the word is not in the cache, add it
                    if word not in self.cache:
                        self.cache[word] = []
                    
                    # // Append the index of the item for this word to the cache
                    self.cache[word].append(i)

    # // Search for a word in the cache
    def search(self, word: str) -> list[int]:
        res: list[int] = []
        for k, v in self.cache.items():
            if word in k:
                res.extend(v)
        return res
    
    # // Convert the indices to the actual items
    def indices_to_data(self, indices: list[int]) -> list[dict]:
        items: list[dict] = []
        for i in indices:
            items.append(self.data[i])
        return items
