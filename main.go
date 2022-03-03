package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre
// o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Quad é responsável por calcula a área de um quadrado
func Quad(base float64) float64 {
	return math.Pow(base,2)
}

// TrianguloEq é responsável por calcular a área de um triangulo Equilátero (Faces iguais)
func TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
