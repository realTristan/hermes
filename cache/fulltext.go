package cache

/*
FullText represents a data structure for storing and searching full-text documents. It uses a map to cache word positions for each word in the text, and an array to store the keys in sorted order for faster iteration. The maximum number of words that can be cached and the maximum size of the text can be set when initializing the struct.

Fields:
- wordCache: a map of words to an array of positions where the word appears in the text
- keys: an array of keys sorted in ascending order
- maxWords: the maximum number of words to cache
- maxSizeBytes: the maximum size of the text to cache in bytes
- isInitialized: a boolean flag indicating whether the struct has been initialized
*/
type FullText struct {
	wordCache     map[string][]int
	keys          []string
	maxWords      int
	maxSizeBytes  int
	isInitialized bool
}
