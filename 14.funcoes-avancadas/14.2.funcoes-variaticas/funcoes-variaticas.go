package main

import "fmt"

func soma(nums ...int) int { // permite que a função receba um número indefinido de parâmetros
	// fmt.Println(nums) imprime um slice com os parâmetros passados
	resultado := 0
	for _, num := range nums {
		resultado += num
	}

	return resultado
}

func escrever(texto string, nums ...int) { // só pode haver um parâmetro váriatico por função e ele deve ser declarado por último
	for _, num := range nums {
		fmt.Println(texto, num)
	}
}

func main() {
	fmt.Println("Total da soma:", soma(6, 7, 8, 9))

	escrever("Texto repetido", 1, 2, 3)
}
