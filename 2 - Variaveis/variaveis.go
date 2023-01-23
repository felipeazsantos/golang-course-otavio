package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	const contante1 string = "Constante 1"
	fmt.Println(contante1)
}
