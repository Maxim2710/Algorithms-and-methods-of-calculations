package controller

import "math"

// Функция для вычисления экспоненты от x
func FExp(x float64) float64 {
	return math.Exp(x)
}

// Функция для вычисления экспоненты от -x
func FExpNeg(x float64) float64 {
	return math.Exp(-x)
}

// Функция для вычисления гиперболического синуса от x
func FSinh(x float64) float64 {
	return math.Sinh(x)
}

// Функция для вычисления гиперболического косинуса от x
func FCosh(x float64) float64 {
	return math.Cosh(x)
}

// Функция для вычисления синуса от x
func FSin(x float64) float64 {
	return math.Sin(x)
}

// Функция для вычисления косинуса от x
func FCos(x float64) float64 {
	return math.Cos(x)
}

// Функция для вычисления натурального логарифма от x
func FLn(x float64) float64 {
	return math.Log(x)
}

// Функция, для которой вы хотите вычислить производную
func FunctionDif(x float64) float64 {
	return x*x + 3*x + 2 // Пример: квадратичная функция
}
