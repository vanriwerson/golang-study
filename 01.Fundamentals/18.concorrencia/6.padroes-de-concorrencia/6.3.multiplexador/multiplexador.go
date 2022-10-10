package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal1 := escrever("Olá, mundo!")
	canal2 := escrever("Aprendendo Golang")

	canalUnico := multiplexar(canal1, canal2)

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-canalUnico)
	}
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string { // junta N canais
	saida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-entrada1:
				saida <- mensagem
			case mensagem := <-entrada2:
				saida <- mensagem
			}
		}
	}()

	return saida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			randomTime := time.Duration(rand.Intn(2000))
			time.Sleep(time.Millisecond * randomTime)
		}
	}()

	return canal
}
