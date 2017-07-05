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

var url correios

// EnderecoLista retorna uma lista de endereços
type EnderecoLista []Endereco

// BuscarViaCep recupera os dados passando o cep como argumento.
func BuscarViaCep(cep uint32) (*Endereco, error) {
	var endereco Endereco

	if err := unmarshalCorreios(url.GetCepURL(cep), &endereco); err != nil {
		return nil, err
	}

	return &endereco, nil
}

// BuscarViaEndereco retorna uma lista dos endereços
func BuscarViaEndereco(estado, cidade, bairro string) (*EnderecoLista, error) {
	var enderecoLista EnderecoLista

	if err := unmarshalCorreios(url.GetAddressURL(estado, cidade, bairro), &enderecoLista); err != nil {
		return nil, err
	}

	return &enderecoLista, nil
}
