package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoValido := enderecos.TipoDeEndereco("Rua dos anjos")
	fmt.Println(tipoValido)
}
