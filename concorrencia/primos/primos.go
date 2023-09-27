package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) // Esse close vai evitar o deadlock
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c { // fazendo um for em cima do canal (como Async iterator)
		fmt.Printf("%d ", primo)
	}
	// o close evitou o deadlock e o Fim foi impresso, terminando o programa de forma correta
	fmt.Println("Fim!")
}
