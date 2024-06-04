package main

// Creates a copy of the varibles passed, any operation wont affect the original variables
func sum(a, b int) int {
	a = 50

	return a + b
}

// Points to the original variables memory address, any operation WILL affect the original varibles
func sumWithPointers(a, b *int) int {
	*a = 50

	return *a + *b
}

func main() {
	a := 10
	b := 120

	// Variable a keeps unchanged
	println(sum(a, b)) // Output: 170
	println(a)         // Output: 10

	// Pointers are passed to the function, changing the value store in the address if any operation is done
	println(sumWithPointers(&a, &b))
	println(a)

}
