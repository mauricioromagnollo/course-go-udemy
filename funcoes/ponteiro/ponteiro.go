package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor
	fmt.Println(n) // 1

	// & é usado para obter o endereço da variável
	inc2(&n)       // por referência
	fmt.Println(n) // 2
}

// Por padrão o Go trabalha com passagem de parâmetros por valor. A não ser que você referencie o ponteiro da variável.
// Mesmo sendo objetos, o Go não trabalha com passagem de parâmetros por referência. A não ser que você referencie o ponteiro da variável.
