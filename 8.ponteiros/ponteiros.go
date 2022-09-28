package main

import "fmt"

func main() {
	fmt.Println("Ponteiros: referências de memória")

	var variavel1 int = 10
	var variavel2 int = variavel1 // uma cópia da variavel1 é atribuída à variavel2

	fmt.Println(variavel1, variavel2)

	variavel1++                       // até aqui, as variáveis são independentes
	fmt.Println(variavel1, variavel2) // apenas variavel1 será alterada

	// UTILIZANDO UM PONTEIRO
	var variavel3 int
	var ponteiro *int //declarando um ponteiro

	variavel3 = 100
	ponteiro = &variavel3             // atribuindo valor a um ponteiro
	fmt.Println(variavel3, ponteiro)  // retorna o endereço de memória salvo pelo ponteiro
	fmt.Println(variavel3, *ponteiro) // desreferenciação: mostra o valor salvo no ponteiro

	variavel3 += 50
	fmt.Println(variavel3, *ponteiro) // valor salvo no ponteiro acompanha as alterações da variavel3
}
