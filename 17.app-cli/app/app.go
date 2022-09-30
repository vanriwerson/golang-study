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

	app.Commands = []cli.Command{ // configuração da execução do CLI
		{ // configura a invocação da Action:
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			}, // go run main.go ip --host github.com
			Action: buscarIPs,
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
