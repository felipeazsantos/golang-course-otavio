package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo!"), escrever("Programando em GO!"))

	for i := 0; i < 10; i ++ {
		fmt.Println(<-canal)
	}
}


func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	CanalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagemCanal1 := <-canalDeEntrada1:
				CanalDeSaida <- mensagemCanal1
			case mensagemCanal2 := <-canalDeEntrada2:
				CanalDeSaida <- mensagemCanal2
			}
		}
	}()

	return CanalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}