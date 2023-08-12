package main

import "math"

// Incializando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)!
// Inicializando com letra minúscula é PRIVADO (visível apenas dentro do pacote)!
// PS: É privado ou público dentro do pacote! Não dentro do arquivo!

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// O nome do pacote que influencia como você vai acessar os métodos, funções, etc. Não o nome do arquivo.

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x) // math.Abs retorna o valor absoluto de um número
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
/*
 * A função Distancia é pública, pois começa com letra maiúscula.
 * A função catetos é privada, pois começa com letra minúscula.
 */
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
