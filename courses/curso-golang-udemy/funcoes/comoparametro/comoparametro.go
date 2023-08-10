package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

// Função que recebe outra função como parâmetro
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	// Passando a função multiplicacao como parâmetro para a função exec
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
