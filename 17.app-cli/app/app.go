package app

import "github.com/urfave/cli"

// Retorna a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Exemplo de aplicação CLI em Golang"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	return app
}
