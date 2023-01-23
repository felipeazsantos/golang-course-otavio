package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Texto da função f")

	fmt.Println(result)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10,15)

	fmt.Println(resultadoSoma, resultadoSubtracao)
}
