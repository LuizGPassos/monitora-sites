package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
		switch comando {
		case 1:
			menu1()
		case 2:
			menu2()
		case 3:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {

	nome := "Luiz"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Esta é a versão", versao, "do programa.")

}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido

}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

}

func menu1() {

	fmt.Println("Iniciando monitoramento...")

	// sites := []string{"https://random-status-code.herokuapp.com", "https://alura.com.br", "https://poggers.com"}
	sites := leSitesArquivo()

	for i := 0; i < 5; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "Foi carregado com sucesso!")
				registraLog(site, true)
			} else {
				fmt.Println("Ocorreu o erro", resp.StatusCode, "ao acessar", site)
				registraLog(site, false)
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Println("")
	}
}

func menu2() {

	fmt.Println("Exibindo Logs...")
	imprimeLog()

}

func leSitesArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

	if err != nil {

		println("Ocorreu o erro:", err)

	}

}

func imprimeLog() {

	arquivo, err := os.Open("log.txt")

	if err != nil {

		println("Erro ao imprimir log: ", err)

	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		println(linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
}
