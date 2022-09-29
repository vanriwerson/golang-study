package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { // recover faz com que o programa possa continuar
		fmt.Println("Execução recuperada!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	// comentar essa linha fará o programa morrer durante a terceira execução
	defer recuperarExecucao() // Essa função será executada antes do panic
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXTAMENTE 6!") // interrompe o fluxo normal do programa após chamar todos os defer
	// o programa 'morre' após entrar em pânico
}

func main() {
	fmt.Println(alunoAprovado(8, 9))
	fmt.Println("Pós primeira execução")
	fmt.Println(alunoAprovado(3, 8))
	fmt.Println("Pós segunda execução")
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós pânico")
}

// panic NÃO é uma boa forma de tratar erros
