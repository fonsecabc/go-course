package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "hello"

	ShowType(x)
	ShowType(y)
}

func ShowType(t interface{}) {
	fmt.Printf("The type of the variable is %T\n", t)
}
