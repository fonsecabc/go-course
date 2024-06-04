package main

import (
	"fmt"
)

func main() {
	// Initial capacity of 5 and length of 5
	// {1, 2, 3, 4, 5}
	list := []int{1, 2, 3, 4, 5}

	// PrintLenAndCap(list)

	// Using :X to show X values after the start of the slice

	// PrintLenAndCap(list[:2])

	// Using X: to ignore X values after the start of the slice, removing X from the slice capacity

	// PrintLenAndCapOfList(list[2:])

	// Using append we can add a value to the end of the slice, where GO recreates the array with twice the capacity of the original one, and adds the values to it
	// Capacity of 10 and length of 6
	// {1, 2, 3, 4, 5, 120}
	list = append(list, 120)
	PrintLenAndCapOfList(list)
}

func PrintLenAndCapOfList(list []int) {
	fmt.Printf("The contents for the list are %v\n", list)
	fmt.Printf("The length of the list is %d\n", len(list))
	fmt.Printf("The capacity of the list is %d\n", cap(list))
}
