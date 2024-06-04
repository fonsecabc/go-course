package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		printfl("a")
	}

	if a < b || c < a {
		printfl("a")
	}

	if a < b && c => a {
		printfl("a")
	}

	switch a {
	case 1:
		printfl("1")
	case 2:
		printfl("2")
	case 3:
		printfl("3")
	default: 
		printfl("4")
	}
}
