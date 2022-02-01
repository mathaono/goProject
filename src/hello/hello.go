package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 3

func main() {
	lendoArquivos()
	intro()

	for {
		showMenu()
		comando := showAction()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Comando Indisponível")
			os.Exit(-1)
		}
	}
}

func intro() {

	nome := "Matheus"
	versao := 1.1

	fmt.Println("Hello World!")
	fmt.Println("Nome:", nome)
	fmt.Println("Programa na versão", versao)
}

func showMenu() {

	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func showAction() int {

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	sites := lendoArquivos()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i)
			testaSites(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSites(site string) {
	resp, error := http.Get(site)

	if error != nil {
		fmt.Println("Erro:", error)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "com problemas! Status Code:", resp.StatusCode)
	}
}

func lendoArquivos() []string {

	var sites []string

	arquivo, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Erro:", error)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if error == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
