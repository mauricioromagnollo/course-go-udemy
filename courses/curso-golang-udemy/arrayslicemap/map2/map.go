package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	// Se Rafael Filho já existisse, o valor seria atualizado. Como não existe, ele será criado!
	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Inexistente") // não dá erro tentar excluir um elemento que não existe

	// O primeiro valor no for é a chave do map, o segundo é o valor
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
