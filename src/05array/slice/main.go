package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Define Slice

	var numbers []int
	numbers = make([]int, 5)

	matrix := make([][]int, 3*3)

	// Insert Data

	for i := 0; i < 3; i++ {
		numbers[i] = rand.Intn(100)
	}

	// Insert Matrix

	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3) // Creating another Slice here
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// Displaying Data

	for i := 0; i < 5; i++ {
		fmt.Printf(" Number = %d", numbers[i])
	}

	// Display Matrix

	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %3d ", matrix[i][j])
		}
		fmt.Println()
	}

}
