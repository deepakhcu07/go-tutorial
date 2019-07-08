package main

import "fmt"

func main() {

	avg := average(2, 4, 6, 8)

	fmt.Printf("Average = %.2f \n", avg)
}

func average(nums ...float64) float64 {

	// For Range Loop
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	avg := sum / (float64)(len(nums))
	return avg
}
