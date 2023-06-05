package main

import (
	"fmt"

	hermes "github.com/realTristan/hermes"
)

func ft() {
	var cache *hermes.Cache = hermes.InitCache()

	// Initialize the FT cache
	cache.FTInit(-1, -1, 3)

	// print cache info
	fmt.Println(cache.InfoForTesting())
}
