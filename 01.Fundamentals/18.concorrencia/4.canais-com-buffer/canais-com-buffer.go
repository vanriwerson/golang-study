package main

import "fmt"

func main() {
	canal := make(chan string, 2) // o canal com buffer só bloqueia a execução quando atinge a capacidade máxima
	canal <- "Olá, Mundo!"
	canal <- "Esse envio faz o canal chegar à capacidade informada (2)"
	// canal <- "Esse envio causaria um deadlock"

	mensagem := <-canal
	mensagem2 := <-canal
	// mensagem3 := <-canal // o valor para essa variável nunca chegaria

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
