package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const viaCEPUrl = "https://viacep.com.br/ws/%s/json/"

func main() {
	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf(viaCEPUrl, cep)
		req, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer req.Body.Close()

		var address Address
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(res, &address)
		if err != nil {
			panic(err)
		}

		fmt.Print(address)
	}
}
