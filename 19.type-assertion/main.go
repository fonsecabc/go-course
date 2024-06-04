package main

import "fmt"

func main() {
	var x interface{} = 10

	fmt.Println(IsNumber(x))
}

func IsNumber(x interface{}) bool {
	x, isNumber := x.(int)
	fmt.Println(x)

	return isNumber
}
