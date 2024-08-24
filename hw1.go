package main

import (
	"fmt"
)

func main() {
	// Step 1: Create a slice of int
	nums := []int{1, 2, 3, 4, 5}

	// Print the slice before modification
	fmt.Println("Before modification:", nums)

	// Step 2: Create a pointer to the slice
	ptr := &nums

	// Step 3: Modify the slice elements through the pointer
	for i := range *ptr {
		(*ptr)[i] *= 2 // Doubling each element in the slice
	}

	// Print the slice after modification
	fmt.Println("After modification:", nums)
}

