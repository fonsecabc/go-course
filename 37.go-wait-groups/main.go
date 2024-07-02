package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Task %s: %d\n", name, i)
		time.Sleep(time.Second)
		waitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	go task("A", &waitGroup)
	go task("B", &waitGroup)

	waitGroup.Wait()
}
