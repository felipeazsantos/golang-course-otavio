package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// goroutines -> executa a função mas não espera o retorno para continuar a execução dos próximos códigos
	// Funções assincronas
	go escrever("Olá mundo!")
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}