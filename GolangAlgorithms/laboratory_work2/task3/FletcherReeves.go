package task3

import (
	"Algorithms/laboratory_work2/controller"
	"math"
)

func FletcherRivesMethod(A [][]float64, B []float64) []float64 {
	// Начальное приближение X0
	X := make([]float64, len(B))

	// Вычисление градиента при X0
	gradient := controller.MatrixMultiply(A, X)
	gradient = controller.SubtractVectors(gradient, B)

	// Начальное значение для направления спуска D
	D := controller.MultiplyVectorByScalar(gradient, -1)

	// Счетчик итераций
	iteration := 0
	maxIterations := 100

	// Точность решения
	epsilon := 0.00001 //4

	// Пока норма градиента больше эпсилон и не достигнуто максимальное количество итераций
	for math.Abs(controller.DotProduct(gradient, gradient)) > epsilon && iteration < maxIterations {
		// Вычисление значения функции на текущей итерации
		f := controller.DotProduct(gradient, gradient)

		// Вычисление значения альфа для текущей итерации
		alpha := f / controller.DotProduct(controller.MatrixMultiply(A, D), D)

		// Обновление значение X
		X = controller.SubtractVectors(X, controller.MultiplyVectorByScalar(D, alpha))

		// Обновление значения градиента
		newGradient := controller.MatrixMultiply(A, X)
		newGradient = controller.SubtractVectors(newGradient, B)

		// Вычисление значения бета для текущей итерации
		beta := controller.DotProduct(newGradient, newGradient) / controller.DotProduct(gradient, gradient)

		// Обновление значения направления спуска
		D = controller.SubtractVectors(controller.MultiplyVectorByScalar(newGradient, -1), controller.MultiplyVectorByScalar(D, beta))

		// Обновление значения градиента
		gradient = newGradient

		iteration++
	}

	return X
}
