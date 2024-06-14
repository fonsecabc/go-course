package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Task %s: %d\n", name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	go task("B")

	fmt.Scanln()
}
