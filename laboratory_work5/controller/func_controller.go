package controller

import "math"

// функцию для вычисления коэффицентов ряда Фурье
func G(x float64) float64 {
	return math.Sin(x) + math.Cos(2*x) + x*x/10.0
}
