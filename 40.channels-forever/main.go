package main

func main() {
	forever := make(chan int)

	go func() {
		forever <- 42
	}()

	<-forever
}
