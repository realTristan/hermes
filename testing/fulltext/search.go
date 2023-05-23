package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes"
)

func search() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache
	cache.FTInit(-1, -1, 3)

	// print cache info
	fmt.Println(cache.InfoForTesting())
}
