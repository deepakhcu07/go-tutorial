package main

import (
	"bufio"
	"os"
)

func main() {
	// Creating a map counts of key as string and value as integer
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

}
