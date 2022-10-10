package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá, Generator!")

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string { // generator encapsula a goroutine e retorna um canal de comunicação
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
