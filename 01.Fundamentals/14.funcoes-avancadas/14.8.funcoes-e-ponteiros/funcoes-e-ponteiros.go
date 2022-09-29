package main

import "fmt"

// passando parâmetro por cópia
func inverterSinal(num int) int {
	return num * -1
}

// passando parâmetro por referência
func inverterSinalComPonteiro(num *int) { // aponta para o endereço de memória onde está o parâmetro
	*num = *num * -1 // *num faz a desreferenciação do endereço de memória, entregando o valor
}

func main() {
	numero := 20
	fmt.Println(numero)
	comSinalInvertido := inverterSinal(numero) // a função cria uma nova variável a partir de numero
	fmt.Println(comSinalInvertido)
	fmt.Println(numero)

	outroNumero := 40
	fmt.Println(outroNumero)
	// utilizando ponteiros, a variável outroNumero é alterada diretamente
	inverterSinalComPonteiro(&outroNumero) // & informa que estamos passando um endereço de memória como parâmetro
	fmt.Println(outroNumero)               // aqui notamos que a variável teve seu valor alterado pela função
}
