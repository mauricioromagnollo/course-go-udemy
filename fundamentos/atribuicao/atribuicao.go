package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Print(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y) // 1 2

	x, y = y, x       // troca o valor das variáveis
	fmt.Println(x, y) // 2 1
}
