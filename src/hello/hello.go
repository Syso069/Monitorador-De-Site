package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			exibeMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 3: 
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)
		}
	}
}

func leEExibeIntroducao() string{ 
	fmt.Println("Digite o seu nome:")
	var nomeLido string
	fmt.Scan(&nomeLido)

	return nomeLido
}

func exibeMenu() {
	nome := leEExibeIntroducao()
	fmt.Println("\nOlá Sr(a)", nome,"o que você gostaria de fazer hoje?")
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := [] string{"https://random-status-code.herokuapp.com/","https://www.alura.com.br","https://www.caelum.com.br"}
	sites := leSitesDoArquivo()
	
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site",i,":",site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Print("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	}else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resposta.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites [] string

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

func registraLog (site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+ " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}