package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter a number : ")
	fmt.Scanf("%d", &num)

	fmt.Printf("You have entered %d \n", num)

	// Loop to print table

	for i := 1; i <= 10; i++ {
		fmt.Printf(" %3d  X  %3d  = %5d \n ", num, i, (num * i))
	}

}
