package main

import "fmt"

/*
	O exemplo abaixo retornará o número que ocupa a posição N da Sequência de Fibonacci, que é uma
	sequência de números inteiros, começando normalmente por 0 e 1, na qual cada termo subsequente
	corresponde à soma dos dois anteriores.
*/
func fibonacci(posicao uint) uint {
	if posicao <= 1 { // condição de parada
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// Essa função recursiva calcula o fatorial de um número
func fatorial(num uint) uint {
	if num == 0 {
		return 1
	}

	return num * fatorial(num-1)
}

func main() {
	// Fibonacci: 1 1 2 3 5 8 13 21 34 55 89
	posicao := uint(12)

	fmt.Println("Fibonacci, posição:", posicao)
	fmt.Println("Elemento:", fibonacci(posicao)) // 144

	// Imprimindo N elementos da Sequência de Fibonacci
	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	num := uint(7)
	fmt.Println("Fatorial de 7:", fatorial(num)) // 5040
}

// Funções recursivas não são recomendadas para casos que exigem muitas execuções
