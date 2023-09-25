package main

import "fmt"

// Funções são cidadãos de primeira classe em Go
// Isso quer dizer que você pode atribuir uma função a uma variável
// E essa variável passa a ser uma função
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// Funções anônimas
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
