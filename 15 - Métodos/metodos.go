package main

import "fmt"

type usuario struct {
	nome string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	u := usuario{"Pedro", 41}
	u.salvar()

	u2 := usuario{"Davi", 30}
	u2.salvar()

	maiorDeIdade := u2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	u2.fazerAniversario()
	fmt.Println(u2.idade)
}
