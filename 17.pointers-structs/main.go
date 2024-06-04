package main

type Client struct {
	Name string
}

func (c *Client) ChangeName(name string) {
	c.Name = name
}

func main() {
	client := Client{
		Name: "Caio",
	}

	println(client.Name)

	client.ChangeName("Caio Filipe")
	println(client.Name)
}
