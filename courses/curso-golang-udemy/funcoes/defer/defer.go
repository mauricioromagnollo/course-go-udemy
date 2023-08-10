package main

import "fmt"

/*
O defer é uma função que adia a execução de uma outra função até o momento
que a função que a contém retorne.

O defer é muito utilizado para fechar conexões com banco de dados, fechar
arquivos, etc.
*/
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
