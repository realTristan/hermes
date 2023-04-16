package main

import (
	"fmt"
	"time"
)

func main() {
	// allocate a slice the size of 4000
	startTime := time.Now()
	var _ []int = make([]int, 400000)
	fmt.Println("Time to allocate a slice of 4000: ", time.Since(startTime))

	// create a new slice with no allocation of memory and fill it with random numbers
	startTime = time.Now()
	var slice2 []int = []int{}
	for i := 0; i < 10000; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Println("Time to create a new slice with random numbers: ", time.Since(startTime))

	// Check if the number 27 is in the slice
	startTime = time.Now()
	for i := 0; i < len(slice2); i++ {
		if slice2[i] == 8977 {
			break
		}
	}
	fmt.Println("Time to check if 27 is in the slice: ", time.Since(startTime))

	// Create a new variable called number
	// and assign it the value of 27
	startTime = time.Now()
	var number int = 8977
	fmt.Println("Time to create a new variable and assign it the value of 8977: ", time.Since(startTime))
	fmt.Println(number)
}
