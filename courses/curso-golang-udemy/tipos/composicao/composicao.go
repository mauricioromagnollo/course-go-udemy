package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// Uma interface pode herdar outras interfaces
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode adicionar outros métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

// Como todos os métodos da interface esportivoLuxuoso são implementados, então a struct bmw7 é compatível com a interface

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
