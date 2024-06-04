package main

import (
	"fmt"
)

type ID int

var (
	userId ID = 1
)

func main() {
	fmt.Printf("User Id type of %T ", userId)
}
