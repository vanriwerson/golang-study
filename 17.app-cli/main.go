package main

import (
	"app-cli/app" // nome-do-módulo/nome-do-pacote
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar() // retorna a aplicação CLI pronta para ser executada

	if erro := aplicacao.Run(os.Args); erro != nil { // Variável erro existe apenas no escopo do if init
		log.Fatal(erro) // Fatal para a execução do programa em caso de erro
	}
}
