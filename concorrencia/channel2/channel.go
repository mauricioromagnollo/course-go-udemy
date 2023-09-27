package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * base // 2

	time.Sleep(time.Second)
	ch <- 3 * base // 3

	time.Sleep(time.Second * 3)
	ch <- 4 * base // 4
}

func main() {
	ch := make(chan int)

	go doisTresQuatroVezes(2, ch)
	fmt.Println("A")

	a, b := <-ch, <-ch // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-ch)
}

/* OUTPUT:
A
B
4 6
8
*/
