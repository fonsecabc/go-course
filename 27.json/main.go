package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number int `json:"n"`
}

func main() {
	account := Account{Number: 1}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		panic(err)
	}

	var account2 Account
	jsonData := []byte(`{"n": 2}`)

	err = json.Unmarshal(jsonData, &account2)
	if err != nil {
		panic(err)
	}

	fmt.Print(account2.Number)
}
