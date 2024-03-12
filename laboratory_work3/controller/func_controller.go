package controller

import "math"

// Функция, для которой мы ищем корень (в данном случае x - ctg(x))
func Function(x float64) float64 {
	return (x - (1 / math.Tan(x)))
}

// Производная функции
func Df(x float64) float64 {
	return 1 + (1 / (math.Sin(x) * math.Sin(x)))
}

// Функции системы уравнений
func Equations(x, y float64) (float64, float64) {
	// Здесь определяются уравнения системы
	function1 := math.Sin(y+1) - x - 1.2
	function2 := 2*y + math.Cos(x) - 2

	return function1, function2
}

// Функция, для которой мы ищем минимум
func QuadraticFunction(x float64) float64 {
	return (math.Pow(x-11, 2)) + 10
}
