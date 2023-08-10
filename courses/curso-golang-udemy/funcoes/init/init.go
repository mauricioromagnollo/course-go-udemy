package main

import "fmt"

/*
A funão init() é uma convenção em Go para inicializar variáveis, conexões com banco de dados, etc.
Ela é executada antes da função main(). E não precisa ser chamada.
*/
func init() {
	fmt.Print("Inicializando...")
}

func main() {
	fmt.Print("Main...")
}
