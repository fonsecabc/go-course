package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 1, 50

	totalSum, err := sum(a, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The total sum of a and b is %d\n", totalSum)
}

func sum(a, b int) (int, error) {
	sum := a + b
	if sum > 50 {
		return 0, errors.New("The sum is greater than 50")
	}

	return sum, nil
}
