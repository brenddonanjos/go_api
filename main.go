package goarea

import (
	"math"
)

//Pi constante pública
const Pi = 3.1416

//Circulo calcula a área da circunferência
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Retangulo área de retangulos e quadradis
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//função privada no pacote
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
