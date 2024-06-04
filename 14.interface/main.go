package main

import "fmt"

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

// This function recreates the struct client inside of it, not affecting the original struct
// func (c Client) Deactivate() {
// 	c.Active = false
// 	fmt.Print(c)
// }

// This function takes a pointer to the original struct client, thus affecthing the original struct
func (c *Client) Deactivate() {
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
	client.Deactivate()

}
