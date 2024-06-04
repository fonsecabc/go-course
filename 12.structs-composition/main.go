package main

import "fmt"

type Address struct {
	City    string
	State   string
	Country string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	client := Client{
		Name:   "Caio",
		Age:    18,
		Active: true,
	}

	client.City = "São Paulo"

	fmt.Printf("Name: %v\nAge: %d\nActive: %v", client.Name, client.Age, client.Active)
}
