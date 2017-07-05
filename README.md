Instalação
========

Para instalar o gorreios execute o comando abaixo em seu terminal:

```go
go get github.com/MarcusMann/Gorreios
```

Para chamar a aplicação dentro do seu arquivo **main.go** escreva o código abaixo:

```go
package main

import (
	"fmt"

	"github.com/MarcusMann/gorreios"
)

func main() {
	correios, err := gorreios.BuscarViaEndereco("BA", "Salvador", "Pituba")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(correios)
}
```
Ao executar o comando  `go run main.go` ele irá trazer todas as informações possíveis, dos bairros encontrados através do Estado, Cidade e Logradouro. A busca também pode ser realizada por cep, veja:

```go
...
func main() {
	correios, err := gorreios.BuscarViaCep(41820460)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println(correios)
}
```
Uma outra forma de obter o resultado esperado é chamando a aplicação dessa forma:
```go
package main

import (
	_ "github.com/MarcusMann/gorreios"
)

func main() {

}
```

Agora você pode executar através de flags, veja:
```go
go run main.go -cep=41820460
```

Veja um pequeno vídeo abaixo de como funcionar a questão das flags.

[![asciicast](https://asciinema.org/a/DQDIWanuTHPzEuWBQ7RoLFf27.png)](https://asciinema.org/a/DQDIWanuTHPzEuWBQ7RoLFf27)

Você pode ver o vídeo completo diretamente no meu canal do [**youtube**](https://youtu.be/i0hz-Q-xjXs).
