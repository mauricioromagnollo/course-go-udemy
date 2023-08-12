package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos (dando a sensação que você está herdando de carro)
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // Repare que dentro de ferrari não existe o atributo nome, mas como ferrari tem um campo anonimo carro, você pode acessar o atributo nome de carro
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f) // {{F40 0} true}
}
