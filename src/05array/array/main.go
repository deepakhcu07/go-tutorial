package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Define Array

	var numbers [5]int
	var cities [5]string
	var matrix [3][3]int // Array 2D

	// Insert Data

	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	// Insert Matrix Data

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// Display Data

	for i := 0; i < 5; i++ {
		fmt.Printf(" Number = %d", numbers[i])
		fmt.Printf(" Citi = %s", cities[i])
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
