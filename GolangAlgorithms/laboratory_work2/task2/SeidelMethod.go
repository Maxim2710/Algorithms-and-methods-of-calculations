package task2

import (
	"Algorithms/laboratory_work2/controller"
	"Algorithms/laboratory_work2/helpers"
)

// SeidelMethod - метод Зейделя
func SeidelMethod(A helpers.Matrix, B helpers.Vector, epsilon float64, maxIterations int) helpers.Vector {
	n := len(A)                     // Получаем размер матрицы A
	X := make(helpers.Vector, n)    // Создаем вектор X длиной n
	Xold := make(helpers.Vector, n) // Создаем вектор Xold длиной n

	for k := 0; k < maxIterations; k++ { // Начинаем цикл итераций
		copy(Xold, X) // Копируем значения из X в Xold

		for i := 0; i < n; i++ { // Начинаем цикл по элементам вектора X
			sum1 := 0.0              // Инициализируем переменную sum1
			for j := 0; j < i; j++ { // Начинаем цикл по элементам вектора X до i-го элемента
				sum1 += A[i][j] * X[j] // Считаем сумму
			}

			sum2 := 0.0                  // Инициализируем переменную sum2
			for j := i + 1; j < n; j++ { // Начинаем цикл по элементам вектора X после i-го элемента
				sum2 += A[i][j] * Xold[j] // Считаем сумму
			}

			X[i] = (B[i] - sum1 - sum2) / A[i][i] // Вычисляем новое значение X[i]
		}

		if controller.CheckAccuracy(X, Xold, epsilon) { // Проверяем точность
			break // Если точность достигнута, выходим из цикла
		}
	}

	return X // Возвращаем результат
}
