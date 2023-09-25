package main

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// return nota >= 6 ? "Aprovado" : "Reprovado" // Não existe operador ternário em Go
}

func main() {
	println(obterResultado(6.2))
}
