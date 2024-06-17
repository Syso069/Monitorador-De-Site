package main

import (
	"fmt"
	"os"
)

func main() {

	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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