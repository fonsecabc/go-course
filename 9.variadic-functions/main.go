package main

import (
	"fmt"
)

func main() {
	a, b := 1, 50

	totalSum := sum(a, b)
	fmt.Printf("The total sum of a and b is %d\n", totalSum)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
