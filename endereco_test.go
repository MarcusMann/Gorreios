package gorreios

import (
	"fmt"
	"testing"
)

func TestGetByCep(t *testing.T) {
	if _, err := BuscarViaCep(41820460); err != nil {
		t.Error(err)
	}

	fmt.Println("Concluído!")
}

func TestGetByEnderecop(t *testing.T) {
	if _, err := BuscarViaEndereco("BA", "Salvador", "Alameda"); err != nil {
		t.Error(err)
	}

	fmt.Println("Concluído!")
}
