package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := Client{
		Name:   "Caio",
		Age:    18,
		Active: true,
	}

	fmt.Printf("Name: %v\nAge: %d\nActive: %v", client.Name, client.Age, client.Active)
}
