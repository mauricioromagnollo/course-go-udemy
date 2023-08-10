package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um slice com 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // cria um slice com 10 posições e capacidade de 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Quando você adiciona um elemento há mais do que o tamanho do slice, o Go aumenta a capacidade do slice para o dobro do tamanho
	// do slice anterior. No caso acima, o slice tinha 10 posições e capacidade de 20 posições. Quando adicionamos o 11º elemento, o
	// Go dobrou a capacidade do slice para 40 posições.
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
