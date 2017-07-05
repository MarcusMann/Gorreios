package gorreios

// Endereco recupera o endereço através do cep passado
type Endereco struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
	Unidade     string
	Ibge        string
	Gia         string
}

// EnderecoLista retorna uma lista de endereços
type EnderecoLista []Endereco

// GetCep recupera os dados passando o cep como argumento.
func GetByCep(cep uint32) (*Endereco, error) {
	var endereco Endereco
	var url url

	if err := unmarshalCorreios(url.GetCepURL(cep), &endereco); err != nil {
		return nil, err
	}

	return &endereco, nil
}

// GetEndereco retorna uma lista dos endereços
func GetByEndereco(estado, cidade, bairro string) (*EnderecoLista, error) {
	var enderecoLista EnderecoLista
	var url url

	if err := unmarshalCorreios(url.GetAddressURL(estado, cidade, bairro), &enderecoLista); err != nil {
		return nil, err
	}

	return &enderecoLista, nil
}
