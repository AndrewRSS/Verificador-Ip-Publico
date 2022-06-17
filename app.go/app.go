package app

import "github.com/urfave/cli"

//Gerar vai retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Verificador de Ip"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	return app
}
