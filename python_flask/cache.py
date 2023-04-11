# //////////////////////////////////////////////////////////////////////////////////////
# //
# // Note: The average size of the cache map is about 1.3 MB for the data.json file
# //       Using this full-text-search method for very large datasets is not recommended
# //
# //////////////////////////////////////////////////////////////////////////////////////
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
                # // Remove all double spaces from the string
                while "  " in v:
                    v = v.replace("  ", " ")

                # // Split the string by spaces, then iterate over the words
                words: list[str] = v.lower().strip().split()
                for word in words:
                    # // If the word is not all alpha, skip it
                    if not word.isalnum() or len(word) <= 1:
                        continue

                    # // If the word is not in the cache, add it
                    if word not in self.cache:
                        self.cache[word] = []

                    # // If the index is already in the cache, skip it
                    if i in self.cache[word]:
                        continue
                    
                    # // Append the index of the item for this word to the cache
                    self.cache[word].append(i)

    # // Search for a word in the cache
    def search(self, word: str, limit: int, strict: bool) -> list[int]:
        res: list[dict] = []

        # // If strict mode is enabled
        if strict:
            # // Check if the word is in the cache
            if word not in self.cache:
                return res

            # // Iterate over the indices in the cache
            for i in self.cache[word]:
                res.append(self.data[i])
            
            # // Return the results
            return res
        
        already_added: list[int] = []
        for k, v in self.cache.items():
            # // Check if the limit has been reached
            if len(res) >= limit:
                return res
            
            # // The word doesn't start with the same letter
            if k[0] != word[0]:
                continue
            
            # // The word is longer than the key string
            if len(k) < len(word):
                continue

            # // The word is the key string
            if word == k:
                for i in v:
                    if i in already_added:
                        continue
                    res.append(self.data[i])
                return res
            
            # // The word is not in the key string
            if word not in k:
                continue
            
            # // Iterate over the indices in the cache
            for i in v:
                if i in already_added:
                    continue
                res.append(self.data[i])
                already_added.append(i)
        
        # // Return the results
        return res
    