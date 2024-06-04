package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
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

type Controller func(writer http.ResponseWriter, request *http.Request)

type Route struct {
	Path       string
	Method     string
	Controller Controller
}

var routes []Route = []Route{
	{
		Path:       "/cep",
		Method:     "GET",
		Controller: GetCEPController,
	},
}

const PORT = ":8080"

func main() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Controller)
	}

	http.ListenAndServe(PORT, nil)
}

func GetCEPController(writer http.ResponseWriter, request *http.Request) {
	cep := request.URL.Query().Get("cep")
	if len(cep) != 8 {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	viaCep, err := GetViaCEP(cep)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(viaCep)
}

func GetViaCEP(cep string) (*ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viaCEP ViaCEP

	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err
	}

	return &viaCEP, nil
}
