package main

import (
	"fmt"
)

func main() {
	// Step 1: Declare an array of int
	arr := [5]int{10, 20, 30, 40, 50}

	// Print the array before modification
	fmt.Println("Before modification:", arr)

	// Step 2: Create a pointer to the array
	ptr := &arr

	// Step 3: Modify the array elements using the pointer
	for i := range *ptr {
		(*ptr)[i] += 5 // Adding 5 to each element in the array
	}

	// Print the array after modification
	fmt.Println("After modification:", arr)
}

