package main

import (
	"fmt"
	"os"
)

func main() {
	// Iterating through the command line arguments using for loop

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
