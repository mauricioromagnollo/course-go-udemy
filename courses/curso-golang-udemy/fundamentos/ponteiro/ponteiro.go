package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável i
	*p++   // desreferenciando (pegando o valor)
	i++

	// p++ // Go não tem aritmética de ponteiros
	fmt.Println(p, *p, i, &i)
}
