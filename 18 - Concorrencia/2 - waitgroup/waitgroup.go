package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2) // quantidade de 1 - goroutines para serem executadas em concorrencia

	go func () {
		escrever("Olá mundo!")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("Programando em GO!")
		waitgroup.Done() // -1
	}()

	waitgroup.Wait()

}

func escrever(texto string) {
	for i:= 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}