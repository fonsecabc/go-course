package main

import "sync"

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(10)

	go reader(c, &wg)
	go publisher(c)

	wg.Wait()
}

func reader(c chan int, wg *sync.WaitGroup) {
	for x := range c {
		println(x)
		wg.Done()
	}
}

func publisher(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
