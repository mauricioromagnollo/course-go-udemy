package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2   // int
	// fmt.Println(x / y) isso vai gerar um erro, pois é necessário converter o y para float64
	fmt.Println(x / float64(y)) // 1.2

	nota := 6.9
	notaFinal := int(nota) // convertendo float64 para int
	fmt.Println(notaFinal) // 6

	// cuidado... isso não converte o valor para string, mas sim para o valor da tabela ASCII
	fmt.Println("Teste " + string(97)) // Teste a

	// int para string

	fmt.Println("Teste " + strconv.Itoa(97)) // Teste 97

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122) // 1

	// string para bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	str := "Texto"
}

// O Go não considera uma string preenchida, um número, ou algo do tipo como verdadeiro ou falso
// Para converter uma string para bool, é necessário usar o strconv.ParseBool("true")
// O strconv.Atoi("123") retorna dois valores, o primeiro é o valor convertido, e o segundo é um erro
// O strconv.Itoa(97) converte um int para string
// O strconv.ParseBool("true") converte uma string para bool
// O strconv.ParseFloat("1.234", 64) converte uma string para float64
// O strconv.ParseInt("123", 10, 64) converte uma string para int64
