package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de linha de comanda pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca ips e nomes de servidor na internet"

	return app
}
