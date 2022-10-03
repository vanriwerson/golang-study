package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Exemplo de aplicação CLI em Golang"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com", // valor default caso a flag não seja fornecida
		},
	}

	app.Commands = []cli.Command{ // configuração da execução do CLI
		{ // configura a invocação da Action:
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIPs, // executada com: go run main.go ip --host amazon.com.br
		},
		{
			Name:   "server",
			Usage:  "Busca o nome de servidores na internet",
			Flags:  flags,
			Action: buscarServidores, // executada com: go run main.go server --host amazon.com.br
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host") // recupera o valor passado pela flag

	ips, erro := net.LookupIP(host) // retorna um slice de IPs e um erro
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips { // imprime os IPs encontrados
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // Name Server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
