package main

import "fmt"

func main() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi: ", comando)
	fmt.Println("E o endereço dela é:", &comando)
}