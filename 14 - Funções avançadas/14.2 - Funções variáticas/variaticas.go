package main

import "fmt"

func somar(numeros ...int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}

	return soma
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma := somar(10, 20, 14, 15, 20, 15, 100)
	fmt.Println(soma)

	escrever("Olá mundo", 10, 5, 4, 5, 8, 5 ,4 ,8, 10)
}
