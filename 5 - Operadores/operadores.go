package main

import "fmt"

func main() {
	// ARITMÉTICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10/4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1, variavel2)

	// FIM ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM LOGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10 // numero = numero / 10
	numero %= 3

	fmt.Println(numero)

	// FIM UNARIOS

}
