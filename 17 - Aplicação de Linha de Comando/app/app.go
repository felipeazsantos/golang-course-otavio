package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gerar vai retornar a aplicação de linha de comanda pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca ips e nomes de servidor na internet"
	flags := []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "www.google.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca ips de endereços da internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca os nomes dos servidores na internet",
			Flags: flags,
			Action: buscaServidoresDaInternet,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidoresDaInternet(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
