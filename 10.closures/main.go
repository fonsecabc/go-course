package main

import (
	"fmt"
)

func main() {
	a, b := 1, 50

	total := sum(a, b)
	fmt.Printf("The total sum of a and b is %d\n", total)

	x := 2
	totalMultiplied := func() int {
		return total * x
	}()
	fmt.Printf("The total mutiplication is %d\n", totalMultiplied)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
