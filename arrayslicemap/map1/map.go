package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados, senão causa erro
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678979] = "Pedro"
	aprovados[12345678980] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678980]) // Ana

	// Excluindo um elemento do map
	delete(aprovados, 12345678980)
}
