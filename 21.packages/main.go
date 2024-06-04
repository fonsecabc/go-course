package main

import (
	"course-package/math"
	"fmt"
)

func main() {
	x, y := 12, 13

	sum := math.Sum(x, y)
	fmt.Println(sum)
}
