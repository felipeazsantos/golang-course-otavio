package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//Go não tem while ou do while, somente for
	//usando o for parecido com while
	i := 0
	for i < 10 {
		time.Sleep(time.Millisecond)
		fmt.Println("Incrementando i ", i)
		i++
	}

	//for padrão
	for j := 0; j < 10; j++{
		time.Sleep(time.Millisecond)
		fmt.Println("Incrementando j ", j)
	}

	// range (estilo foreach)
	fmt.Println("---------------------------")
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Davi",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome string
		sobrenome string
	}

	//Não da pra usar range com struct
	//usuario2 := usuarioStruct{"Zé", "dos Santos"}
	//
	//for chave, valor := range usuario2 {
	//
	//}
}
