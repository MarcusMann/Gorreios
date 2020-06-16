## Instalação
Você pode obter o pacote rodando o comando abaixo:

```go
go get github.com/marcuxyz/gorreios
```

## Como usar:
Você pode importar o pacote desta forma:

```go
import "github.com/marcuxyz/gorreios"
```

E para obter um endereço através de um cep você pode usar o código a seguir:

```go
package main

import (
  "fmt"
  "github.com/marcuxyz/gorreios"
)

func main() {
  address := gorreios.GetAddressByCep("41820-460")
  fmt.Printlnf("%+v\n", address)
}
```
