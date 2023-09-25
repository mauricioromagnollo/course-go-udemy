package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	// Você pode ignorar um dos retornos da função
	// Apenas se você utilizar o _, por exemplo: _, r5 := f5()
	// Não é possível utilizar o _ em mais de uma variável
	// Não é possível utilizar apenas um dos valores retornados, por exemplo r5 := f5(), dado que a função retorna dois valores
	// Você pode chamar a função sem atribuir os valores retornados a variáveis, por exemplo: f5()
	r51, r52 := f5()
	fmt.Println("F5:", r51)
	fmt.Println("F5:", r52)
}
