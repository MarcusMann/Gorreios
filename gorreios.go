package gorreios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type correios string

func (correios) GetCepURL(cep uint32) string {
	return fmt.Sprintf("https://viacep.com.br/ws/%08d/json/", cep)
}

func (correios) GetAddressURL(endereco ...string) string {
	return fmt.Sprintf("https://viacep.com.br/ws/%s/%s/%s/json/", endereco[0], endereco[1], endereco[2])
}

// UnmarshalCorreios analisa os dados codificados em JSON e armazena o resultado no valor apontado por v
func unmarshalCorreios(URL string, v interface{}) error {
	client := &http.Client{}

	// Faz a requisição dos dados que irão ser recebidos via Get, caso não consiga receber, fecha a requisição e retorna um erro.
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}

	// Seta que a requisição tem que ser via Json
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// fecha a response e retorna um erro
	defer response.Body.Close()

	// Se os status recebidos for diferente de 200, fecha o corpo e retorna um erro
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("ocorreu um erro! StatusCode: %d", response.StatusCode)
	}

	// Ler todos os dados recebidos e retorna um []bytes e caso haja um erro, fecha a resposta e retorna o erro.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Faz um parse entre os dados recebidos e interface passada no argumento.
	return json.Unmarshal(body, &v)
}
