package main

import (
	"fmt"
)

var (
	list [3]int
)

func main() {
	list[0] = 10
	list[1] = 20
	list[2] = 30

	lastItem := list[len(list)-1]

	fmt.Println(list)
	fmt.Println(lastItem)

	for i, v := range list {
		fmt.Printf("The index is %d has the value of %d\n", i, v)
	}
}
