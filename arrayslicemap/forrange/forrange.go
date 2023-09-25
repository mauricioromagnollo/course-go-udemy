package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta! (tamanho do array) se não tiver os ..., ele se torna um slice
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// ignorando o indice
	for _, num := range numeros {
		fmt.Println(num)
	}

	// o primeiro valor é o indice, o segundo é o valor
}
