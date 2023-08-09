package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// for parecido com while (não existe while em Go)
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	fmt.Println("Misturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
