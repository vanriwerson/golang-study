package main

import "fmt"

// Go é uma linguagem fortemente tipada
func main() {
	var variavel1 string = "Tipo (string) declarado explícitamente"
	fmt.Println(variavel1)

	variavel2 := "Tipo (string) declarado implícitamente" // inferência de tipo
	fmt.Println(variavel2)

	var (
		variavel3 string = "Declaração simultânea"
		variavel4 string = "de várias variáveis"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Declaração simultânea", "por inferência de tipo"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Essa é uma constante"

	fmt.Println("Invertendo os valores de variavel5 e variavel6")
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
