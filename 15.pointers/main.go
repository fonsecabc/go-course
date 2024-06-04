package main

import "fmt"

func main() {
	a := 10

	// Prints the value of the variable
	fmt.Println(a)

	// Prints the address of the variable in the memory
	fmt.Println(&a)

	// Defining a pointer, a variable that has the memory address of another variable as its value
	pointer := &a
	fmt.Println(pointer)

	// Returns the value stored in that memory address
	*pointer = 12
	fmt.Println(a)
}
