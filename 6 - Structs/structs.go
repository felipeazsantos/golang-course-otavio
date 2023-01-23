package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "João"
	u.idade = 18
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua XYZ", 50}

	usuario2 := usuario{"Felipe", 23, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome:"Pedro"}
	fmt.Println(usuario3)
}
