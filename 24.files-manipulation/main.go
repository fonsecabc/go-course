package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("The size of the string writen is %d\n", size)
	f.Close()

	f, err = os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
