package controller

import "math"

type Point struct {
	X, Y float64
}

// Базисные функции Бернштейна
func Bernstein(n, i int, t float64) float64 {
	// Реализация формулы Бернштейна
	return float64(Factorial(n)) / (float64(Factorial(i)) * float64(Factorial(n-i))) * math.Pow(t, float64(i)) * math.Pow(1-t, float64(n-i))
}

// Вычисление факториала
func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}
