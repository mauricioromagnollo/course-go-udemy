package main

import "fmt"

// Como o retorno da função está nomeado, não é necessário utilizar o return para retornar os valores
// O retorno limpo é uma funcionalidade do Go
// Quer dizer que eu já atribui os valores que serão retornados pela função, por ter dado nome a eles
// Então basta um return no final da função para retornar os valores
func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
