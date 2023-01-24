package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome":"Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro":"renan",
			"segunda":"silva",
		},
		"curso": {
			"nome": "engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	// apagar dados pela chave map
	delete(usuario2, "nome")

	fmt.Println(usuario2)

	//adicionando chaves no map
	usuario2["signo"] = map[string]string {
		"nome":"sagitario",
	}
	fmt.Println(usuario2)

}
