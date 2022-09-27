package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}

// go build cria um binário a partir do módulo atual
// utilize ./modulo para executar
