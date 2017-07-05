package gorreios

import (
	"flag"
	"fmt"
	"strings"
)

func init() {
	cep := flag.Uint64("cep", 0, "-cep=41820460 faz a busca via cep")
	endereco := flag.String("endereco", "", "-endereco faz a busca via endereco")

	flag.Parse()

	if *cep != 0 {
		bcep, err := BuscarViaCep(uint32(*cep))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(bcep)
	}

	enderecos := strings.Split(*endereco, ",")

	if *endereco != "" {
		bendereco, err := BuscarViaEndereco(enderecos[0], enderecos[1], enderecos[2])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(bendereco)
	}
}
