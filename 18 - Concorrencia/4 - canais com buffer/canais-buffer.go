package main

import "fmt"

func main() {
	// buffer -> canal somente bloqueia a execução quando atingir o buffer
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
