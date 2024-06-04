package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func SumSalaries[T Number](salaries map[string]T) T {
	var sum T

	for _, value := range salaries {
		sum += value
	}

	return sum
}

func CompareSalaries[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// salaries := map[string]int{
	// 	"Wesley":   1000,
	// 	"Fernando": 2000,
	// 	"Marcos":   3000,
	// }

	salaries2 := map[string]MyNumber{
		"Wesley":   1000,
		"Fernando": 2000,
		"Marcos":   3000,
	}

	fmt.Printf("The sum of the salaries is %d\n", SumSalaries(salaries2))
	fmt.Printf("The salaries are of equal type? %v", CompareSalaries(salaries2["Wesley"], salaries2["Fernando"]))
}
