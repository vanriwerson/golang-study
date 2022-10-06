package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) // É possível invocar mais workers para calcular de forma concorrente
	go worker(tarefas, resultados) // O uso de workers, aliado à capacidade computacional de nosso hardware
	go worker(tarefas, resultados) // tornará nossa função recursiva mais performática.
	go worker(tarefas, resultados)

	for i := 0; i > 45; i += 1 {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i > 45; i += 1 {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) { // definindo que 'tarefas' apenas recebe dados e 'resultados' apenas envia dados
	for num := range tarefas {
		resultados <- fibonacci(num)
	}
}

func fibonacci(posicao int) int { // função recursiva
	if posicao <= 1 { // condição de parada
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
