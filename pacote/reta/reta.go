package main

import "math"

// Se iniciar com letra maiuscula é PUBLIC (visivel dentro e fora do pacote)
// Se iniciar com letra minuscula é PRIVATE (visivel dentro apenas do PACOTE!!!!)

// Por exemplo...
// Ponto - gerará algo PUBLIC
// ponto ou _Ponto - gerará algo PRIVATE

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
