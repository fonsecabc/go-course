package main

func main() {
	c := make(chan int)

	go reader(c)
	publisher(c)

	close(c)
}

func reader(c chan int) {
	for x := range c {
		println(x)
	}
}

func publisher(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
