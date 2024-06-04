package main

import "fmt"

// A map of key of type string and value of type int
type Salaries map[string]int

func main() {
	// A variable of type Salaries
	var salaries = Salaries{
		"Wesley": 1000,
		"John":   2000,
		"Maria":  3000,
	}

	fmt.Println(salaries)

	// Deleting a key from a map
	delete(salaries, "Wesley")
	fmt.Println(salaries)

	// Adding a new key to a map
	salaries["Kyle"] = 5000
	fmt.Println(salaries)

	// Creating an empty map
	emptyMap := make(map[string]int) // map[string]in{}
	fmt.Println(emptyMap)

	// Creating an empty map of type Salaries
	emptySalaries := make(Salaries) // Salaries{}
	fmt.Println(emptySalaries)

	// Looping through the map
	for key, value := range salaries {
		fmt.Printf("%v's salary is $%d\n", key, value)
	}

	// To ignore a variable, use the _ as the blank identifier
	for _, value := range salaries {
		fmt.Printf("Salary is $%d\n", value)
	}
}
