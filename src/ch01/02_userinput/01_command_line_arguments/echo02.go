package main

import (
	"fmt"
	"os"
)

func main() {
	// Iterating over the command line arguments using for range

	for i, v := range os.Args[1:] {
		fmt.Println("Argument ", v, " found at Index ", i)
	}
}
