package main

type ID int
type Name string

var (
	userId   ID   = 1
	userName Name = "Caio"
)

func main() {
	println(userId)
	println(userName)
}
