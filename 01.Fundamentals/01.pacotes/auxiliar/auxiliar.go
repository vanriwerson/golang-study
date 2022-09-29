package auxiliar

import "fmt"

func Escrever() { // o nome da função deve iniciar em maiúscula para que ela fique visível para outros pacotes
	fmt.Println("Escrevendo do Pacote auxiliar")
	escrever2() // funções do mesmo pacote não precisam iniciar em letra maiúscula
}
