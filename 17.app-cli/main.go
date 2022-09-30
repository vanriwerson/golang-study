package main

import (
	"app-cli/app" // nome-do-módulo/nome-do-pacote
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil { // Variável erro existe apenas no escopo do if init
		log.Fatal(erro) // Fatal para a execução do programa em caso de erro
	}
}
