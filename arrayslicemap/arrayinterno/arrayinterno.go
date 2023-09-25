package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // cria um slice com 10 posições e capacidade de 20 posições
	s2 := append(s1, 1, 2, 3) // cria um slice com 10 posições e capacidade de 20 posições
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)
}
