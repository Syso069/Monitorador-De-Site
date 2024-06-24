package main

import (
	"fmt"
	"net/http"
	"os"
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
	sites := [] string{"https://random-status-code.herokuapp.com/","https://www.alura.com.br","https://www.caelum.com.br"}
	
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
	resposta, _ := http.Get(site)

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resposta.StatusCode)
	}
}