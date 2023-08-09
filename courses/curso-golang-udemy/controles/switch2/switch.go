package main

import "time"

func main() {
	t := time.Now()

	// Quando você passa um switch sem nada, ele assume true e vai procurar um primeiro case que seja verdadeiro
	switch {
	case t.Hour() < 12:
		println("Bom dia!")
	case t.Hour() < 18:
		println("Boa tarde.")
	default:
		println("Boa noite.")
	}
}
