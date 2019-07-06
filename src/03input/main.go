package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Enter Radius: ")
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("Area = %.2f", area)
}
