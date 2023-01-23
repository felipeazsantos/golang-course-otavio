package main

import "fmt"

func main() {
	var numero int64 = 10000000000
	fmt.Println(numero)

	// somente números positivos
	var numero2 uint32 = 100
	fmt.Println(numero2)

	// alias
	var numero3 rune = 12345 //int32
	fmt.Println(numero3)

	var numero4 byte = 12 // uint8
	fmt.Println(numero4)

	var numeroReal1 float32 = 152.89
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.57
	fmt.Println(numeroReal2)

	numeroReal3 := 122345.67
	fmt.Println(numeroReal3)

	//STRINGS
	var str string = "lalalalla"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//FIM Strings

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}
