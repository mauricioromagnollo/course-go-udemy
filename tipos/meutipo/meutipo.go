package main

type nota float64 // alias

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

// Você não consegue fazer isso para tipos primitivos diretamente
// func (n float64) entre(inicio, fim float64) bool { por exemplo

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(6.9))
	println(notaParaConceito(2.1))
}
