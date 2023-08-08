package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo => ", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 2
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 3
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 1

	c := 3.0
	d := 2.0

	// outras operações usando math...
	fmt.Println("Maior =>", math.Max(c, d))         // Retorna o maior dos dois valores
	fmt.Println("Menor =>", math.Min(c, d))         // Retorna o menor dos dois valores
	fmt.Println("Exponenciação =>", math.Pow(c, d)) // Retorna o primeiro valor elevado ao segundo valor
}
