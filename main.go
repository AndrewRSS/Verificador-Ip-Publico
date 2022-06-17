package main

import (
	"fmt"
	"log"
	"os"
	"verificador-ip/app.go"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
