package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Imprime o número de processadores disponíveis na sua máquina
	fmt.Println(runtime.NumCPU())
}
