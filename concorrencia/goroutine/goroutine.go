package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Fale agora", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	/* O programa irá finalizar, antes de chamar Ei ou Opa, imprimindo Fim! */
	/* Não tem nada "esperando" as Goroutines acabarem, para finalizar o programa */
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	// fmt.Println("Fim!")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
}
