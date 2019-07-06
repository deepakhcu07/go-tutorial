package main

import "fmt"

func main() {
	// Declare Variables

	var str string
	var n, m int
	var x float32

	// Assign Values

	str = "Welcome in Golang"
	n = 20
	m = 30

	x = 3.14

	fmt.Println(" The value of str is ", str)
	fmt.Println(" The value of n is ", n)
	fmt.Println(" The value of m is ", m)
	fmt.Println(" The value of x is ", x)

	// Define Multiple Variables
	var (
		name  string
		email string
		age   int
	)

	name = "Aditya Singh"
	email = "aditya@gmail.com"
	age = 11

	fmt.Println(" Name  ", name)
	fmt.Println(" Email", email)
	fmt.Println(" Age ", age)

	// Short Variable Declarations
	val := 15
	fmt.Printf(" The value of val is %v and Type %T", val, val)

}
