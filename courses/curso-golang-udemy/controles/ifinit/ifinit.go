/*
Essa é uma forma de criar um if com inicialização de variável, que é uma
estrutura de controle que não existe em outras linguagens. O escopo da variável
criada é apenas dentro do if.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // também é suportado com switch
		fmt.Print("Ganhou!")
	} else {
		fmt.Print("Perdeu!")
	}
}
