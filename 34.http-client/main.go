package main

import (
	"bytes"
	"net/http"
)

func main() {
	client := http.Client{}
	jsonBuffer := bytes.NewBuffer([]byte(`{"cep":"01001-000"}`))

	resp, err := client.Post("http://localhost:8080", "application/json", jsonBuffer)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
