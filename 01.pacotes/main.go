package main

import (
	"fmt"
	"modulo/auxiliar"

	// Pacote externo instalado com go get github.com/badoux/checkmail
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	// Invocam-se os modulos com o nome que está após a última barra
	erro := checkmail.ValidateFormat("user@gmail.com")
	fmt.Println(erro) // retorna <nil> (nulo)
}

// go build cria um binário a partir do módulo atual
// utilize ./modulo para executar
