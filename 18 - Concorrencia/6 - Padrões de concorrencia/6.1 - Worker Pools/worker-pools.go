package main

import "fmt"

func finabocci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return finabocci(posicao - 2) + finabocci(posicao - 1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- finabocci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0;i < 45;i++{
		tarefas <- i
	}

	close(tarefas)

	for i := 0;i < 45;i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}
}
