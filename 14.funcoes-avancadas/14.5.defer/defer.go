package main

import "fmt"

func func1() {
	fmt.Println("Executando a função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado:") // essa linha será executada antes do return
	fmt.Println("Iniciando cálculo da média")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	func1()
	func2()

	defer func1() // a cláusula defer adia a execução de um código até o último momento possível
	func2()

	fmt.Println(alunoAprovado(8, 10))
}

// Defer é comumente usado para fechar a conexão com um banco de dados
