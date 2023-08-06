package main

import "fmt"

func main() {
	// a função print não quebra linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// quebra de linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	fmt.Println("O valor de x é", x)
	xs := fmt.Sprint(x) // Retorna uma string

	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x, "!!!")

	// Formatação de string
	// Diminuindo as casas decimais no print
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// %d = int
	// %f = float
	// %.1f = float com uma casa decimal
	// %t = bool
	// %s = string
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// %v = valor genérico (pode ser qualquer tipo)
	fmt.Printf("%v %v %v %v %v", a, b, b, c, d)

	// Podemos interpolar uma string
	str := fmt.Sprintf("Valores: %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Println(str)
}
