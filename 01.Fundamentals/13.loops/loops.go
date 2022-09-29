package main

import (
	"fmt"
)

// Em Go temos apenas o loop 'for' que pode ser usado de diversas maneiras
func main() {
	a := 0 // análogo ao while
	for a < 10 {
		a++
		// time.Sleep(time.Second) // pausa a execução por 1 segundo
		fmt.Println("a =", a)
	}
	fmt.Println("último valor de a:", a)

	for j := 0; j < 10; j++ { // for raiz ^^
		fmt.Println("j =", j) // a variável j só existe no escopo do for
	}

	array := [4]string{"George", "John", "Paull", "Ringo"}

	for index, value := range array { // usado para percorrer elementos iteráveis
		fmt.Println(index, value)
	}
	for _, value := range array {
		fmt.Println(value) // imprime somente os nomes
	}

	for i, char := range "PALAVRA" { // range itera pelos caracteres de uma string
		fmt.Println(i, char)         // retorna os códigos da tabela ASCII
		fmt.Println(i, string(char)) // retorna as letras
	}

	usuario := map[string]string{
		"nome":      "Eddie",
		"sobrenome": "Van Hallen",
	}

	for chave, valor := range usuario { // range itera pelas chaves e valores de um map
		fmt.Println(chave, valor)
	}

	// Não é possivel utilizar o range em um struct

	// Loop Infinito (não façam isso em casa, crianças):
	// for {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Executando infinitamente")
	// }
}
