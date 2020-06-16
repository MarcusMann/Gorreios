package gorreios

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetAddressByCep Get address via cep
func GetAddressByCep(cep string) *Address {
	address := new(Address)

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	checkError(err)

	err = json.NewDecoder(resp.Body).Decode(&address)
	checkError(err)

	return address
}
