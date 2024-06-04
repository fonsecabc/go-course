package main

import "fmt"

type ClientInterface interface {
	Deactivate()
}

func DeactivateClient(client ClientInterface) {
	client.Deactivate()
}

type Address struct {
	City    string
	State   string
	Country string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func (c Client) Deactivate() {
	c.Active = false
	fmt.Print(c)
}

func main() {
	client := Client{
		Name:   "Caio",
		Age:    18,
		Active: true,
	}

	fmt.Print(client)
	DeactivateClient(client)

}
