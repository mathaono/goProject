package main

import (
	"fmt"
)

func main() {
	nome := "Matheus"
	versao := 1.1

	fmt.Println("Hello World!")
	fmt.Println("Nome:", nome)
	fmt.Println("Programa na versão", versao)
	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)
}
