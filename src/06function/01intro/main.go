package main

import "fmt"

func main() {

	foo()

	greet("Aditya")

	add(10, 5)

	x := subtract(19, 5)
	fmt.Printf(" %d - %d = %d \n", 19, 5, x)

	a, b, c := compute(5, 8, 9, "test")

	fmt.Printf(" a= %f \n b = %f \n c = %s \n", a, b, c)

}

// Function Declaration which takes no argument and no returns
func foo() {
	fmt.Println("This is foo ...")
}

// A function declaration with parameters

func greet(name string) {
	fmt.Println("We are inside Greet")
	fmt.Printf(" Hello %s \n", name)
}

// Passing multiple parameters

func add(a, b int) {
	result := a + b
	fmt.Printf("Sum = %d\n ", result)
}

// A function with parameters and return value

func subtract(a, b int) int {
	return a - b
}

// A function with parameters and multiple return values

func compute(a, b, c float64, name string) (float64, float64, string) {
	result1 := a * b * c
	result2 := a + b + c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s value = %.2f", name, result3)

	return result1, result2, newName
}
