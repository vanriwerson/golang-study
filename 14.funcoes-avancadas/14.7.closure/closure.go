package main

import "fmt"

func closure() func() { // funções que referenciam variáveis que estão fora do seu corpo
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto) // essa função se 'lembrará' da vaiável texto criada aqui
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure() // mesmo quando ela é chamada em um lugar que possua outra variável texto
	funcaoNova()
}
