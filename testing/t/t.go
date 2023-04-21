package main

import (
	"fmt"
	"internal/bytealg"
)

func main() {
	fmt.Println(byte('a'))
	fmt.Println(byte('A'))

	fmt.Println(byte('z'))
	fmt.Println(byte('Z'))

	char := 'C'
	if char >= 65 && char <= 90 {
		fmt.Println("upper")
	}
}

func isUpper(char byte) bool {
	return char >= 65 && char <= 90
}

func charToLower(char byte) byte {
	if isUpper(char) {
		return char + 32
	}
	return char
}

func StringsContains(s, substr string) bool {
	return Index(s, substr) >= 0
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int {
	return bytealg.IndexByteString(s, c)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	var n int = len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexByte(s, substr[0])
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case n <= bytealg.MaxLen:
		// Use brute force when s and substr both are small
		if len(s) <= bytealg.MaxBruteForce {
			return bytealg.IndexString(s, substr)
		}
		var (
			c0    byte = substr[0]
			c1    byte = substr[1]
			i     int  = 0
			t     int  = len(s) - n + 1
			fails int  = 0
		)
		for i < t {
			if s[i] != c0 {
				// IndexByte is faster than bytealg.IndexString, so use it as long as
				// we're not getting lots of false positives.
				var o int = IndexByte(s[i+1:t], c0)
				if o < 0 {
					return -1
				}
				i += o + 1
			}
			if s[i+1] == c1 && s[i:i+n] == substr {
				return i
			}
			fails++
			i++
			// Switch to bytealg.IndexString when IndexByte produces too many false positives.
			if fails > bytealg.Cutover(i) {
				var r int = bytealg.IndexString(s[i:], substr)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}
	var (
		c0    byte = substr[0]
		c1    byte = substr[1]
		i     int  = 0
		t     int  = len(s) - n + 1
		fails int  = 0
	)
	for i < t {
		if s[i] != c0 {
			var o int = IndexByte(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i+1] == c1 && s[i:i+n] == substr {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// See comment in ../bytes/bytes.go.
			var j int = bytealg.IndexRabinKarp(s[i:], substr)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}
