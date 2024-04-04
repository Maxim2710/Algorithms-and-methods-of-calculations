package controller

import "math"

func RungeForAdaptSimpson(a float64, b float64) float64 {
	// Инициализация переменной k значением 4
	var k float64 = 4

	// Возвращение абсолютного значения разности a и b, деленной на (2 в степени k - 1)
	return math.Abs(a-b) / (math.Pow(2, k) - 1)
}
